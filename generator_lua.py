#!/usr/bin/python3
import sys
import argparse
import os
import json
from enum import IntEnum


INVALID_NAMES = [
    'and',       # Logical operator
    'break',     # Loop control
    'do',        # Block delimiter
    'else',      # Conditional
    'elseif',    # Conditional
    'end',       # Block delimiter
    'false',     # Boolean literal
    'for',       # Loop control
    'function',  # Function declaration
    'goto',      # Goto statement
    'if',        # Conditional
    'in',        # Loop operator
    'local',     # Local variable declaration
    'nil',       # Null value
    'not',       # Logical operator
    'or',        # Logical operator
    'repeat',    # Loop control
    'return',    # Return statement
    'then',      # Conditional
    'true',      # Boolean literal
    'until',     # Loop control
    'while'      # Loop control
]


def generate_name(name: str) -> str:
    """Generate a valid Lua variable name."""
    return f'{name}_' if name in INVALID_NAMES else name


def gen_params(params: list[dict]) -> str:
    """Generate Lua function parameters as strings."""
    def gen_param(index: int, param: dict) -> str:
        return generate_name(param.get('name', f'p{index}'))

    return ', '.join(gen_param(i, p) for i, p in enumerate(params))


def gen_documentation(method: dict) -> str:
    """
    Generate a Lua-style function documentation string from a JSON block.

    Args:
        method (Dict[str, Any]): The input JSON data describing the function.

    Returns:
        str: The generated Lua documentation string.
    """
    name = method.get('name', 'UnnamedFunction')
    description = method.get('description', 'No description provided.')
    param_types = method.get('paramTypes', [])
    ret_type_info = method.get('retType', {})
    ret_type = ret_type_info.get('type', 'void')

    docstring = [f"--- {description}"]

    # Add parameters
    for param in param_types:
        param_name = param.get('name', 'UnnamedParam')
        param_type = param.get('type', 'any')
        param_desc = param.get('description', 'No description available.')
        docstring.append(f'-- @param {param_name} {param_type} {param_desc}')

    # Add return type
    if ret_type.lower() != 'void':
        ret_desc = ret_type_info.get('description', 'No description available.')
        docstring.append(f'-- @return {ret_type} {ret_desc}')

    # Add callback prototype if present
    for param in param_types:
        if param.get('type') == 'function' and 'prototype' in param:
            prototype = param['prototype']
            proto_name = prototype.get('name', 'UnnamedCallback')
            proto_desc = prototype.get('description', 'No description provided.')
            docstring.append(f"-- @callback {proto_name} {param.get('name', '')} - {proto_desc}")
            for proto_param in prototype.get('paramTypes', []):
                p_name = proto_param.get('name', 'UnnamedParam')
                p_type = proto_param.get('type', 'any')
                p_desc = proto_param.get('description', 'No description available.')
                docstring.append(f'-- @param {p_name} {p_type} {p_desc}')
            if 'retType' in prototype:
                p_ret = prototype['retType']
                p_ret_type = p_ret.get('type', 'void')
                p_ret_desc = p_ret.get('description', 'No description available.')
                docstring.append(f'-- @return {p_ret_type} {p_ret_desc}')

    return '\n'.join(docstring)


def gen_enum_body(enum: dict, enums: set[str]) -> str:
    """
    Generates a Lua enum-like table definition from the provided enum metadata.

    Args:
        enum (dict): The JSON dictionary describing the enum.
        enums (set): A set to track already defined enums to prevent duplicates.

    Returns:
        str: The generated Lua enum table code or an empty string if the enum already exists.
    """
    enum_name = enum.get('name', 'InvalidEnum')
    enum_description = enum.get('description', '')
    enum_values = enum.get('values', [])

    if enum_name in enums:
        return ''  # Skip if already generated

    enums.add(enum_name)

    enum_code = []

    # Add description as a Lua comment
    if enum_description:
        enum_code.append(f'-- {enum_name}: {enum_description}')
    else:
        enum_code.append(f'-- Enum: {enum_name}')

    # Begin Lua table
    enum_code.append(f'{enum_name} = {{')

    for i, value in enumerate(enum_values):
        name = value.get('name', f'InvalidName_{i}')
        enum_value = value.get('value', str(i))
        description = value.get('description', '')

        if description:
            enum_code.append(f'  -- {description}')
        enum_code.append(f'  {name} = {enum_value},')

    # Close the table
    enum_code.append('}\n')

    return '\n'.join(enum_code)


