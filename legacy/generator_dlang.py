import json
import os

# Convert manifest types to native types
TYPES_MAP = {
    'void': 'void',
    'bool': 'bool',
    'char8': 'char',
    'char16': 'wchar',
    'int8': 'byte',
    'int16': 'short',
    'int32': 'int',
    'int64': 'long',
    'uint8': 'ubyte',
    'uint16': 'ushort',
    'uint32': 'uint',
    'uint64': 'ulong',
    'ptr64': 'void*',
    'float': 'float',
    'double': 'double',
    'function': 'function',
    'string': 'string',
    'any': 'PlgV',
    'bool[]': 'bool[]',
    'char8[]': 'char[]',
    'char16[]': 'wchar[]',
    'int8[]': 'byte[]',
    'int16[]': 'short[]',
    'int32[]': 'int[]',
    'int64[]': 'long[]',
    'uint8[]': 'ubyte[]',
    'uint16[]': 'ushort[]',
    'uint32[]': 'uint[]',
    'uint64[]': 'ulong[]',
    'ptr64[]': 'void*[]',
    'float[]': 'float[]',
    'double[]': 'double[]',
    'string[]': 'string[]',
    'any[]': 'PlgV[]',
    'vec2[]': 'Vec2[]',
    'vec3[]': 'Vec3[]',
    'vec4[]': 'Vec4[]',
    'mat4x4[]': 'Mat4x4[]',
    'vec2': 'Vec2',
    'vec3': 'Vec3',
    'vec4': 'Vec4',
    'mat4x4': 'Mat4x4'
}

# Convert manifest types to C types
TYPES_C_MAP = {
    'void': 'void',
    'bool': 'bool',
    'char8': 'char',
    'char16': 'wchar',
    'int8': 'byte',
    'int16': 'short',
    'int32': 'int',
    'int64': 'long',
    'uint8': 'ubyte',
    'uint16': 'ushort',
    'uint32': 'uint',
    'uint64': 'ulong',
    'ptr64': 'void*',
    'float': 'float',
    'double': 'double',
    'function': 'function',
    'string': 'ref PlgS',
    'any': 'ref PlgV',
    'bool[]': 'ref PlgA!bool',
    'char8[]': 'ref PlgA!char',
    'char16[]': 'ref PlgA!wchar',
    'int8[]': 'ref PlgA!byte',
    'int16[]': 'ref PlgA!short',
    'int32[]': 'ref PlgA!int',
    'int64[]': 'ref PlgA!long',
    'uint8[]': 'ref PlgA!ubyte',
    'uint16[]': 'ref PlgA!ushort',
    'uint32[]': 'ref PlgA!uint',
    'uint64[]': 'ref PlgA!ulong',
    'ptr64[]': 'ref PlgA!(void*)',
    'float[]': 'ref PlgA!float',
    'double[]': 'ref PlgA!double',
    'string[]': 'ref PlgA!string',
    'any[]': 'ref PlgA!PlgV',
    'vec2[]': 'ref PlgA!Vec2',
    'vec3[]': 'ref PlgA!Vec3',
    'vec4[]': 'ref PlgA!Vec4',
    'mat4x4[]': 'ref PlgA!Mat4x4',
    'vec2': 'ref Vec2',
    'vec3': 'ref Vec3',
    'vec4': 'ref Vec4',
    'mat4x4': 'ref Mat4x4'
}

# Convert native types to C types
N_TYPES_MAP = {
    'void': 'void',
    'bool': 'bool',
    'char': 'char',
    'wchar': 'wchar',
    'byte': 'byte',
    'short': 'short',
    'int': 'int',
    'long': 'long',
    'ubyte': 'ubyte',
    'ushort': 'ushort',
    'uint': 'uint',
    'ulong': 'ulong',
    'void*': 'void*',
    'float': 'float',
    'double': 'double',
    'function': 'function',
    'string': 'ref PlgS',
    'PlgV': 'ref PlgV',
    'bool[]': 'ref PlgA!bool',
    'char[]': 'ref PlgA!char',
    'wchar[]': 'ref PlgA!wchar',
    'byte[]': 'ref PlgA!byte',
    'short[]': 'ref PlgA!short',
    'int[]': 'ref PlgA!int',
    'long[]': 'ref PlgA!long',
    'ubyte[]': 'ref PlgA!ubyte',
    'ushort[]': 'ref PlgA!ushort',
    'uint[]': 'ref PlgA!uint',
    'ulong[]': 'ref PlgA!ulong',
    'void*[]': 'ref PlgA!(void*)',
    'float[]': 'ref PlgA!float',
    'double[]': 'ref PlgA!double',
    'string[]': 'ref PlgA!string',
    'PlgV[]': 'ref PlgA!PlgV',
    'Vec2[]': 'ref PlgA!Vec2',
    'Vec3[]': 'ref PlgA!Vec3',
    'Vec4[]': 'ref PlgA!Vec4',
    'Mat4x4[]': 'ref PlgA!Max4x4',
    'Vec2': 'ref Vec2',
    'Vec3': 'ref Vec3',
    'Vec4': 'ref Vec4',
    'Mat4x4': 'ref Mat4x4'
}

INVALID_NAMES = {
    "out",
    "version",
    "module",
    "function"
}

def get_enum(param, tp):
    enum = {
        "type": TYPES_MAP[tp],
        "desc": param["description"] if "description" in param else "",
        "values": {}
    }
    for v in param["values"]:
        enum["values"][v["name"]] = {
            "value": v["value"],
            "desc": v["description"] if "description" in v else ""
        }
    return enum

modules = {}
enums = {}

def gen_import(file_path):
    obj = None
    with open(file_path, 'r') as f:
        obj = json.load(f)

    for m in obj["methods"]:
        module_name = m["group"].lower()
        if not (module_name in modules):
            modules[module_name] = {
                "funcs": {},
                "alias": {}
            }
        func_name_disp = m["name"]
        func_name = m["funcName"]
        func_desc = m["description"]

        # Parse parameters
        params = {}
        for param in m["paramTypes"]:
            param_name = param["name"]
            if param_name in INVALID_NAMES:
                param_name = '_' + param_name
            if param["type"] == "function":
                params[param_name] = {
                    "type": param["prototype"]["funcName"],
                    "ref": False,
                    "desc": param["description"]
                }
                alias = {
                    "desc": param["prototype"]["description"],
                    "params": {},
                    "ret": {}
                }
                for param2 in param["prototype"]["paramTypes"]:
                    param_name2 = param2["name"]
                    if param_name2 in INVALID_NAMES:
                        param_name2 = '_' + param_name2
                    alias["params"][param_name2] = {
                        "type": TYPES_C_MAP[param2["type"]],
                        "ref": param2["ref"] if "ref" in param2 else False,
                        "desc": param2["description"]
                    }
                    if not alias["params"][param_name2]["ref"]:
                        alias["params"][param_name2]["type"] = "const " + alias["params"][param_name2]["type"]
                    if "enum" in param2:
                        alias["params"][param_name2]["type"] = param2["enum"]["name"]
                        enums[param2["enum"]["name"]] = get_enum(param2["enum"], param2["type"])
                ret = {
                    "type": TYPES_MAP[param["prototype"]["retType"]["type"]],
                    "desc": param["prototype"]["retType"]["description"] if "description" in param["prototype"]["retType"] else ""
                }
                if "enum" in param["prototype"]["retType"]:
                    ret["type"] = param["prototype"]["retType"]["enum"]["name"]
                    enums[param["prototype"]["retType"]["enum"]["name"]] = get_enum(param["prototype"]["retType"]["enum"], param["prototype"]["retType"]["type"])
                alias["ret"] = ret

                modules[module_name]["alias"][param["prototype"]["funcName"]] = alias
            else:
                params[param_name] = {
                    "type": TYPES_MAP[param["type"]],
                    "ref": param["ref"] if "ref" in param else False,
                    "desc": param["description"]
                }
                if "enum" in param:
                    params[param_name]["type"] = param["enum"]["name"]
                    enums[param["enum"]["name"]] = get_enum(param["enum"], param["type"])

        # Parse return
        ret = {
            "type": TYPES_MAP[m["retType"]["type"]],
            "desc": m["retType"]["description"] if "description" in m["retType"] else ""
        }
        if "enum" in m["retType"]:
            ret["type"] = m["retType"]["enum"]["name"]

            enums[m["retType"]["enum"]["name"]] = get_enum(m["retType"]["enum"], m["retType"]["type"])

        modules[module_name]["funcs"][func_name_disp] = {
            "internalName": func_name,
            "desc": func_desc,
            "params": params,
            "ret": ret
        }

    # All methods parsed
    path = "source" + os.sep + "imported" + os.sep + obj["name"].lower() + os.sep
    os.makedirs(path, exist_ok=True)
    for module_name, kv in modules.items():
        with open(path + module_name + ".d", "wt") as f:
            f.write("module imported." + obj["name"].lower() + "." + module_name + ";\n\n")
            f.write("import plugify.internals;\n")
            f.write("public import plugify;\n")
            f.write("public import imported." + obj["name"].lower() + ".enums;\n\n")

            for alias_name, akv in kv["alias"].items():
                params = []
                body_text = "alias " + alias_name + " = extern (C) " + akv["ret"]["type"] + " function("
                f.write("/++\n")
                if akv["desc"]:
                    f.write('\t' + akv["desc"] + "\n\n")
                header_text = ""
                if akv["params"]:
                    f.write("\tParams:\n")
                    for param_name, pkv in akv["params"].items():
                        header_text += "\t\t" + param_name + " = " + pkv["desc"] + "\n"
                        params.append(pkv["type"] + ' ' + param_name)
                header_text += "+/\n"

                body_text += ", ".join(params)
                body_text += ");\n\n"
                f.write(header_text)
                f.write(body_text)

            for func_name, fkv in kv["funcs"].items():
                params = []
                body_text = fkv["ret"]["type"] + " " + func_name + "("
                f.write("/++\n")
                if fkv["desc"]:
                    f.write('\t' + fkv["desc"] + "\n\n")
                header_text = ""
                body2_text = ""
                call_text = []

                counter = 0
                if fkv["params"]:
                    f.write("\tParams:\n")
                    for param_name, pkv in fkv["params"].items():
                        header_text += '\t\t' + param_name + " = " + pkv["desc"] + "\n"
                        p = pkv["type"] + ' ' + param_name
                        if pkv["ref"]:
                            p = "ref " + p
                        params.append(p)

                        c_ret_type = ""
                        if pkv["type"] in N_TYPES_MAP:
                            c_ret_type = N_TYPES_MAP[pkv["type"]]
                        else:
                            c_ret_type = pkv["type"]
                        if c_ret_type.startswith("ref") and ("PlgA" in c_ret_type or "PlgS" in c_ret_type):
                            val_type = c_ret_type[4:]
                            body2_text += '\t' + val_type + " _var" + str(counter) + " = (" + param_name + ");\n"
                            if pkv["ref"]:
                                body2_text += "\tscope(exit) { " + param_name + " = _var" + str(counter) + "; }\n"
                            body2_text += '\n'
                            call_text.append("_var" + str(counter))
                        else:
                            call_text.append(param_name)
                        counter += 1
                if fkv["ret"]["type"] != "void":
                    header_text += "\tReturns:\n"
                    header_text += "\t\t" + fkv["ret"]["desc"] + '\n'
                header_text += "+/\n"

                body_text += ", ".join(params)
                body_text += ") {\n"
                f.write(header_text)
                f.write(body_text)
                f.write(body2_text + '\t')
                c_ret_type = ""
                if fkv["ret"]["type"] in N_TYPES_MAP:
                    c_ret_type = N_TYPES_MAP[fkv["ret"]["type"]]
                else:
                    c_ret_type = fkv["ret"]["type"]
                if not (c_ret_type == "void"):
                    f.write("return ")
                f.write("__" + func_name + '(')
                f.write(", ".join(call_text) + ')')
                if "!" in c_ret_type:
                    f.write(".value")
                f.write(";\n}\n\n")

            aliased = ""
            statics = ""
            for func_name, fkv in kv["funcs"].items():
                params = []
                c_ret_type = ""
                if fkv["ret"]["type"] in N_TYPES_MAP:
                    c_ret_type = N_TYPES_MAP[fkv["ret"]["type"]]
                else:
                    c_ret_type = fkv["ret"]["type"]
                if c_ret_type.startswith("ref"):
                    c_ret_type = c_ret_type[4:]
                aliased += "\talias _" + func_name + " = extern (C) " + c_ret_type + " function("
                for param_name, pkv in fkv["params"].items():
                    c_ret_type = ""
                    if pkv["type"] in N_TYPES_MAP:
                        c_ret_type = N_TYPES_MAP[pkv["type"]]
                    else:
                        c_ret_type = pkv["type"]

                    if c_ret_type.startswith("ref"):
                        if not pkv["ref"]:
                            c_ret_type = "const " + c_ret_type
                    elif pkv["ref"]:
                        c_ret_type = "ref " + c_ret_type
                    if c_ret_type == "function":
                        c_ret_type = pkv["type"]
                    p = c_ret_type + ' ' + param_name
                    params.append(p)
                aliased += ", ".join(params)
                aliased += ");\n"
                aliased += '\t__gshared _' + func_name + " __" + func_name + ";\n"
                statics += '\t__' + func_name + " = cast(_" + func_name + ')_MethodPointer("' + fkv["internalName"] + '");\n'
            f.write("private {\n")
            f.write(aliased)
            f.write("}\n\nshared static this() {\n")
            f.write(statics)
            f.write("}\n")


    with open(path + "enums.d", "wt") as f:
        f.write("module imported." + obj["name"].lower() + ".enums;\n\n")
        for enum_name, ekv in enums.items():
            if ekv["desc"]:
                f.write("/// " + ekv["desc"] + "\n")
            f.write("enum " + enum_name + " : " + ekv["type"] + " {\n")
            for enum_v_name, enum_v in ekv["values"].items():
                enum_val = f"{enum_v["value"]:,}".split(',')
                f.write("\t" + enum_v_name + " = " + '_'.join(enum_val) + ",")
                if enum_v["desc"]:
                    f.write(" /// " + enum_v["desc"])
                f.write('\n')
            f.write("}\n\n")


with os.scandir("import_plugins") as entries:
    for entry in entries:
        if entry.is_file() and entry.name.endswith(".pplugin.in"):
            gen_import("import_plugins" + os.sep + entry.name)