def generate_enum_code(pplugin: dict, enums: set[str]) -> str:
    """
    Generate  JavaScript enum-like object code from a plugin definition.
    """
    # Container for all generated enum code
    content = []

    def process_enum(enum_data: dict):
        """
        Generate enum code from the given enum data if it hasn't been processed.
        """
        enum_code = gen_enum_body(enum_data, enums)
        if enum_code:
            content.append(enum_code)
            content.append('\n')

    def process_prototype(prototype: dict):
        """
        Recursively process a function prototype for enums.
        """
        if 'enum' in prototype.get('retType', {}):
            process_enum(prototype['retType']['enum'])

        for param in prototype.get('paramTypes', []):
            if 'enum' in param:
                process_enum(param['enum'])
            if 'prototype' in param:  # Process nested prototypes
                process_prototype(param['prototype'])

    # Main loop: Process all exported methods in the plugin
    for method in pplugin.get('methods', []):
        if 'retType' in method and 'enum' in method['retType']:
            process_enum(method['retType']['enum'])

        for param in method.get('paramTypes', []):
            if 'enum' in param:
                process_enum(param['enum'])
            if 'prototype' in param:  # Handle nested function prototypes
                process_prototype(param['prototype'])

    # Join all generated enums into a single string
    return '\n'.join(content)

    
def generate_stub(plugin_name: str, pplugin: dict) -> str:
    """Generate JavaScript stub content."""
    link = 'https://github.com/untrustedmodders/plugify-module-lua/blob/main/generator/generator.py'
    content = [
        f'-- Generated from {plugin_name}.pplugin by {link}\n\n']

    # Append enum definitions
    enums = set()
    content.append(generate_enum_code(pplugin, enums))

    # Append method stubs
    for method in pplugin.get('methods', []):
        method_name = method.get('name', 'UnnamedMethod')
        param_types = method.get('paramTypes', [])
        ret_type = method.get('retType', {})

        # Add the method signature and documentation
        content.append(gen_documentation(method))  # Use JS-specific doc generator
        content.append(f'function {method_name}({gen_params(param_types)}) end\n')

    return '\n'.join(content)


def main(manifest_path: str, output_dir: str, override: bool):
    """Main entry point for the script."""
    if not os.path.isfile(manifest_path):
        print(f'Manifest file does not exist: {manifest_path}')
        return 1
    if not os.path.isdir(output_dir):
        print(f'Output directory does not exist: {output_dir}')
        return 1

    try:
        with open(manifest_path, 'r', encoding='utf-8') as file:
            pplugin = json.load(file)

    except Exception as e:
        print(f'An error occurred: {e}')
        return 1

    plugin_name = pplugin.get('name', os.path.basename(manifest_path).rsplit('.', 3)[0])
    output_path = os.path.join(output_dir, 'pps', f'{plugin_name}.lua')
    os.makedirs(os.path.dirname(output_path), exist_ok=True)

    if os.path.isfile(output_path) and not override:
        print(f'Output file already exists: {output_path}. Use --override to overwrite existing file.')
        return 1

    try:
        content = generate_stub(plugin_name, pplugin)
        
        with open(output_path, 'w', encoding='utf-8') as file:
            file.write(content)

    except Exception as e:
        print(f'An error occurred: {e}')
        return 1
    
    print(f'Stub generated at: {output_path}')
    return 0


def get_args():
    """Parse command-line arguments."""
    parser = argparse.ArgumentParser(description='Generate Lua .lua stub files for plugin manifests.')
    parser.add_argument('manifest', help='Path to the plugin manifest file')
    parser.add_argument('output', help='Output directory for the generated stub')
    parser.add_argument('--override', action='store_true', help='Override existing files')
    return parser.parse_args()


if __name__ == '__main__':
    args = get_args()
    sys.exit(main(args.manifest, args.output, args.override))