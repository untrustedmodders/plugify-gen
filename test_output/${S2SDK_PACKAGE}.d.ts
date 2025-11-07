// Generated from ${S2SDK_PACKAGE}.pplugin

declare module "plugify" {
  type Vector2 = { x: number; y: number };
  type Vector3 = { x: number; y: number; z: number };
  type Vector4 = { x: number; y: number; z: number; w: number };
  type Matrix4x4 = { data: number[] };
}

declare module ":${S2SDK_PACKAGE}" {
  import { Vector2, Vector3, Vector4, Matrix4x4 } from "plugify";

  /** Enum representing various movement types for entities. */
  export const enum MoveType {
    /** Never moves. */
    None = 0,
    /** Previously isometric movement type. */
    Isometric = 1,
    /** Player only - moving on the ground. */
    Walk = 2,
    /** No gravity, but still collides with stuff. */
    Fly = 3,
    /** Flies through the air and is affected by gravity. */
    Flygravity = 4,
    /** Uses VPHYSICS for simulation. */
    Vphysics = 5,
    /** No clip to world, push and crush. */
    Push = 6,
    /** No gravity, no collisions, still has velocity/avelocity. */
    Noclip = 7,
    /** Used by players only when going onto a ladder. */
    Ladder = 8,
    /** Observer movement, depends on player's observer mode. */
    Observer = 9,
    /** Allows the entity to describe its own physics. */
    Custom = 10
  }

  /** Enum representing rendering modes for materials. */
  export const enum RenderMode {
    /** Standard rendering mode (src). */
    Normal = 0,
    /** Composite: c*a + dest*(1-a). */
    TransColor = 1,
    /** Composite: src*a + dest*(1-a). */
    TransTexture = 2,
    /** Composite: src*a + dest -- No Z buffer checks -- Fixed size in screen space. */
    Glow = 3,
    /** Composite: src*srca + dest*(1-srca). */
    TransAlpha = 4,
    /** Composite: src*a + dest. */
    TransAdd = 5,
    /** Not drawn, used for environmental effects. */
    Environmental = 6,
    /** Uses a fractional frame value to blend between animation frames. */
    TransAddFrameBlend = 7,
    /** Composite: src + dest*(1-a). */
    TransAlphaAdd = 8,
    /** Same as Glow but not fixed size in screen space. */
    WorldGlow = 9,
    /** No rendering. */
    None = 10,
    /** Developer visualizer rendering mode. */
    DevVisualizer = 11
  }

  /** Enum representing the possible teams in Counter-Strike. */
  export const enum CSTeam {
    /** No team. */
    None = 0,
    /** Spectator team. */
    Spectator = 1,
    /** Terrorist team. */
    T = 2,
    /** Counter-Terrorist team. */
    CT = 3
  }

  /** Represents the possible types of data that can be passed as a value in input actions. */
  export const enum FieldType {
    /** Automatically detect the type of the value. */
    Auto = 0,
    /** A 32-bit floating-point number. */
    Float32 = 1,
    /** A 64-bit floating-point number. */
    Float64 = 2,
    /** A 32-bit signed integer. */
    Int32 = 3,
    /** A 32-bit unsigned integer. */
    UInt32 = 4,
    /** A 64-bit signed integer. */
    Int64 = 5,
    /** A 64-bit unsigned integer. */
    UInt64 = 6,
    /** A boolean value (true or false). */
    Boolean = 7,
    /** A single character. */
    Character = 8,
    /** A managed string object. */
    String = 9,
    /** A null-terminated C-style string. */
    CString = 10,
    /** A script handle, typically for scripting integration. */
    HScript = 11,
    /** An entity handle, used to reference an entity within the system. */
    EHandle = 12,
    /** A resource handle, such as a file or asset reference. */
    Resource = 13,
    /** A 3D vector, typically representing position or direction. */
    Vector3d = 14,
    /** A 2D vector, for planar data or coordinates. */
    Vector2d = 15,
    /** A 4D vector, often used for advanced mathematical representations. */
    Vector4d = 16,
    /** A 32-bit color value (RGBA). */
    Color32 = 17,
    /** A quaternion-based angle representation. */
    QAngle = 18,
    /** A quaternion, used for rotation and orientation calculations. */
    Quaternion = 19
  }

  /** Enum representing various damage types. */
  export const enum DamageTypes {
    /** Generic damage. */
    DMG_GENERIC = 0,
    /** Crush damage. */
    DMG_CRUSH = 1,
    /** Bullet damage. */
    DMG_BULLET = 2,
    /** Slash damage. */
    DMG_SLASH = 4,
    /** Burn damage. */
    DMG_BURN = 8,
    /** Vehicle damage. */
    DMG_VEHICLE = 16,
    /** Fall damage. */
    DMG_FALL = 32,
    /** Blast damage. */
    DMG_BLAST = 64,
    /** Club damage. */
    DMG_CLUB = 128,
    /** Shock damage. */
    DMG_SHOCK = 256,
    /** Sonic damage. */
    DMG_SONIC = 512,
    /** Energy beam damage. */
    DMG_ENERGYBEAM = 1024,
    /** Drowning damage. */
    DMG_DROWN = 16384,
    /** Poison damage. */
    DMG_POISON = 32768,
    /** Radiation damage. */
    DMG_RADIATION = 65536,
    /** Recovering from drowning damage. */
    DMG_DROWNRECOVER = 131072,
    /** Acid damage. */
    DMG_ACID = 262144,
    /** Physgun damage. */
    DMG_PHYSGUN = 1048576,
    /** Dissolve damage. */
    DMG_DISSOLVE = 2097152,
    /** Surface blast damage. */
    DMG_BLAST_SURFACE = 4194304,
    /** Buckshot damage. */
    DMG_BUCKSHOT = 16777216,
    /** Last generic flag damage. */
    DMG_LASTGENERICFLAG = 16777216,
    /** Headshot damage. */
    DMG_HEADSHOT = 33554432,
    /** Danger zone damage. */
    DMG_DANGERZONE = 67108864
  }

  /** Enum representing various flags for ConVars and ConCommands. */
  export const enum ConVarFlag {
    /** The default, no flags at all. */
    None = 0,
    /** Linked to a ConCommand. */
    LinkedConcommand = 1,
    /** Hidden in released products. Automatically removed if ALLOW_DEVELOPMENT_CVARS is defined. */
    DevelopmentOnly = 2,
    /** Defined by the game DLL. */
    GameDll = 4,
    /** Defined by the client DLL. */
    ClientDll = 8,
    /** Hidden. Doesn't appear in find or auto-complete. Like DEVELOPMENTONLY but cannot be compiled out. */
    Hidden = 16,
    /** Server cvar; data is not sent since it's sensitive (e.g., passwords). */
    Protected = 32,
    /** This cvar cannot be changed by clients connected to a multiplayer server. */
    SpOnly = 64,
    /** Saved to vars.rc. */
    Archive = 128,
    /** Notifies players when changed. */
    Notify = 256,
    /** Changes the client's info string. */
    UserInfo = 512,
    /** Hides the cvar from lookups. */
    Missing0 = 1024,
    /** If this is a server cvar, changes are not logged to the file or console. */
    Unlogged = 2048,
    /** Hides the cvar from lookups. */
    Missing1 = 4096,
    /** Server-enforced setting on clients. */
    Replicated = 8192,
    /** Only usable in singleplayer/debug or multiplayer with sv_cheats. */
    Cheat = 16384,
    /** Causes auto-generated varnameN for splitscreen slots. */
    PerUser = 32768,
    /** Records this cvar when starting a demo file. */
    Demo = 65536,
    /** Excluded from demo files. */
    DontRecord = 131072,
    /** Reserved for future use. */
    Missing2 = 262144,
    /** Cvars tagged with this are available to customers. */
    Release = 524288,
    /** Marks the cvar as a menu bar item. */
    MenuBarItem = 1048576,
    /** Reserved for future use. */
    Missing3 = 2097152,
    /** Cannot be changed by a client connected to a server. */
    NotConnected = 4194304,
    /** Enables fuzzy matching for vconsole. */
    VconsoleFuzzyMatching = 8388608,
    /** The server can execute this command on clients. */
    ServerCanExecute = 16777216,
    /** Allows clients to execute this command. */
    ClientCanExecute = 33554432,
    /** The server cannot query this cvar's value. */
    ServerCannotQuery = 67108864,
    /** Sets focus in the vconsole. */
    VconsoleSetFocus = 134217728,
    /** IVEngineClient::ClientCmd can execute this command. */
    ClientCmdCanExecute = 268435456,
    /** Executes the cvar every tick. */
    ExecutePerTick = 536870912
  }

  /** Handles the execution of a command triggered by a caller. This function processes the command, interprets its context, and handles any provided arguments. */
  export type CommandCallback = (number caller, number context, string[] arguments_) => number;

  /** Enum representing the type of callback. */
  export const enum HookMode {
    /** Callback will be executed before the original function */
    Pre = 0,
    /** Callback will be executed after the original function */
    Post = 1
  }

  /** The command execution context. */
  export const enum CommandCallingContext {
    /** The command execute from the client's console. */
    Console = 0,
    /** The command execute from the client's chat. */
    Chat = 1
  }

  export const enum ConVarType {
    /** Invalid type */
    Invalid = -1,
    /** Boolean type */
    Bool = 0,
    /** 16-bit signed integer */
    Int16 = 1,
    /** 16-bit unsigned integer */
    UInt16 = 2,
    /** 32-bit signed integer */
    Int32 = 3,
    /** 32-bit unsigned integer */
    UInt32 = 4,
    /** 64-bit signed integer */
    Int64 = 5,
    /** 64-bit unsigned integer */
    UInt64 = 6,
    /** 32-bit floating point */
    Float32 = 7,
    /** 64-bit floating point (double) */
    Float64 = 8,
    /** String type */
    String = 9,
    /** Color type */
    Color = 10,
    /** 2D vector */
    Vector2 = 11,
    /** 3D vector */
    Vector3 = 12,
    /** 4D vector */
    Vector4 = 13,
    /** Quaternion angle */
    Qangle = 14,
    /** Maximum value (used for bounds checking) */
    Max = 15
  }

  /** Handles changes to a console variable's value. This function is called whenever the value of a specific console variable is modified. */
  export type ChangeCallback = (bigint conVarHandle, string newValue, string oldValue) => void;

  /** Handles changes to a console variable's value. This function is called whenever the value of a specific console variable is modified. */
  export type CvarValueCallback = (number playerSlot, number cookie, number code, string name, string value, any[] data) => void;

  /** Defines a QueueTask Callback. */
  export type TaskCallback = (any[] userData) => void;

  /** This function is a callback handler for entity output events. It is triggered when a specific output event is activated, and it handles the process by passing the activator, the caller, and a delay parameter for the output. */
  export type HookEntityOutputCallback = (number activatorHandle, number callerHandle, number flDelay) => number;

  /** Enum representing the type of callback. */
  export const enum EventHookError {
    /** Indicates that the event hook was successfully created. */
    Okay = 0,
    /** Indicates that the event name provided is invalid or does not exist. */
    InvalidEvent = 1,
    /** Indicates that the event system is not currently active or initialized. */
    NotActive = 2,
    /** Indicates that the callback function provided is invalid or not compatible with the event system. */
    InvalidCallback = 3
  }

  /** Handles events triggered by the game event system. This function processes the event data, determines the necessary action, and optionally prevents event broadcasting. */
  export type EventCallback = (string name, bigint event, boolean dontBroadcast) => number;

  /** Enum representing the possible verbosity of a logger. */
  export const enum LoggingVerbosity {
    /** Turns off all spew. */
    Off = 0,
    /** Turns on vital logs. */
    Essential = 1,
    /** Turns on most messages. */
    Default = 2,
    /** Allows for walls of text that are usually useful. */
    Detailed = 3,
    /** Allows everything. */
    Max = 4
  }

  /** Enum representing the possible verbosity of a logger. */
  export const enum LoggingSeverity {
    /** Turns off all spew. */
    Off = 0,
    /** A debug message. */
    Detailed = 1,
    /** An informative logging message. */
    Message = 2,
    /** A warning, typically non-fatal. */
    Warning = 3,
    /** A message caused by an Assert**() operation. */
    Assert = 4,
    /** An error, typically fatal/unrecoverable. */
    Error = 5
  }

  /** Logging channel behavior flags, set on channel creation. */
  export const enum LoggingChannelFlags {
    /** Indicates that the spew is only relevant to interactive consoles. */
    ConsoleOnly = 1,
    /** Indicates that spew should not be echoed to any output devices. */
    DoNotEcho = 2
  }

  /** Enum representing the possible reasons a vote creation or processing has failed. */
  export const enum VoteCreateFailed {
    /** Generic vote failure. */
    Generic = 0,
    /** Vote failed due to players transitioning. */
    TransitioningPlayers = 1,
    /** Vote failed because vote rate limit was exceeded. */
    RateExceeded = 2,
    /** Vote failed because Yes votes must exceed No votes. */
    YesMustExceedNo = 3,
    /** Vote failed due to quorum not being met. */
    QuorumFailure = 4,
    /** Vote failed because the issue is disabled. */
    IssueDisabled = 5,
    /** Vote failed because the map was not found. */
    MapNotFound = 6,
    /** Vote failed because map name is required. */
    MapNameRequired = 7,
    /** Vote failed because a similar vote failed recently. */
    FailedRecently = 8,
    /** Vote to kick failed recently. */
    FailedRecentKick = 9,
    /** Vote to change map failed recently. */
    FailedRecentChangeMap = 10,
    /** Vote to swap teams failed recently. */
    FailedRecentSwapTeams = 11,
    /** Vote to scramble teams failed recently. */
    FailedRecentScrambleTeams = 12,
    /** Vote to restart failed recently. */
    FailedRecentRestart = 13,
    /** Team is not allowed to call vote. */
    TeamCantCall = 14,
    /** Vote failed because game is waiting for players. */
    WaitingForPlayers = 15,
    /** Target player was not found. */
    PlayerNotFound = 16,
    /** Cannot kick an admin. */
    CannotKickAdmin = 17,
    /** Scramble is currently in progress. */
    ScrambleInProgress = 18,
    /** Swap is currently in progress. */
    SwapInProgress = 19,
    /** Spectators are not allowed to vote. */
    Spectator = 20,
    /** Voting is disabled. */
    Disabled = 21,
    /** Next level is already set. */
    NextLevelSet = 22,
    /** Rematch vote failed. */
    Rematch = 23,
    /** Vote to surrender failed due to being too early. */
    TooEarlySurrender = 24,
    /** Vote to continue failed. */
    Continue = 25,
    /** Vote failed because match is already paused. */
    MatchPaused = 26,
    /** Vote failed because match is not paused. */
    MatchNotPaused = 27,
    /** Vote failed because game is not in warmup. */
    NotInWarmup = 28,
    /** Vote failed because there are not 10 players. */
    Not10Players = 29,
    /** Vote failed due to an active timeout. */
    TimeoutActive = 30,
    /** Vote failed because timeout is inactive. */
    TimeoutInactive = 31,
    /** Vote failed because timeout has been exhausted. */
    TimeoutExhausted = 32,
    /** Vote failed because the round can't end now. */
    CantRoundEnd = 33,
    /** Sentinel value. Not a real failure reason. */
    Max = 34
  }

  /** Handles the final result of a Yes/No vote. This function is called when a vote concludes, and is responsible for determining whether the vote passed based on the number of 'yes' and 'no' votes. Also receives context about the clients who participated in the vote. */
  export type YesNoVoteResult = (number numVotes, number yesVotes, number noVotes, number numClients, number[] clientInfoSlot, number[] clientInfoItem) => boolean;

  export type YesNoVoteHandler = (number action, number clientSlot, number choice) => void;

  /** Enum representing the possible types of a vote. */
  export const enum VoteEndReason {
    /** All possible votes were cast. */
    AllVotes = 0,
    /** Time ran out. */
    TimeUp = 1,
    /** The vote got cancelled. */
    Cancelled = 2
  }

  /** This function is invoked when a timer event occurs. It handles the timer-related logic and performs necessary actions based on the event. */
  export type TimerCallback = (number timer, any[] userData) => void;

  /** Enum representing the possible flags of a timer. */
  export const enum TimerFlag {
    /** Timer with no unique properties. */
    Default = 0,
    /** Timer will repeat until stopped. */
    Repeat = 1,
    /** Timer will not carry over mapchanges. */
    NoMapChange = 2
  }

  /** Called on client connection. If you return true, the client will be allowed in the server. If you return false (or return nothing), the client will be rejected. If the client is rejected by this forward or any other, OnClientDisconnect will not be called.<br>Note: Do not write to rejectmsg if you plan on returning true. If multiple plugins write to the string buffer, it is not defined which plugin's string will be shown to the client, but it is guaranteed one of them will. */
  export type OnClientConnectCallback = (number playerSlot, string name, string networkId) => boolean;

  /** Called on client connection. */
  export type OnClientConnect_PostCallback = (number playerSlot) => void;

  /** Called once a client successfully connects. This callback is paired with OnClientDisconnect. */
  export type OnClientConnectedCallback = (number playerSlot) => void;

  /** Called when a client is entering the game. */
  export type OnClientPutInServerCallback = (number playerSlot) => void;

  /** Called when a client is disconnecting from the server. */
  export type OnClientDisconnectCallback = (number playerSlot) => void;

  /** Called when a client is disconnected from the server. */
  export type OnClientDisconnect_PostCallback = (number playerSlot, number reason) => void;

  /** Called when a client is activated by the game. */
  export type OnClientActiveCallback = (number playerSlot, boolean isActive) => void;

  /** Called when a client is fully connected to the game. */
  export type OnClientFullyConnectCallback = (number playerSlot) => void;

  /** Called whenever the client's settings are changed. */
  export type OnClientSettingsChangedCallback = (number playerSlot) => void;

  /** Called when a client is fully connected to the game. */
  export type OnClientAuthenticatedCallback = (number playerSlot, bigint steamID) => void;

  /** Called right before a round terminates. */
  export type OnRoundTerminatedCallback = (number delay, number reason) => void;

  /** Called when an entity is created. */
  export type OnEntityCreatedCallback = (number entityHandle) => void;

  /** Called when when an entity is destroyed. */
  export type OnEntityDeletedCallback = (number entityHandle) => void;

  /** When an entity is reparented to another entity. */
  export type OnEntityParentChangedCallback = (number entityHandle, number parentHandle) => void;

  /** Called on every server startup. */
  export type OnServerStartupCallback = () => void;

  /** Called on every server activate. */
  export type OnServerActivateCallback = () => void;

  /** Called on every server spawn. */
  export type OnServerSpawnCallback = () => void;

  /** Called on every server started only once. */
  export type OnServerStartedCallback = () => void;

  /** Called on every map start. */
  export type OnMapStartCallback = () => void;

  /** Called on every map end. */
  export type OnMapEndCallback = () => void;

  /** Called before every server frame. Note that you should avoid doing expensive computations or declaring large local arrays. */
  export type OnGameFrameCallback = (boolean simulating, boolean firstTick, boolean lastTick) => void;

  /** Called when the server is not in game. */
  export type OnUpdateWhenNotInGameCallback = (number deltaTime) => void;

  /** Called before every server frame, before entities are updated. */
  export type OnPreWorldUpdateCallback = (boolean simulating) => void;

  /** Enum representing the possible reasons for a round ending in Counter-Strike. */
  export const enum CSRoundEndReason {
    /** Target successfully bombed. */
    TargetBombed = 1,
    /** The VIP has escaped (not present in CS:GO). */
    VIPEscaped = 2,
    /** VIP has been assassinated (not present in CS:GO). */
    VIPKilled = 3,
    /** The terrorists have escaped. */
    TerroristsEscaped = 4,
    /** The CTs have prevented most of the terrorists from escaping. */
    CTStoppedEscape = 5,
    /** Escaping terrorists have all been neutralized. */
    TerroristsStopped = 6,
    /** The bomb has been defused. */
    BombDefused = 7,
    /** Counter-Terrorists win. */
    CTWin = 8,
    /** Terrorists win. */
    TerroristWin = 9,
    /** Round draw. */
    Draw = 10,
    /** All hostages have been rescued. */
    HostagesRescued = 11,
    /** Target has been saved. */
    TargetSaved = 12,
    /** Hostages have not been rescued. */
    HostagesNotRescued = 13,
    /** Terrorists have not escaped. */
    TerroristsNotEscaped = 14,
    /** VIP has not escaped (not present in CS:GO). */
    VIPNotEscaped = 15,
    /** Game commencing. */
    GameStart = 16,
    /** Terrorists surrender. */
    TerroristsSurrender = 17,
    /** CTs surrender. */
    CTSurrender = 18,
    /** Terrorists planted the bomb. */
    TerroristsPlanted = 19,
    /** CTs reached the hostage. */
    CTsReachedHostage = 20,
    /** Survival mode win. */
    SurvivalWin = 21,
    /** Survival mode draw. */
    SurvivalDraw = 22
  }

  /** Callback function for user messages. */
  export type UserMessageCallback = (bigint userMessage) => number;

  /** Enum representing different weapon types. */
  export const enum CSWeaponType {
    Knife = 0,
    Pistol = 1,
    SubmachineGun = 2,
    Rifle = 3,
    Shotgun = 4,
    SniperRifle = 5,
    MachineGun = 6,
    C4 = 7,
    Taser = 8,
    Grenade = 9,
    Equipment = 10,
    StackableItem = 11,
    Unknown = 12
  }

  /** Enum representing different weapon categories. */
  export const enum CSWeaponCategory {
    Other = 0,
    Melee = 1,
    Secondary = 2,
    SMG = 3,
    Rifle = 4,
    Heavy = 5,
    Count = 6
  }

  /** Enum representing different gear slots. */
  export const enum GearSlot {
    Invalid = 4294967295,
    Rifle = 0,
    Pistol = 1,
    Knife = 2,
    Grenades = 3,
    C4 = 4,
    ReservedSlot6 = 5,
    ReservedSlot7 = 6,
    ReservedSlot8 = 7,
    ReservedSlot9 = 8,
    ReservedSlot10 = 9,
    ReservedSlot11 = 10,
    Boosts = 11,
    Utility = 12,
    Count = 13,
    First = 0,
    Last = 12
  }

  /** Enum representing different weapon definition indices. */
  export const enum WeaponDefIndex {
    Invalid = 0,
    Deagle = 1,
    Elite = 2,
    FiveSeven = 3,
    Glock = 4,
    AK47 = 7,
    AUG = 8,
    AWP = 9,
    FAMAS = 10,
    G3SG1 = 11,
    GalilAR = 13,
    M249 = 14,
    M4A1 = 16,
    MAC10 = 17,
    P90 = 19,
    MP5SD = 23,
    UMP45 = 24,
    XM1014 = 25,
    Bizon = 26,
    MAG7 = 27,
    Negev = 28,
    SawedOff = 29,
    Tec9 = 30,
    Taser = 31,
    HKP2000 = 32,
    MP7 = 33,
    MP9 = 34,
    Nova = 35,
    P250 = 36,
    SCAR20 = 38,
    SG556 = 39,
    SSG08 = 40,
    KnifeGG = 41,
    Knife = 42,
    Flashbang = 43,
    HEGrenade = 44,
    SmokeGrenade = 45,
    Molotov = 46,
    Decoy = 47,
    IncGrenade = 48,
    C4 = 49,
    Kevlar = 50,
    AssaultSuit = 51,
    HeavyAssaultSuit = 52,
    Defuser = 55,
    KnifeT = 59,
    M4A1Silencer = 60,
    USPSilencer = 61,
    CZ75A = 63,
    Revolver = 64,
    Bayonet = 500,
    KnifeCSS = 503,
    KnifeFlip = 505,
    KnifeGut = 506,
    KnifeKarambit = 507,
    KnifeM9Bayonet = 508,
    KnifeTactical = 509,
    KnifeFalchion = 512,
    KnifeBowie = 514,
    KnifeButterfly = 515,
    KnifePush = 516,
    KnifeCord = 517,
    KnifeCanis = 518,
    KnifeUrsus = 519,
    KnifeGypsyJackknife = 520,
    KnifeOutdoor = 521,
    KnifeStiletto = 522,
    KnifeWidowmaker = 523,
    KnifeSkeleton = 525,
    KnifeKukri = 526
  }


  /**
   * Creates a new KeyValues instance
   * @param setName - The name to assign to this KeyValues instance
   * @returns Pointer to the newly created KeyValues object
   */
  export function Kv1Create(string setName): bigint;

  /**
   * Destroys a KeyValues instance
   * @param kv - Pointer to the KeyValues object to destroy
   */
  export function Kv1Destroy(bigint kv): void;

  /**
   * Gets the section name of a KeyValues instance
   * @param kv - Pointer to the KeyValues object
   * @returns The name of the KeyValues section
   */
  export function Kv1GetName(bigint kv): string;

  /**
   * Sets the section name of a KeyValues instance
   * @param kv - Pointer to the KeyValues object
   * @param name - The new name to assign to this KeyValues section
   */
  export function Kv1SetName(bigint kv, string name): void;

  /**
   * Finds a key by name
   * @param kv - Pointer to the KeyValues object to search in
   * @param keyName - The name of the key to find
   * @returns Pointer to the found KeyValues subkey, or NULL if not found
   */
  export function Kv1FindKey(bigint kv, string keyName): bigint;

  /**
   * Finds a key by name or creates it if it doesn't exist
   * @param kv - Pointer to the KeyValues object to search in
   * @param keyName - The name of the key to find or create
   * @returns Pointer to the found or newly created KeyValues subkey (never NULL)
   */
  export function Kv1FindKeyOrCreate(bigint kv, string keyName): bigint;

  /**
   * Creates a new subkey with the specified name
   * @param kv - Pointer to the parent KeyValues object
   * @param keyName - The name for the new key
   * @returns Pointer to the newly created KeyValues subkey
   */
  export function Kv1CreateKey(bigint kv, string keyName): bigint;

  /**
   * Creates a new subkey with an autogenerated name
   * @param kv - Pointer to the parent KeyValues object
   * @returns Pointer to the newly created KeyValues subkey
   */
  export function Kv1CreateNewKey(bigint kv): bigint;

  /**
   * Adds a subkey to this KeyValues instance
   * @param kv - Pointer to the parent KeyValues object
   * @param subKey - Pointer to the KeyValues object to add as a child
   */
  export function Kv1AddSubKey(bigint kv, bigint subKey): void;

  /**
   * Gets the first subkey in the list
   * @param kv - Pointer to the parent KeyValues object
   * @returns Pointer to the first subkey, or NULL if there are no children
   */
  export function Kv1GetFirstSubKey(bigint kv): bigint;

  /**
   * Gets the next sibling key in the list
   * @param kv - Pointer to the current KeyValues object
   * @returns Pointer to the next sibling key, or NULL if this is the last sibling
   */
  export function Kv1GetNextKey(bigint kv): bigint;

  /**
   * Gets a color value from a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to retrieve the color from
   * @param defaultValue - The default color value to return if the key is not found
   * @returns The color value as a 32-bit integer (RGBA)
   */
  export function Kv1GetColor(bigint kv, string keyName, number defaultValue): number;

  /**
   * Sets a color value for a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to set the color for
   * @param value - The color value as a 32-bit integer (RGBA)
   */
  export function Kv1SetColor(bigint kv, string keyName, number value): void;

  /**
   * Gets an integer value from a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to retrieve the integer from
   * @param defaultValue - The default value to return if the key is not found
   * @returns The integer value associated with the key, or defaultValue if not found
   */
  export function Kv1GetInt(bigint kv, string keyName, number defaultValue): number;

  /**
   * Sets an integer value for a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to set the integer for
   * @param value - The integer value to set
   */
  export function Kv1SetInt(bigint kv, string keyName, number value): void;

  /**
   * Gets a float value from a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to retrieve the float from
   * @param defaultValue - The default value to return if the key is not found
   * @returns The float value associated with the key, or defaultValue if not found
   */
  export function Kv1GetFloat(bigint kv, string keyName, number defaultValue): number;

  /**
   * Sets a float value for a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to set the float for
   * @param value - The float value to set
   */
  export function Kv1SetFloat(bigint kv, string keyName, number value): void;

  /**
   * Gets a string value from a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to retrieve the string from
   * @param defaultValue - The default string to return if the key is not found
   * @returns The string value associated with the key, or defaultValue if not found
   */
  export function Kv1GetString(bigint kv, string keyName, string defaultValue): string;

  /**
   * Sets a string value for a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to set the string for
   * @param value - The string value to set
   */
  export function Kv1SetString(bigint kv, string keyName, string value): void;

  /**
   * Gets a pointer value from a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to retrieve the pointer from
   * @param defaultValue - The default pointer to return if the key is not found
   * @returns The pointer value associated with the key, or defaultValue if not found
   */
  export function Kv1GetPtr(bigint kv, string keyName, bigint defaultValue): bigint;

  /**
   * Sets a pointer value for a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to set the pointer for
   * @param value - The pointer value to set
   */
  export function Kv1SetPtr(bigint kv, string keyName, bigint value): void;

  /**
   * Gets a boolean value from a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to retrieve the boolean from
   * @param defaultValue - The default value to return if the key is not found
   * @returns The boolean value associated with the key, or defaultValue if not found
   */
  export function Kv1GetBool(bigint kv, string keyName, boolean defaultValue): boolean;

  /**
   * Sets a boolean value for a key
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to set the boolean for
   * @param value - The boolean value to set
   */
  export function Kv1SetBool(bigint kv, string keyName, boolean value): void;

  /**
   * Makes a deep copy of a KeyValues tree
   * @param kv - Pointer to the KeyValues object to copy
   * @returns Pointer to the newly allocated copy of the KeyValues tree
   */
  export function Kv1MakeCopy(bigint kv): bigint;

  /**
   * Clears all subkeys and the current value
   * @param kv - Pointer to the KeyValues object to clear
   */
  export function Kv1Clear(bigint kv): void;

  /**
   * Checks if a key exists and has no value or subkeys
   * @param kv - Pointer to the KeyValues object
   * @param keyName - The name of the key to check
   * @returns true if the key exists and is empty, false otherwise
   */
  export function Kv1IsEmpty(bigint kv, string keyName): boolean;

  /**
   * Creates a new KeyValues3 object with specified type and subtype
   * @param type - The KV3 type enumeration value
   * @param subtype - The KV3 subtype enumeration value
   * @returns Pointer to the newly created KeyValues3 object
   */
  export function Kv3Create(number type, number subtype): bigint;

  /**
   * Creates a new KeyValues3 object with cluster element, type, and subtype
   * @param cluster_elem - The cluster element index
   * @param type - The KV3 type enumeration value
   * @param subtype - The KV3 subtype enumeration value
   * @returns Pointer to the newly created KeyValues3 object
   */
  export function Kv3CreateWithCluster(number cluster_elem, number type, number subtype): bigint;

  /**
   * Creates a copy of an existing KeyValues3 object
   * @param other - Pointer to the KeyValues3 object to copy
   * @returns Pointer to the newly created copy, or nullptr if other is null
   */
  export function Kv3CreateCopy(bigint other): bigint;

  /**
   * Destroys a KeyValues3 object and frees its memory
   * @param kv - Pointer to the KeyValues3 object to destroy
   */
  export function Kv3Destroy(bigint kv): void;

  /**
   * Copies data from another KeyValues3 object
   * @param kv - Pointer to the destination KeyValues3 object
   * @param other - Pointer to the source KeyValues3 object
   */
  export function Kv3CopyFrom(bigint kv, bigint other): void;

  /**
   * Overlays keys from another KeyValues3 object
   * @param kv - Pointer to the destination KeyValues3 object
   * @param other - Pointer to the source KeyValues3 object
   * @param depth - Whether to perform a deep overlay
   */
  export function Kv3OverlayKeysFrom(bigint kv, bigint other, boolean depth): void;

  /**
   * Gets the context associated with a KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @returns Pointer to the CKeyValues3Context, or nullptr if kv is null
   */
  export function Kv3GetContext(bigint kv): bigint;

  /**
   * Gets the metadata associated with a KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param ppCtx - Pointer to store the context pointer
   * @returns Pointer to the KV3MetaData_t structure, or nullptr if kv is null
   */
  export function Kv3GetMetaData(bigint kv, bigint ppCtx): bigint;

  /**
   * Checks if a specific flag is set
   * @param kv - Pointer to the KeyValues3 object
   * @param flag - The flag to check
   * @returns true if the flag is set, false otherwise
   */
  export function Kv3HasFlag(bigint kv, number flag): boolean;

  /**
   * Checks if any flags are set
   * @param kv - Pointer to the KeyValues3 object
   * @returns true if any flags are set, false otherwise
   */
  export function Kv3HasAnyFlags(bigint kv): boolean;

  /**
   * Gets all flags as a bitmask
   * @param kv - Pointer to the KeyValues3 object
   * @returns Bitmask of all flags, or 0 if kv is null
   */
  export function Kv3GetAllFlags(bigint kv): number;

  /**
   * Sets all flags from a bitmask
   * @param kv - Pointer to the KeyValues3 object
   * @param flags - Bitmask of flags to set
   */
  export function Kv3SetAllFlags(bigint kv, number flags): void;

  /**
   * Sets or clears a specific flag
   * @param kv - Pointer to the KeyValues3 object
   * @param flag - The flag to modify
   * @param state - true to set the flag, false to clear it
   */
  export function Kv3SetFlag(bigint kv, number flag, boolean state): void;

  /**
   * Gets the basic type of the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @returns The type enumeration value, or 0 if kv is null
   */
  export function Kv3GetType(bigint kv): number;

  /**
   * Gets the extended type of the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @returns The extended type enumeration value, or 0 if kv is null
   */
  export function Kv3GetTypeEx(bigint kv): number;

  /**
   * Gets the subtype of the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @returns The subtype enumeration value, or 0 if kv is null
   */
  export function Kv3GetSubType(bigint kv): number;

  /**
   * Checks if the object has invalid member names
   * @param kv - Pointer to the KeyValues3 object
   * @returns true if invalid member names exist, false otherwise
   */
  export function Kv3HasInvalidMemberNames(bigint kv): boolean;

  /**
   * Sets the invalid member names flag
   * @param kv - Pointer to the KeyValues3 object
   * @param bValue - true to mark as having invalid member names, false otherwise
   */
  export function Kv3SetHasInvalidMemberNames(bigint kv, boolean bValue): void;

  /**
   * Gets the type as a string representation
   * @param kv - Pointer to the KeyValues3 object
   * @returns String representation of the type, or empty string if kv is null
   */
  export function Kv3GetTypeAsString(bigint kv): string;

  /**
   * Gets the subtype as a string representation
   * @param kv - Pointer to the KeyValues3 object
   * @returns String representation of the subtype, or empty string if kv is null
   */
  export function Kv3GetSubTypeAsString(bigint kv): string;

  /**
   * Converts the KeyValues3 object to a string representation
   * @param kv - Pointer to the KeyValues3 object
   * @param flags - Formatting flags for the string conversion
   * @returns String representation of the object, or empty string if kv is null
   */
  export function Kv3ToString(bigint kv, number flags): string;

  /**
   * Checks if the KeyValues3 object is null
   * @param kv - Pointer to the KeyValues3 object
   * @returns true if the object is null or the pointer is null, false otherwise
   */
  export function Kv3IsNull(bigint kv): boolean;

  /**
   * Sets the KeyValues3 object to null
   * @param kv - Pointer to the KeyValues3 object
   */
  export function Kv3SetToNull(bigint kv): void;

  /**
   * Checks if the KeyValues3 object is an array
   * @param kv - Pointer to the KeyValues3 object
   * @returns true if the object is an array, false otherwise
   */
  export function Kv3IsArray(bigint kv): boolean;

  /**
   * Checks if the KeyValues3 object is a KV3 array
   * @param kv - Pointer to the KeyValues3 object
   * @returns true if the object is a KV3 array, false otherwise
   */
  export function Kv3IsKV3Array(bigint kv): boolean;

  /**
   * Checks if the KeyValues3 object is a table
   * @param kv - Pointer to the KeyValues3 object
   * @returns true if the object is a table, false otherwise
   */
  export function Kv3IsTable(bigint kv): boolean;

  /**
   * Checks if the KeyValues3 object is a string
   * @param kv - Pointer to the KeyValues3 object
   * @returns true if the object is a string, false otherwise
   */
  export function Kv3IsString(bigint kv): boolean;

  /**
   * Gets the boolean value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns Boolean value or defaultValue
   */
  export function Kv3GetBool(bigint kv, boolean defaultValue): boolean;

  /**
   * Gets the char value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns Char value or defaultValue
   */
  export function Kv3GetChar(bigint kv, number defaultValue): number;

  /**
   * Gets the 32-bit Unicode character value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns 32-bit Unicode character value or defaultValue
   */
  export function Kv3GetUChar32(bigint kv, number defaultValue): number;

  /**
   * Gets the signed 8-bit integer value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns int8_t value or defaultValue
   */
  export function Kv3GetInt8(bigint kv, number defaultValue): number;

  /**
   * Gets the unsigned 8-bit integer value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns uint8_t value or defaultValue
   */
  export function Kv3GetUInt8(bigint kv, number defaultValue): number;

  /**
   * Gets the signed 16-bit integer value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns int16_t value or defaultValue
   */
  export function Kv3GetShort(bigint kv, number defaultValue): number;

  /**
   * Gets the unsigned 16-bit integer value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns uint16_t value or defaultValue
   */
  export function Kv3GetUShort(bigint kv, number defaultValue): number;

  /**
   * Gets the signed 32-bit integer value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns int32_t value or defaultValue
   */
  export function Kv3GetInt(bigint kv, number defaultValue): number;

  /**
   * Gets the unsigned 32-bit integer value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns uint32_t value or defaultValue
   */
  export function Kv3GetUInt(bigint kv, number defaultValue): number;

  /**
   * Gets the signed 64-bit integer value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns int64_t value or defaultValue
   */
  export function Kv3GetInt64(bigint kv, number defaultValue): number;

  /**
   * Gets the unsigned 64-bit integer value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns uint64_t value or defaultValue
   */
  export function Kv3GetUInt64(bigint kv, bigint defaultValue): bigint;

  /**
   * Gets the float value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns Float value or defaultValue
   */
  export function Kv3GetFloat(bigint kv, number defaultValue): number;

  /**
   * Gets the double value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null or conversion fails
   * @returns Double value or defaultValue
   */
  export function Kv3GetDouble(bigint kv, number defaultValue): number;

  /**
   * Sets the KeyValues3 object to a boolean value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - Boolean value to set
   */
  export function Kv3SetBool(bigint kv, boolean value): void;

  /**
   * Sets the KeyValues3 object to a char value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - Char value to set
   */
  export function Kv3SetChar(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to a 32-bit Unicode character value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - 32-bit Unicode character value to set
   */
  export function Kv3SetUChar32(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to a signed 8-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - int8_t value to set
   */
  export function Kv3SetInt8(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to an unsigned 8-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - uint8_t value to set
   */
  export function Kv3SetUInt8(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to a signed 16-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - int16_t value to set
   */
  export function Kv3SetShort(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to an unsigned 16-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - uint16_t value to set
   */
  export function Kv3SetUShort(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to a signed 32-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - int32_t value to set
   */
  export function Kv3SetInt(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to an unsigned 32-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - uint32_t value to set
   */
  export function Kv3SetUInt(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to a signed 64-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - int64_t value to set
   */
  export function Kv3SetInt64(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to an unsigned 64-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - uint64_t value to set
   */
  export function Kv3SetUInt64(bigint kv, bigint value): void;

  /**
   * Sets the KeyValues3 object to a float value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - Float value to set
   */
  export function Kv3SetFloat(bigint kv, number value): void;

  /**
   * Sets the KeyValues3 object to a double value
   * @param kv - Pointer to the KeyValues3 object
   * @param value - Double value to set
   */
  export function Kv3SetDouble(bigint kv, number value): void;

  /**
   * Gets the pointer value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default value to return if kv is null
   * @returns Pointer value as uintptr_t or defaultValue
   */
  export function Kv3GetPointer(bigint kv, bigint defaultValue): bigint;

  /**
   * Sets the KeyValues3 object to a pointer value
   * @param kv - Pointer to the KeyValues3 object
   * @param ptr - Pointer value as uintptr_t to set
   */
  export function Kv3SetPointer(bigint kv, bigint ptr): void;

  /**
   * Gets the string token value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default token value to return if kv is null
   * @returns String token hash code or defaultValue
   */
  export function Kv3GetStringToken(bigint kv, number defaultValue): number;

  /**
   * Sets the KeyValues3 object to a string token value
   * @param kv - Pointer to the KeyValues3 object
   * @param token - String token hash code to set
   */
  export function Kv3SetStringToken(bigint kv, number token): void;

  /**
   * Gets the entity handle value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default entity handle value to return if kv is null
   * @returns Entity handle as int32_t or defaultValue
   */
  export function Kv3GetEHandle(bigint kv, number defaultValue): number;

  /**
   * Sets the KeyValues3 object to an entity handle value
   * @param kv - Pointer to the KeyValues3 object
   * @param ehandle - Entity handle value to set
   */
  export function Kv3SetEHandle(bigint kv, number ehandle): void;

  /**
   * Gets the string value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default string to return if kv is null or value is empty
   * @returns String value or defaultValue
   */
  export function Kv3GetString(bigint kv, string defaultValue): string;

  /**
   * Sets the KeyValues3 object to a string value (copies the string)
   * @param kv - Pointer to the KeyValues3 object
   * @param str - String value to set
   * @param subtype - String subtype enumeration value
   */
  export function Kv3SetString(bigint kv, string str, number subtype): void;

  /**
   * Sets the KeyValues3 object to an external string value (does not copy)
   * @param kv - Pointer to the KeyValues3 object
   * @param str - External string value to reference
   * @param subtype - String subtype enumeration value
   */
  export function Kv3SetStringExternal(bigint kv, string str, number subtype): void;

  /**
   * Gets the binary blob from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @returns Vector containing the binary blob data, or empty vector if kv is null
   */
  export function Kv3GetBinaryBlob(bigint kv): number[];

  /**
   * Gets the size of the binary blob in the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @returns Size of the binary blob in bytes, or 0 if kv is null
   */
  export function Kv3GetBinaryBlobSize(bigint kv): number;

  /**
   * Sets the KeyValues3 object to a binary blob (copies the data)
   * @param kv - Pointer to the KeyValues3 object
   * @param blob - Vector containing the binary blob data
   */
  export function Kv3SetToBinaryBlob(bigint kv, number[] blob): void;

  /**
   * Sets the KeyValues3 object to an external binary blob (does not copy)
   * @param kv - Pointer to the KeyValues3 object
   * @param blob - Vector containing the external binary blob data
   * @param free_mem - Whether to free the memory when the object is destroyed
   */
  export function Kv3SetToBinaryBlobExternal(bigint kv, number[] blob, boolean free_mem): void;

  /**
   * Gets the color value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default color value to return if kv is null
   * @returns Color value as int32_t or defaultValue
   */
  export function Kv3GetColor(bigint kv, number defaultValue): number;

  /**
   * Sets the KeyValues3 object to a color value
   * @param kv - Pointer to the KeyValues3 object
   * @param color - Color value as int32_t to set
   */
  export function Kv3SetColor(bigint kv, number color): void;

  /**
   * Gets the 3D vector value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default vector to return if kv is null
   * @returns 3D vector or defaultValue
   */
  export function Kv3GetVector(bigint kv, Vector3 defaultValue): Vector3;

  /**
   * Gets the 2D vector value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default 2D vector to return if kv is null
   * @returns 2D vector or defaultValue
   */
  export function Kv3GetVector2D(bigint kv, Vector2 defaultValue): Vector2;

  /**
   * Gets the 4D vector value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default 4D vector to return if kv is null
   * @returns 4D vector or defaultValue
   */
  export function Kv3GetVector4D(bigint kv, Vector4 defaultValue): Vector4;

  /**
   * Gets the quaternion value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default quaternion to return if kv is null
   * @returns Quaternion as vec4 or defaultValue
   */
  export function Kv3GetQuaternion(bigint kv, Vector4 defaultValue): Vector4;

  /**
   * Gets the angle (QAngle) value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default angle to return if kv is null
   * @returns QAngle as vec3 or defaultValue
   */
  export function Kv3GetQAngle(bigint kv, Vector3 defaultValue): Vector3;

  /**
   * Gets the 3x4 matrix value from the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   * @param defaultValue - Default matrix to return if kv is null
   * @returns 3x4 matrix as mat4x4 or defaultValue
   */
  export function Kv3GetMatrix3x4(bigint kv, Matrix4x4 defaultValue): Matrix4x4;

  /**
   * Sets the KeyValues3 object to a 3D vector value
   * @param kv - Pointer to the KeyValues3 object
   * @param vec - 3D vector to set
   */
  export function Kv3SetVector(bigint kv, Vector3 vec): void;

  /**
   * Sets the KeyValues3 object to a 2D vector value
   * @param kv - Pointer to the KeyValues3 object
   * @param vec2d - 2D vector to set
   */
  export function Kv3SetVector2D(bigint kv, Vector2 vec2d): void;

  /**
   * Sets the KeyValues3 object to a 4D vector value
   * @param kv - Pointer to the KeyValues3 object
   * @param vec4d - 4D vector to set
   */
  export function Kv3SetVector4D(bigint kv, Vector4 vec4d): void;

  /**
   * Sets the KeyValues3 object to a quaternion value
   * @param kv - Pointer to the KeyValues3 object
   * @param quat - Quaternion to set (as vec4)
   */
  export function Kv3SetQuaternion(bigint kv, Vector4 quat): void;

  /**
   * Sets the KeyValues3 object to an angle (QAngle) value
   * @param kv - Pointer to the KeyValues3 object
   * @param ang - QAngle to set (as vec3)
   */
  export function Kv3SetQAngle(bigint kv, Vector3 ang): void;

  /**
   * Sets the KeyValues3 object to a 3x4 matrix value
   * @param kv - Pointer to the KeyValues3 object
   * @param matrix - 3x4 matrix to set (as mat4x4)
   */
  export function Kv3SetMatrix3x4(bigint kv, Matrix4x4 matrix): void;

  /**
   * Gets the number of elements in the array
   * @param kv - Pointer to the KeyValues3 object
   * @returns Number of array elements, or 0 if kv is null or not an array
   */
  export function Kv3GetArrayElementCount(bigint kv): number;

  /**
   * Sets the number of elements in the array
   * @param kv - Pointer to the KeyValues3 object
   * @param count - Number of elements to set
   * @param type - Type of array elements
   * @param subtype - Subtype of array elements
   */
  export function Kv3SetArrayElementCount(bigint kv, number count, number type, number subtype): void;

  /**
   * Sets the KeyValues3 object to an empty KV3 array
   * @param kv - Pointer to the KeyValues3 object
   */
  export function Kv3SetToEmptyKV3Array(bigint kv): void;

  /**
   * Gets an array element at the specified index
   * @param kv - Pointer to the KeyValues3 object
   * @param elem - Index of the element to get
   * @returns Pointer to the element KeyValues3 object, or nullptr if invalid
   */
  export function Kv3GetArrayElement(bigint kv, number elem): bigint;

  /**
   * Inserts a new element before the specified index
   * @param kv - Pointer to the KeyValues3 object
   * @param elem - Index before which to insert
   * @returns Pointer to the newly inserted element, or nullptr if invalid
   */
  export function Kv3ArrayInsertElementBefore(bigint kv, number elem): bigint;

  /**
   * Inserts a new element after the specified index
   * @param kv - Pointer to the KeyValues3 object
   * @param elem - Index after which to insert
   * @returns Pointer to the newly inserted element, or nullptr if invalid
   */
  export function Kv3ArrayInsertElementAfter(bigint kv, number elem): bigint;

  /**
   * Adds a new element to the end of the array
   * @param kv - Pointer to the KeyValues3 object
   * @returns Pointer to the newly added element, or nullptr if invalid
   */
  export function Kv3ArrayAddElementToTail(bigint kv): bigint;

  /**
   * Swaps two array elements
   * @param kv - Pointer to the KeyValues3 object
   * @param idx1 - Index of the first element
   * @param idx2 - Index of the second element
   */
  export function Kv3ArraySwapItems(bigint kv, number idx1, number idx2): void;

  /**
   * Removes an element from the array
   * @param kv - Pointer to the KeyValues3 object
   * @param elem - Index of the element to remove
   */
  export function Kv3ArrayRemoveElement(bigint kv, number elem): void;

  /**
   * Sets the KeyValues3 object to an empty table
   * @param kv - Pointer to the KeyValues3 object
   */
  export function Kv3SetToEmptyTable(bigint kv): void;

  /**
   * Gets the number of members in the table
   * @param kv - Pointer to the KeyValues3 object
   * @returns Number of table members, or 0 if kv is null or not a table
   */
  export function Kv3GetMemberCount(bigint kv): number;

  /**
   * Checks if a member with the specified name exists
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member to check
   * @returns true if the member exists, false otherwise
   */
  export function Kv3HasMember(bigint kv, string name): boolean;

  /**
   * Finds a member by name
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member to find
   * @returns Pointer to the member KeyValues3 object, or nullptr if not found
   */
  export function Kv3FindMember(bigint kv, string name): bigint;

  /**
   * Finds a member by name, or creates it if it doesn't exist
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member to find or create
   * @returns Pointer to the member KeyValues3 object, or nullptr if kv is null
   */
  export function Kv3FindOrCreateMember(bigint kv, string name): bigint;

  /**
   * Removes a member from the table
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member to remove
   * @returns true if the member was removed, false otherwise
   */
  export function Kv3RemoveMember(bigint kv, string name): boolean;

  /**
   * Gets the name of a member at the specified index
   * @param kv - Pointer to the KeyValues3 object
   * @param index - Index of the member
   * @returns Name of the member, or empty string if invalid
   */
  export function Kv3GetMemberName(bigint kv, number index): string;

  /**
   * Gets a member by index
   * @param kv - Pointer to the KeyValues3 object
   * @param index - Index of the member to get
   * @returns Pointer to the member KeyValues3 object, or nullptr if invalid
   */
  export function Kv3GetMemberByIndex(bigint kv, number index): bigint;

  /**
   * Gets a boolean value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns Boolean value or defaultValue
   */
  export function Kv3GetMemberBool(bigint kv, string name, boolean defaultValue): boolean;

  /**
   * Gets a char value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns Char value or defaultValue
   */
  export function Kv3GetMemberChar(bigint kv, string name, number defaultValue): number;

  /**
   * Gets a 32-bit Unicode character value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns 32-bit Unicode character value or defaultValue
   */
  export function Kv3GetMemberUChar32(bigint kv, string name, number defaultValue): number;

  /**
   * Gets a signed 8-bit integer value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns int8_t value or defaultValue
   */
  export function Kv3GetMemberInt8(bigint kv, string name, number defaultValue): number;

  /**
   * Gets an unsigned 8-bit integer value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns uint8_t value or defaultValue
   */
  export function Kv3GetMemberUInt8(bigint kv, string name, number defaultValue): number;

  /**
   * Gets a signed 16-bit integer value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns int16_t value or defaultValue
   */
  export function Kv3GetMemberShort(bigint kv, string name, number defaultValue): number;

  /**
   * Gets an unsigned 16-bit integer value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns uint16_t value or defaultValue
   */
  export function Kv3GetMemberUShort(bigint kv, string name, number defaultValue): number;

  /**
   * Gets a signed 32-bit integer value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns int32_t value or defaultValue
   */
  export function Kv3GetMemberInt(bigint kv, string name, number defaultValue): number;

  /**
   * Gets an unsigned 32-bit integer value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns uint32_t value or defaultValue
   */
  export function Kv3GetMemberUInt(bigint kv, string name, number defaultValue): number;

  /**
   * Gets a signed 64-bit integer value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns int64_t value or defaultValue
   */
  export function Kv3GetMemberInt64(bigint kv, string name, number defaultValue): number;

  /**
   * Gets an unsigned 64-bit integer value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns uint64_t value or defaultValue
   */
  export function Kv3GetMemberUInt64(bigint kv, string name, bigint defaultValue): bigint;

  /**
   * Gets a float value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns Float value or defaultValue
   */
  export function Kv3GetMemberFloat(bigint kv, string name, number defaultValue): number;

  /**
   * Gets a double value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns Double value or defaultValue
   */
  export function Kv3GetMemberDouble(bigint kv, string name, number defaultValue): number;

  /**
   * Gets a pointer value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default value to return if member not found
   * @returns Pointer value as uintptr_t or defaultValue
   */
  export function Kv3GetMemberPointer(bigint kv, string name, bigint defaultValue): bigint;

  /**
   * Gets a string token value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default token value to return if member not found
   * @returns String token hash code or defaultValue
   */
  export function Kv3GetMemberStringToken(bigint kv, string name, number defaultValue): number;

  /**
   * Gets an entity handle value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default entity handle value to return if member not found
   * @returns Entity handle as int32_t or defaultValue
   */
  export function Kv3GetMemberEHandle(bigint kv, string name, number defaultValue): number;

  /**
   * Gets a string value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default string to return if member not found
   * @returns String value or defaultValue
   */
  export function Kv3GetMemberString(bigint kv, string name, string defaultValue): string;

  /**
   * Gets a color value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default color value to return if member not found
   * @returns Color value as int32_t or defaultValue
   */
  export function Kv3GetMemberColor(bigint kv, string name, number defaultValue): number;

  /**
   * Gets a 3D vector value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default vector to return if member not found
   * @returns 3D vector or defaultValue
   */
  export function Kv3GetMemberVector(bigint kv, string name, Vector3 defaultValue): Vector3;

  /**
   * Gets a 2D vector value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default 2D vector to return if member not found
   * @returns 2D vector or defaultValue
   */
  export function Kv3GetMemberVector2D(bigint kv, string name, Vector2 defaultValue): Vector2;

  /**
   * Gets a 4D vector value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default 4D vector to return if member not found
   * @returns 4D vector or defaultValue
   */
  export function Kv3GetMemberVector4D(bigint kv, string name, Vector4 defaultValue): Vector4;

  /**
   * Gets a quaternion value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default quaternion to return if member not found
   * @returns Quaternion as vec4 or defaultValue
   */
  export function Kv3GetMemberQuaternion(bigint kv, string name, Vector4 defaultValue): Vector4;

  /**
   * Gets an angle (QAngle) value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default angle to return if member not found
   * @returns QAngle as vec3 or defaultValue
   */
  export function Kv3GetMemberQAngle(bigint kv, string name, Vector3 defaultValue): Vector3;

  /**
   * Gets a 3x4 matrix value from a table member
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param defaultValue - Default matrix to return if member not found
   * @returns 3x4 matrix as mat4x4 or defaultValue
   */
  export function Kv3GetMemberMatrix3x4(bigint kv, string name, Matrix4x4 defaultValue): Matrix4x4;

  /**
   * Sets a table member to null
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   */
  export function Kv3SetMemberToNull(bigint kv, string name): void;

  /**
   * Sets a table member to an empty array
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   */
  export function Kv3SetMemberToEmptyArray(bigint kv, string name): void;

  /**
   * Sets a table member to an empty table
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   */
  export function Kv3SetMemberToEmptyTable(bigint kv, string name): void;

  /**
   * Sets a table member to a binary blob (copies the data)
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param blob - Vector containing the binary blob data
   */
  export function Kv3SetMemberToBinaryBlob(bigint kv, string name, number[] blob): void;

  /**
   * Sets a table member to an external binary blob (does not copy)
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param blob - Vector containing the external binary blob data
   * @param free_mem - Whether to free the memory when the object is destroyed
   */
  export function Kv3SetMemberToBinaryBlobExternal(bigint kv, string name, number[] blob, boolean free_mem): void;

  /**
   * Sets a table member to a copy of another KeyValues3 value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param other - Pointer to the KeyValues3 object to copy
   */
  export function Kv3SetMemberToCopyOfValue(bigint kv, string name, bigint other): void;

  /**
   * Sets a table member to a boolean value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - Boolean value to set
   */
  export function Kv3SetMemberBool(bigint kv, string name, boolean value): void;

  /**
   * Sets a table member to a char value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - Char value to set
   */
  export function Kv3SetMemberChar(bigint kv, string name, number value): void;

  /**
   * Sets a table member to a 32-bit Unicode character value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - 32-bit Unicode character value to set
   */
  export function Kv3SetMemberUChar32(bigint kv, string name, number value): void;

  /**
   * Sets a table member to a signed 8-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - int8_t value to set
   */
  export function Kv3SetMemberInt8(bigint kv, string name, number value): void;

  /**
   * Sets a table member to an unsigned 8-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - uint8_t value to set
   */
  export function Kv3SetMemberUInt8(bigint kv, string name, number value): void;

  /**
   * Sets a table member to a signed 16-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - int16_t value to set
   */
  export function Kv3SetMemberShort(bigint kv, string name, number value): void;

  /**
   * Sets a table member to an unsigned 16-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - uint16_t value to set
   */
  export function Kv3SetMemberUShort(bigint kv, string name, number value): void;

  /**
   * Sets a table member to a signed 32-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - int32_t value to set
   */
  export function Kv3SetMemberInt(bigint kv, string name, number value): void;

  /**
   * Sets a table member to an unsigned 32-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - uint32_t value to set
   */
  export function Kv3SetMemberUInt(bigint kv, string name, number value): void;

  /**
   * Sets a table member to a signed 64-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - int64_t value to set
   */
  export function Kv3SetMemberInt64(bigint kv, string name, number value): void;

  /**
   * Sets a table member to an unsigned 64-bit integer value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - uint64_t value to set
   */
  export function Kv3SetMemberUInt64(bigint kv, string name, bigint value): void;

  /**
   * Sets a table member to a float value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - Float value to set
   */
  export function Kv3SetMemberFloat(bigint kv, string name, number value): void;

  /**
   * Sets a table member to a double value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param value - Double value to set
   */
  export function Kv3SetMemberDouble(bigint kv, string name, number value): void;

  /**
   * Sets a table member to a pointer value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param ptr - Pointer value as uintptr_t to set
   */
  export function Kv3SetMemberPointer(bigint kv, string name, bigint ptr): void;

  /**
   * Sets a table member to a string token value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param token - String token hash code to set
   */
  export function Kv3SetMemberStringToken(bigint kv, string name, number token): void;

  /**
   * Sets a table member to an entity handle value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param ehandle - Entity handle value to set
   */
  export function Kv3SetMemberEHandle(bigint kv, string name, number ehandle): void;

  /**
   * Sets a table member to a string value (copies the string)
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param str - String value to set
   * @param subtype - String subtype enumeration value
   */
  export function Kv3SetMemberString(bigint kv, string name, string str, number subtype): void;

  /**
   * Sets a table member to an external string value (does not copy)
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param str - External string value to reference
   * @param subtype - String subtype enumeration value
   */
  export function Kv3SetMemberStringExternal(bigint kv, string name, string str, number subtype): void;

  /**
   * Sets a table member to a color value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param color - Color value as int32_t to set
   */
  export function Kv3SetMemberColor(bigint kv, string name, number color): void;

  /**
   * Sets a table member to a 3D vector value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param vec - 3D vector to set
   */
  export function Kv3SetMemberVector(bigint kv, string name, Vector3 vec): void;

  /**
   * Sets a table member to a 2D vector value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param vec2d - 2D vector to set
   */
  export function Kv3SetMemberVector2D(bigint kv, string name, Vector2 vec2d): void;

  /**
   * Sets a table member to a 4D vector value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param vec4d - 4D vector to set
   */
  export function Kv3SetMemberVector4D(bigint kv, string name, Vector4 vec4d): void;

  /**
   * Sets a table member to a quaternion value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param quat - Quaternion to set (as vec4)
   */
  export function Kv3SetMemberQuaternion(bigint kv, string name, Vector4 quat): void;

  /**
   * Sets a table member to an angle (QAngle) value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param ang - QAngle to set (as vec3)
   */
  export function Kv3SetMemberQAngle(bigint kv, string name, Vector3 ang): void;

  /**
   * Sets a table member to a 3x4 matrix value
   * @param kv - Pointer to the KeyValues3 object
   * @param name - Name of the member
   * @param matrix - 3x4 matrix to set (as mat4x4)
   */
  export function Kv3SetMemberMatrix3x4(bigint kv, string name, Matrix4x4 matrix): void;

  /**
   * Prints debug information about the KeyValues3 object
   * @param kv - Pointer to the KeyValues3 object
   */
  export function Kv3DebugPrint(bigint kv): void;

  /**
   * Loads KeyValues3 data from a buffer into a context
   * @param context - Pointer to the KeyValues3 context
   * @param error - Output string for error messages
   * @param input - Vector containing the input buffer data
   * @param kv_name - Name for the KeyValues3 object
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromBuffer(bigint context, string error, number[] input, string kv_name, number flags): boolean;

  /**
   * Loads KeyValues3 data from a buffer
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param input - Vector containing the input buffer data
   * @param kv_name - Name for the KeyValues3 object
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3Load(bigint kv, string error, number[] input, string kv_name, number flags): boolean;

  /**
   * Loads KeyValues3 data from a text string
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param input - Text string containing KV3 data
   * @param kv_name - Name for the KeyValues3 object
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromText(bigint kv, string error, string input, string kv_name, number flags): boolean;

  /**
   * Loads KeyValues3 data from a file into a context
   * @param context - Pointer to the KeyValues3 context
   * @param error - Output string for error messages
   * @param filename - Name of the file to load
   * @param path - Path to the file
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromFileToContext(bigint context, string error, string filename, string path, number flags): boolean;

  /**
   * Loads KeyValues3 data from a file
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param filename - Name of the file to load
   * @param path - Path to the file
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromFile(bigint kv, string error, string filename, string path, number flags): boolean;

  /**
   * Loads KeyValues3 data from a JSON string
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param input - JSON string
   * @param kv_name - Name for the KeyValues3 object
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromJSON(bigint kv, string error, string input, string kv_name, number flags): boolean;

  /**
   * Loads KeyValues3 data from a JSON file
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param path - Path to the file
   * @param filename - Name of the file to load
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromJSONFile(bigint kv, string error, string path, string filename, number flags): boolean;

  /**
   * Loads KeyValues3 data from a KeyValues1 file
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param path - Path to the file
   * @param filename - Name of the file to load
   * @param esc_behavior - Escape sequence behavior for KV1 text
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromKV1File(bigint kv, string error, string path, string filename, number esc_behavior, number flags): boolean;

  /**
   * Loads KeyValues3 data from a KeyValues1 text string
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param input - KV1 text string
   * @param esc_behavior - Escape sequence behavior for KV1 text
   * @param kv_name - Name for the KeyValues3 object
   * @param unk - Unknown boolean parameter
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromKV1Text(bigint kv, string error, string input, number esc_behavior, string kv_name, boolean unk, number flags): boolean;

  /**
   * Loads KeyValues3 data from a KeyValues1 text string with translation
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param input - KV1 text string
   * @param esc_behavior - Escape sequence behavior for KV1 text
   * @param translation - Pointer to translation table
   * @param unk1 - Unknown integer parameter
   * @param kv_name - Name for the KeyValues3 object
   * @param unk2 - Unknown boolean parameter
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromKV1TextTranslated(bigint kv, string error, string input, number esc_behavior, bigint translation, number unk1, string kv_name, boolean unk2, number flags): boolean;

  /**
   * Loads data from a buffer that may be KV3 or KV1 format
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param input - Vector containing the input buffer data
   * @param kv_name - Name for the KeyValues3 object
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromKV3OrKV1(bigint kv, string error, number[] input, string kv_name, number flags): boolean;

  /**
   * Loads KeyValues3 data from old schema text format
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param input - Vector containing the input buffer data
   * @param kv_name - Name for the KeyValues3 object
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadFromOldSchemaText(bigint kv, string error, number[] input, string kv_name, number flags): boolean;

  /**
   * Loads KeyValues3 text without a header
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param input - Text string containing KV3 data
   * @param kv_name - Name for the KeyValues3 object
   * @param flags - Loading flags
   * @returns true if successful, false otherwise
   */
  export function Kv3LoadTextNoHeader(bigint kv, string error, string input, string kv_name, number flags): boolean;

  /**
   * Saves KeyValues3 data to a buffer
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param output - Vector to store the output buffer data
   * @param flags - Saving flags
   * @returns true if successful, false otherwise
   */
  export function Kv3Save(bigint kv, string error, number[] output, number flags): boolean;

  /**
   * Saves KeyValues3 data as JSON to a buffer
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param output - Vector to store the output JSON data
   * @returns true if successful, false otherwise
   */
  export function Kv3SaveAsJSON(bigint kv, string error, number[] output): boolean;

  /**
   * Saves KeyValues3 data as a JSON string
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param output - String to store the JSON output
   * @returns true if successful, false otherwise
   */
  export function Kv3SaveAsJSONString(bigint kv, string error, string output): boolean;

  /**
   * Saves KeyValues3 data as KeyValues1 text to a buffer
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param output - Vector to store the output KV1 text data
   * @param esc_behavior - Escape sequence behavior for KV1 text
   * @returns true if successful, false otherwise
   */
  export function Kv3SaveAsKV1Text(bigint kv, string error, number[] output, number esc_behavior): boolean;

  /**
   * Saves KeyValues3 data as KeyValues1 text with translation to a buffer
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param output - Vector to store the output KV1 text data
   * @param esc_behavior - Escape sequence behavior for KV1 text
   * @param translation - Pointer to translation table
   * @param unk - Unknown integer parameter
   * @returns true if successful, false otherwise
   */
  export function Kv3SaveAsKV1TextTranslated(bigint kv, string error, number[] output, number esc_behavior, bigint translation, number unk): boolean;

  /**
   * Saves KeyValues3 text without a header to a buffer
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param output - Vector to store the output text data
   * @param flags - Saving flags
   * @returns true if successful, false otherwise
   */
  export function Kv3SaveTextNoHeaderToBuffer(bigint kv, string error, number[] output, number flags): boolean;

  /**
   * Saves KeyValues3 text without a header to a string
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param output - String to store the text output
   * @param flags - Saving flags
   * @returns true if successful, false otherwise
   */
  export function Kv3SaveTextNoHeader(bigint kv, string error, string output, number flags): boolean;

  /**
   * Saves KeyValues3 text to a string
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param output - String to store the text output
   * @param flags - Saving flags
   * @returns true if successful, false otherwise
   */
  export function Kv3SaveTextToString(bigint kv, string error, string output, number flags): boolean;

  /**
   * Saves KeyValues3 data to a file
   * @param kv - Pointer to the KeyValues3 object
   * @param error - Output string for error messages
   * @param filename - Name of the file to save
   * @param path - Path to save the file
   * @param flags - Saving flags
   * @returns true if successful, false otherwise
   */
  export function Kv3SaveToFile(bigint kv, string error, string filename, string path, number flags): boolean;

  /**
   * Triggers a breakpoint in the debugger.
   */
  export function DebugBreak(): void;

  /**
   * Draws a debug overlay box.
   * @param center - Center of the box in world space.
   * @param mins - Minimum bounds relative to the center.
   * @param maxs - Maximum bounds relative to the center.
   * @param r - Red color value.
   * @param g - Green color value.
   * @param b - Blue color value.
   * @param a - Alpha (transparency) value.
   * @param duration - Duration (in seconds) to display the box.
   */
  export function DebugDrawBox(Vector3 center, Vector3 mins, Vector3 maxs, number r, number g, number b, number a, number duration): void;

  /**
   * Draws a debug box oriented in the direction of a forward vector.
   * @param center - Center of the box.
   * @param mins - Minimum bounds.
   * @param maxs - Maximum bounds.
   * @param forward - Forward direction vector.
   * @param color - RGB color vector.
   * @param alpha - Alpha transparency.
   * @param duration - Duration (in seconds) to display the box.
   */
  export function DebugDrawBoxDirection(Vector3 center, Vector3 mins, Vector3 maxs, Vector3 forward, Vector3 color, number alpha, number duration): void;

  /**
   * Draws a debug circle.
   * @param center - Center of the circle.
   * @param color - RGB color vector.
   * @param alpha - Alpha transparency.
   * @param radius - Circle radius.
   * @param zTest - Whether to perform depth testing.
   * @param duration - Duration (in seconds) to display the circle.
   */
  export function DebugDrawCircle(Vector3 center, Vector3 color, number alpha, number radius, boolean zTest, number duration): void;

  /**
   * Clears all debug overlays.
   */
  export function DebugDrawClear(): void;

  /**
   * Draws a debug overlay line.
   * @param origin - Start point of the line.
   * @param target - End point of the line.
   * @param r - Red color value.
   * @param g - Green color value.
   * @param b - Blue color value.
   * @param zTest - Whether to perform depth testing.
   * @param duration - Duration (in seconds) to display the line.
   */
  export function DebugDrawLine(Vector3 origin, Vector3 target, number r, number g, number b, boolean zTest, number duration): void;

  /**
   * Draws a debug line using a color vector.
   * @param start - Start point of the line.
   * @param end - End point of the line.
   * @param color - RGB color vector.
   * @param zTest - Whether to perform depth testing.
   * @param duration - Duration (in seconds) to display the line.
   */
  export function DebugDrawLine_vCol(Vector3 start, Vector3 end, Vector3 color, boolean zTest, number duration): void;

  /**
   * Draws text at a specified screen position with line offset.
   * @param x - X coordinate in screen space.
   * @param y - Y coordinate in screen space.
   * @param lineOffset - Line offset value.
   * @param text - The text string to display.
   * @param r - Red color value.
   * @param g - Green color value.
   * @param b - Blue color value.
   * @param a - Alpha transparency value.
   * @param duration - Duration (in seconds) to display the text.
   */
  export function DebugDrawScreenTextLine(number x, number y, number lineOffset, string text, number r, number g, number b, number a, number duration): void;

  /**
   * Draws a debug sphere.
   * @param center - Center of the sphere.
   * @param color - RGB color vector.
   * @param alpha - Alpha transparency.
   * @param radius - Radius of the sphere.
   * @param zTest - Whether to perform depth testing.
   * @param duration - Duration (in seconds) to display the sphere.
   */
  export function DebugDrawSphere(Vector3 center, Vector3 color, number alpha, number radius, boolean zTest, number duration): void;

  /**
   * Draws text in 3D space.
   * @param origin - World-space position to draw the text at.
   * @param text - The text string to display.
   * @param viewCheck - If true, only draws when visible to camera.
   * @param duration - Duration (in seconds) to display the text.
   */
  export function DebugDrawText(Vector3 origin, string text, boolean viewCheck, number duration): void;

  /**
   * Draws styled debug text on screen.
   * @param x - X coordinate.
   * @param y - Y coordinate.
   * @param lineOffset - Line offset value.
   * @param text - Text string.
   * @param r - Red color value.
   * @param g - Green color value.
   * @param b - Blue color value.
   * @param a - Alpha transparency.
   * @param duration - Duration (in seconds) to display the text.
   * @param font - Font name.
   * @param size - Font size.
   * @param bold - Whether text should be bold.
   */
  export function DebugScreenTextPretty(number x, number y, number lineOffset, string text, number r, number g, number b, number a, number duration, string font, number size, boolean bold): void;

  /**
   * Performs an assertion and logs a message if the assertion fails.
   * @param assertion - Boolean value to test.
   * @param message - Message to display if the assertion fails.
   */
  export function DebugScriptAssert(boolean assertion, string message): void;

  /**
   * Returns angular difference in degrees
   * @param angle1 - First angle in degrees
   * @param angle2 - Second angle in degrees
   * @returns Angular difference in degrees
   */
  export function AnglesDiff(number angle1, number angle2): number;

  /**
   * Converts QAngle to directional Vector
   * @param angles - The QAngle to convert
   * @returns Directional vector
   */
  export function AnglesToVector(Vector3 angles): Vector3;

  /**
   * Converts axis-angle representation to quaternion
   * @param axis - Rotation axis (should be normalized)
   * @param angle - Rotation angle in radians
   * @returns Resulting quaternion
   */
  export function AxisAngleToQuaternion(Vector3 axis, number angle): Vector4;

  /**
   * Computes closest point on an entity's oriented bounding box (OBB)
   * @param entityHandle - Handle of the entity
   * @param position - Position to find closest point from
   * @returns Closest point on the entity's OBB, or vec3_origin if entity is invalid
   */
  export function CalcClosestPointOnEntityOBB(number entityHandle, Vector3 position): Vector3;

  /**
   * Computes distance between two entities' oriented bounding boxes (OBBs)
   * @param entityHandle1 - Handle of the first entity
   * @param entityHandle2 - Handle of the second entity
   * @returns Distance between OBBs, or -1.0f if either entity is invalid
   */
  export function CalcDistanceBetweenEntityOBB(number entityHandle1, number entityHandle2): number;

  /**
   * Computes shortest 2D distance from a point to a line segment
   * @param p - The point
   * @param vLineA - First endpoint of the line segment
   * @param vLineB - Second endpoint of the line segment
   * @returns Shortest 2D distance
   */
  export function CalcDistanceToLineSegment2D(Vector3 p, Vector3 vLineA, Vector3 vLineB): number;

  /**
   * Computes cross product of two vectors
   * @param v1 - First vector
   * @param v2 - Second vector
   * @returns Cross product vector (v1 × v2)
   */
  export function CrossVectors(Vector3 v1, Vector3 v2): Vector3;

  /**
   * Smooth exponential decay function
   * @param decayTo - Target value to decay towards
   * @param decayTime - Time constant for decay
   * @param dt - Delta time
   * @returns Decay factor
   */
  export function ExponentDecay(number decayTo, number decayTime, number dt): number;

  /**
   * Linear interpolation between two vectors
   * @param start - Starting vector
   * @param end - Ending vector
   * @param factor - Interpolation factor (0.0 to 1.0)
   * @returns Interpolated vector
   */
  export function LerpVectors(Vector3 start, Vector3 end, number factor): Vector3;

  /**
   * Quaternion spherical linear interpolation for angles
   * @param fromAngle - Starting angle
   * @param toAngle - Ending angle
   * @param time - Interpolation time (0.0 to 1.0)
   * @returns Interpolated angle
   */
  export function QSlerp(Vector3 fromAngle, Vector3 toAngle, number time): Vector3;

  /**
   * Rotate one QAngle by another
   * @param a1 - Base orientation
   * @param a2 - Rotation to apply
   * @returns Rotated orientation
   */
  export function RotateOrientation(Vector3 a1, Vector3 a2): Vector3;

  /**
   * Rotate a vector around a point by specified angle
   * @param rotationOrigin - Origin point of rotation
   * @param rotationAngle - Angle to rotate by
   * @param vectorToRotate - Vector to be rotated
   * @returns Rotated vector
   */
  export function RotatePosition(Vector3 rotationOrigin, Vector3 rotationAngle, Vector3 vectorToRotate): Vector3;

  /**
   * Rotates quaternion by axis-angle representation
   * @param q - Quaternion to rotate
   * @param axis - Rotation axis
   * @param angle - Rotation angle in radians
   * @returns Rotated quaternion
   */
  export function RotateQuaternionByAxisAngle(Vector4 q, Vector3 axis, number angle): Vector4;

  /**
   * Finds angular delta between two QAngles
   * @param src - Source angle
   * @param dest - Destination angle
   * @returns Delta angle from src to dest
   */
  export function RotationDelta(Vector3 src, Vector3 dest): Vector3;

  /**
   * Converts delta QAngle to angular velocity vector
   * @param a1 - First angle
   * @param a2 - Second angle
   * @returns Angular velocity vector
   */
  export function RotationDeltaAsAngularVelocity(Vector3 a1, Vector3 a2): Vector3;

  /**
   * Interpolates between two quaternions using spline
   * @param q0 - Starting quaternion
   * @param q1 - Ending quaternion
   * @param t - Interpolation parameter (0.0 to 1.0)
   * @returns Interpolated quaternion
   */
  export function SplineQuaternions(Vector4 q0, Vector4 q1, number t): Vector4;

  /**
   * Interpolates between two vectors using spline
   * @param v0 - Starting vector
   * @param v1 - Ending vector
   * @param t - Interpolation parameter (0.0 to 1.0)
   * @returns Interpolated vector
   */
  export function SplineVectors(Vector3 v0, Vector3 v1, number t): Vector3;

  /**
   * Converts directional vector to QAngle (no roll)
   * @param input - Direction vector
   * @returns Angle representation with pitch and yaw (roll is 0)
   */
  export function VectorToAngles(Vector3 input): Vector3;

  /**
   * Returns random float between min and max
   * @param min - Minimum value (inclusive)
   * @param max - Maximum value (inclusive)
   * @returns Random float in range [min, max]
   */
  export function RandomFlt(number min, number max): number;

  /**
   * Returns random integer between min and max (inclusive)
   * @param min - Minimum value (inclusive)
   * @param max - Maximum value (inclusive)
   * @returns Random integer in range [min, max]
   */
  export function RandomInt(number min, number max): number;

  /**
   * Performs a collideable trace using the VScript-compatible table call, exposing it through C++ exports.
   * @param start - Trace start position (world space)
   * @param end - Trace end position (world space)
   * @param entityHandle - Entity handle of the collideable
   * @param outPos - Output: position of impact
   * @param outFraction - Output: fraction of trace completed
   * @param outHit - Output: whether a hit occurred
   * @param outStartSolid - Output: whether trace started inside solid
   * @param outNormal - Output: surface normal at impact
   * @returns True if trace hit something, false otherwise
   */
  export function TraceCollideable(Vector3 start, Vector3 end, number entityHandle, Vector3 outPos, number outFraction, boolean outHit, boolean outStartSolid, Vector3 outNormal): boolean;

  /**
   * Performs a collideable trace using the VScript-compatible table call, exposing it through C++ exports.
   * @param start - Trace start position (world space)
   * @param end - Trace end position (world space)
   * @param entityHandle - Entity handle of the collideable
   * @param mins - Bounding box minimums
   * @param maxs - Bounding box maximums
   * @param outPos - Output: position of impact
   * @param outFraction - Output: fraction of trace completed
   * @param outHit - Output: whether a hit occurred
   * @param outStartSolid - Output: whether trace started inside solid
   * @param outNormal - Output: surface normal at impact
   * @returns True if trace hit something, false otherwise
   */
  export function TraceCollideable2(Vector3 start, Vector3 end, number entityHandle, bigint mins, bigint maxs, Vector3 outPos, number outFraction, boolean outHit, boolean outStartSolid, Vector3 outNormal): boolean;

  /**
   * Performs a hull trace with specified dimensions and mask.
   * @param start - Trace start position
   * @param end - Trace end position
   * @param min - Local bounding box minimums
   * @param max - Local bounding box maximums
   * @param mask - Trace mask
   * @param ignoreHandle - Entity handle to ignore during trace
   * @param outPos - Output: position of impact
   * @param outFraction - Output: fraction of trace completed
   * @param outHit - Output: whether a hit occurred
   * @param outEntHit - Output: handle of entity hit
   * @param outStartSolid - Output: whether trace started inside solid
   * @returns True if trace hit something, false otherwise
   */
  export function TraceHull(Vector3 start, Vector3 end, Vector3 min, Vector3 max, number mask, number ignoreHandle, Vector3 outPos, number outFraction, boolean outHit, number outEntHit, boolean outStartSolid): boolean;

  /**
   * Performs a line trace between two points.
   * @param startPos - Trace start position
   * @param endPos - Trace end position
   * @param mask - Trace mask
   * @param ignoreHandle - Entity handle to ignore during trace
   * @param outPos - Output: position of impact
   * @param outFraction - Output: fraction of trace completed
   * @param outHit - Output: whether a hit occurred
   * @param outEntHit - Output: handle of entity hit
   * @param outStartSolid - Output: whether trace started inside solid
   * @returns True if trace hit something, false otherwise
   */
  export function TraceLine(Vector3 startPos, Vector3 endPos, number mask, number ignoreHandle, Vector3 outPos, number outFraction, boolean outHit, number outEntHit, boolean outStartSolid): boolean;

  /**
   * Applies an impulse to an entity at a specific world position.
   * @param entityHandle - The handle of the entity.
   * @param position - The world position where the impulse will be applied.
   * @param impulse - The impulse vector to apply.
   */
  export function AddBodyImpulseAtPosition(number entityHandle, Vector3 position, Vector3 impulse): void;

  /**
   * Adds linear and angular velocity to the entity's physics object.
   * @param entityHandle - The handle of the entity.
   * @param linearVelocity - The linear velocity vector to add.
   * @param angularVelocity - The angular velocity vector to add.
   */
  export function AddBodyVelocity(number entityHandle, Vector3 linearVelocity, Vector3 angularVelocity): void;

  /**
   * Detaches the entity from its parent.
   * @param entityHandle - The handle of the entity.
   */
  export function DetachBodyFromParent(number entityHandle): void;

  /**
   * Retrieves the currently active sequence of the entity.
   * @param entityHandle - The handle of the entity.
   * @returns The sequence ID of the active sequence, or -1 if invalid.
   */
  export function GetBodySequence(number entityHandle): number;

  /**
   * Checks whether the entity is attached to a parent.
   * @param entityHandle - The handle of the entity.
   * @returns True if attached to a parent, false otherwise.
   */
  export function IsBodyAttachedToParent(number entityHandle): boolean;

  /**
   * Looks up a sequence ID by its name.
   * @param entityHandle - The handle of the entity.
   * @param name - The name of the sequence.
   * @returns The sequence ID, or -1 if not found.
   */
  export function LookupBodySequence(number entityHandle, string name): number;

  /**
   * Retrieves the duration of a specified sequence.
   * @param entityHandle - The handle of the entity.
   * @param sequenceName - The name of the sequence.
   * @returns The duration of the sequence in seconds, or 0 if invalid.
   */
  export function SetBodySequenceDuration(number entityHandle, string sequenceName): number;

  /**
   * Sets the angular velocity of the entity.
   * @param entityHandle - The handle of the entity.
   * @param angVelocity - The new angular velocity vector.
   */
  export function SetBodyAngularVelocity(number entityHandle, Vector3 angVelocity): void;

  /**
   * Sets the material group of the entity.
   * @param entityHandle - The handle of the entity.
   * @param materialGroup - The material group token to assign.
   */
  export function SetBodyMaterialGroup(number entityHandle, string materialGroup): void;

  /**
   * Sets the linear velocity of the entity.
   * @param entityHandle - The handle of the entity.
   * @param velocity - The new velocity vector.
   */
  export function SetBodyVelocity(number entityHandle, Vector3 velocity): void;

  /**
   * Retrieves the player slot from a given entity pointer.
   * @param entity - A pointer to the entity (CBaseEntity*).
   * @returns The player slot if valid, otherwise -1.
   */
  export function EntPointerToPlayerSlot(bigint entity): number;

  /**
   * Returns a pointer to the entity instance by player slot index.
   * @param playerSlot - Index of the player slot.
   * @returns Pointer to the entity instance, or nullptr if the slot is invalid.
   */
  export function PlayerSlotToEntPointer(number playerSlot): bigint;

  /**
   * Returns the entity handle associated with a player slot index.
   * @param playerSlot - Index of the player slot.
   * @returns The index of the entity, or -1 if the handle is invalid.
   */
  export function PlayerSlotToEntHandle(number playerSlot): number;

  /**
   * Retrieves the client object from a given player slot.
   * @param playerSlot - The index of the player's slot (0-based).
   * @returns A pointer to the client object if found, otherwise nullptr.
   */
  export function PlayerSlotToClientPtr(number playerSlot): bigint;

  /**
   * Retrieves the index of a given client object.
   * @param client - A pointer to the client object (CServerSideClient*).
   * @returns The player slot if found, otherwise -1.
   */
  export function ClientPtrToPlayerSlot(bigint client): number;

  /**
   * Returns the entity index for a given player slot.
   * @param playerSlot - The index of the player's slot.
   * @returns The entity index if valid, otherwise 0.
   */
  export function PlayerSlotToClientIndex(number playerSlot): number;

  /**
   * Retrieves the player slot from a given client index.
   * @param clientIndex - The index of the client.
   * @returns The player slot if valid, otherwise -1.
   */
  export function ClientIndexToPlayerSlot(number clientIndex): number;

  /**
   * Retrieves the player slot from a given player service.
   * @param service - The service pointer. Like CCSPlayer_ItemServices, CCSPlayer_WeaponServices ect.
   * @returns The player slot if valid, otherwise -1.
   */
  export function PlayerServicesToPlayerSlot(bigint service): number;

  /**
   * Retrieves a client's authentication string (SteamID).
   * @param playerSlot - The index of the player's slot whose authentication string is being retrieved.
   * @returns The authentication string.
   */
  export function GetClientAuthId(number playerSlot): string;

  /**
   * Returns the client's Steam account ID, a unique number identifying a given Steam account.
   * @param playerSlot - The index of the player's slot.
   * @returns uint32_t The client's steam account ID.
   */
  export function GetClientAccountId(number playerSlot): number;

  /**
   * Returns the client's SteamID64 â€” a unique 64-bit identifier of a Steam account.
   * @param playerSlot - The index of the player's slot.
   * @returns uint64_t The client's SteamID64.
   */
  export function GetClientSteamID64(number playerSlot): bigint;

  /**
   * Retrieves a client's IP address.
   * @param playerSlot - The index of the player's slot.
   * @returns The client's IP address.
   */
  export function GetClientIp(number playerSlot): string;

  /**
   * Retrieves a client's language.
   * @param playerSlot - The index of the player's slot.
   * @returns The client's language.
   */
  export function GetClientLanguage(number playerSlot): string;

  /**
   * Retrieves a client's operating system.
   * @param playerSlot - The index of the player's slot.
   * @returns The client's operating system.
   */
  export function GetClientOS(number playerSlot): string;

  /**
   * Returns the client's name.
   * @param playerSlot - The index of the player's slot.
   * @returns The client's name.
   */
  export function GetClientName(number playerSlot): string;

  /**
   * Returns the client's connection time in seconds.
   * @param playerSlot - The index of the player's slot.
   * @returns float Connection time in seconds.
   */
  export function GetClientTime(number playerSlot): number;

  /**
   * Returns the client's current latency (RTT).
   * @param playerSlot - The index of the player's slot.
   * @returns float Latency value.
   */
  export function GetClientLatency(number playerSlot): number;

  /**
   * Returns the client's access flags.
   * @param playerSlot - The index of the player's slot.
   * @returns uint64 Access flags as a bitmask.
   */
  export function GetUserFlagBits(number playerSlot): bigint;

  /**
   * Sets the access flags on a client using a bitmask.
   * @param playerSlot - The index of the player's slot.
   * @param flags - Bitmask representing the flags to be set.
   */
  export function SetUserFlagBits(number playerSlot, bigint flags): void;

  /**
   * Adds access flags to a client.
   * @param playerSlot - The index of the player's slot.
   * @param flags - Bitmask representing the flags to be added.
   */
  export function AddUserFlags(number playerSlot, bigint flags): void;

  /**
   * Removes access flags from a client.
   * @param playerSlot - The index of the player's slot.
   * @param flags - Bitmask representing the flags to be removed.
   */
  export function RemoveUserFlags(number playerSlot, bigint flags): void;

  /**
   * Checks if a certain player has been authenticated.
   * @param playerSlot - The index of the player's slot.
   * @returns true if the player is authenticated, false otherwise.
   */
  export function IsClientAuthorized(number playerSlot): boolean;

  /**
   * Checks if a certain player is connected.
   * @param playerSlot - The index of the player's slot.
   * @returns true if the player is connected, false otherwise.
   */
  export function IsClientConnected(number playerSlot): boolean;

  /**
   * Checks if a certain player has entered the game.
   * @param playerSlot - The index of the player's slot.
   * @returns true if the player is in the game, false otherwise.
   */
  export function IsClientInGame(number playerSlot): boolean;

  /**
   * Checks if a certain player is the SourceTV bot.
   * @param playerSlot - The index of the player's slot.
   * @returns true if the client is the SourceTV bot, false otherwise.
   */
  export function IsClientSourceTV(number playerSlot): boolean;

  /**
   * Checks if the client is alive or dead.
   * @param playerSlot - The index of the player's slot.
   * @returns true if the client is alive, false if dead.
   */
  export function IsClientAlive(number playerSlot): boolean;

  /**
   * Checks if a certain player is a fake client.
   * @param playerSlot - The index of the player's slot.
   * @returns true if the client is a fake client, false otherwise.
   */
  export function IsFakeClient(number playerSlot): boolean;

  /**
   * Retrieves the movement type of an client.
   * @param playerSlot - The index of the player's slot whose movement type is to be retrieved.
   * @returns The movement type of the entity, or 0 if the entity is invalid.
   */
  export function GetClientMoveType(number playerSlot): number;

  /**
   * Sets the movement type of an client.
   * @param playerSlot - The index of the player's slot whose movement type is to be set.
   * @param moveType - The movement type of the entity, or 0 if the entity is invalid.
   */
  export function SetClientMoveType(number playerSlot, number moveType): void;

  /**
   * Retrieves the gravity scale of an client.
   * @param playerSlot - The index of the player's slot whose gravity scale is to be retrieved.
   * @returns The gravity scale of the client, or 0.0f if the client is invalid.
   */
  export function GetClientGravity(number playerSlot): number;

  /**
   * Sets the gravity scale of an client.
   * @param playerSlot - The index of the player's slot whose gravity scale is to be set.
   * @param gravity - The new gravity scale to set for the client.
   */
  export function SetClientGravity(number playerSlot, number gravity): void;

  /**
   * Retrieves the flags of an client.
   * @param playerSlot - The index of the player's slot whose flags are to be retrieved.
   * @returns The flags of the client, or 0 if the client is invalid.
   */
  export function GetClientFlags(number playerSlot): number;

  /**
   * Sets the flags of an client.
   * @param playerSlot - The index of the player's slot whose flags are to be set.
   * @param flags - The new flags to set for the client.
   */
  export function SetClientFlags(number playerSlot, number flags): void;

  /**
   * Retrieves the render color of an client.
   * @param playerSlot - The index of the player's slot whose render color is to be retrieved.
   * @returns The raw color value of the client's render color, or 0 if the client is invalid.
   */
  export function GetClientRenderColor(number playerSlot): number;

  /**
   * Sets the render color of an client.
   * @param playerSlot - The index of the player's slot whose render color is to be set.
   * @param color - The new raw color value to set for the client's render color.
   */
  export function SetClientRenderColor(number playerSlot, number color): void;

  /**
   * Retrieves the render mode of an client.
   * @param playerSlot - The index of the player's slot whose render mode is to be retrieved.
   * @returns The render mode of the client, or 0 if the client is invalid.
   */
  export function GetClientRenderMode(number playerSlot): number;

  /**
   * Sets the render mode of an client.
   * @param playerSlot - The index of the player's slot whose render mode is to be set.
   * @param renderMode - The new render mode to set for the client.
   */
  export function SetClientRenderMode(number playerSlot, number renderMode): void;

  /**
   * Retrieves the mass of an client.
   * @param playerSlot - The index of the player's slot whose mass is to be retrieved.
   * @returns The mass of the client, or 0 if the client is invalid.
   */
  export function GetClientMass(number playerSlot): number;

  /**
   * Sets the mass of an client.
   * @param playerSlot - The index of the player's slot whose mass is to be set.
   * @param mass - The new mass value to set for the client.
   */
  export function SetClientMass(number playerSlot, number mass): void;

  /**
   * Retrieves the friction of an client.
   * @param playerSlot - The index of the player's slot whose friction is to be retrieved.
   * @returns The friction of the client, or 0 if the client is invalid.
   */
  export function GetClientFriction(number playerSlot): number;

  /**
   * Sets the friction of an client.
   * @param playerSlot - The index of the player's slot whose friction is to be set.
   * @param friction - The new friction value to set for the client.
   */
  export function SetClientFriction(number playerSlot, number friction): void;

  /**
   * Retrieves the health of an client.
   * @param playerSlot - The index of the player's slot whose health is to be retrieved.
   * @returns The health of the client, or 0 if the client is invalid.
   */
  export function GetClientHealth(number playerSlot): number;

  /**
   * Sets the health of an client.
   * @param playerSlot - The index of the player's slot whose health is to be set.
   * @param health - The new health value to set for the client.
   */
  export function SetClientHealth(number playerSlot, number health): void;

  /**
   * Retrieves the max health of an client.
   * @param playerSlot - The index of the player's slot whose max health is to be retrieved.
   * @returns The max health of the client, or 0 if the client is invalid.
   */
  export function GetClientMaxHealth(number playerSlot): number;

  /**
   * Sets the max health of an client.
   * @param playerSlot - The index of the player's slot whose max health is to be set.
   * @param maxHealth - The new max health value to set for the client.
   */
  export function SetClientMaxHealth(number playerSlot, number maxHealth): void;

  /**
   * Retrieves the team number of an client.
   * @param playerSlot - The index of the player's slot whose team number is to be retrieved.
   * @returns The team number of the client, or 0 if the client is invalid.
   */
  export function GetClientTeam(number playerSlot): number;

  /**
   * Sets the team number of an client.
   * @param playerSlot - The index of the player's slot whose team number is to be set.
   * @param team - The new team number to set for the client.
   */
  export function SetClientTeam(number playerSlot, number team): void;

  /**
   * Retrieves the absolute origin of an client.
   * @param playerSlot - The index of the player's slot whose absolute origin is to be retrieved.
   * @returns A vector where the absolute origin will be stored.
   */
  export function GetClientAbsOrigin(number playerSlot): Vector3;

  /**
   * Sets the absolute origin of an client.
   * @param playerSlot - The index of the player's slot whose absolute origin is to be set.
   * @param origin - The new absolute origin to set for the client.
   */
  export function SetClientAbsOrigin(number playerSlot, Vector3 origin): void;

  /**
   * Retrieves the absolute scale of an client.
   * @param playerSlot - The index of the player's slot whose absolute scale is to be retrieved.
   * @returns A vector where the absolute scale will be stored.
   */
  export function GetClientAbsScale(number playerSlot): number;

  /**
   * Sets the absolute scale of an client.
   * @param playerSlot - The index of the player's slot whose absolute scale is to be set.
   * @param scale - The new absolute scale to set for the client.
   */
  export function SetClientAbsScale(number playerSlot, number scale): void;

  /**
   * Retrieves the angular rotation of an client.
   * @param playerSlot - The index of the player's slot whose angular rotation is to be retrieved.
   * @returns A QAngle where the angular rotation will be stored.
   */
  export function GetClientAbsAngles(number playerSlot): Vector3;

  /**
   * Sets the angular rotation of an client.
   * @param playerSlot - The index of the player's slot whose angular rotation is to be set.
   * @param angle - The new angular rotation to set for the client.
   */
  export function SetClientAbsAngles(number playerSlot, Vector3 angle): void;

  /**
   * Retrieves the local origin of an client.
   * @param playerSlot - The index of the player's slot whose local origin is to be retrieved.
   * @returns A vector where the local origin will be stored.
   */
  export function GetClientLocalOrigin(number playerSlot): Vector3;

  /**
   * Sets the local origin of an client.
   * @param playerSlot - The index of the player's slot whose local origin is to be set.
   * @param origin - The new local origin to set for the client.
   */
  export function SetClientLocalOrigin(number playerSlot, Vector3 origin): void;

  /**
   * Retrieves the local scale of an client.
   * @param playerSlot - The index of the player's slot whose local scale is to be retrieved.
   * @returns A vector where the local scale will be stored.
   */
  export function GetClientLocalScale(number playerSlot): number;

  /**
   * Sets the local scale of an client.
   * @param playerSlot - The index of the player's slot whose local scale is to be set.
   * @param scale - The new local scale to set for the client.
   */
  export function SetClientLocalScale(number playerSlot, number scale): void;

  /**
   * Retrieves the angular rotation of an client.
   * @param playerSlot - The index of the player's slot whose angular rotation is to be retrieved.
   * @returns A QAngle where the angular rotation will be stored.
   */
  export function GetClientLocalAngles(number playerSlot): Vector3;

  /**
   * Sets the angular rotation of an client.
   * @param playerSlot - The index of the player's slot whose angular rotation is to be set.
   * @param angle - The new angular rotation to set for the client.
   */
  export function SetClientLocalAngles(number playerSlot, Vector3 angle): void;

  /**
   * Retrieves the absolute velocity of an client.
   * @param playerSlot - The index of the player's slot whose absolute velocity is to be retrieved.
   * @returns A vector where the absolute velocity will be stored.
   */
  export function GetClientAbsVelocity(number playerSlot): Vector3;

  /**
   * Sets the absolute velocity of an client.
   * @param playerSlot - The index of the player's slot whose absolute velocity is to be set.
   * @param velocity - The new absolute velocity to set for the client.
   */
  export function SetClientAbsVelocity(number playerSlot, Vector3 velocity): void;

  /**
   * Retrieves the base velocity of an client.
   * @param playerSlot - The index of the player's slot whose base velocity is to be retrieved.
   * @returns A vector where the base velocity will be stored.
   */
  export function GetClientBaseVelocity(number playerSlot): Vector3;

  /**
   * Retrieves the local angular velocity of an client.
   * @param playerSlot - The index of the player's slot whose local angular velocity is to be retrieved.
   * @returns A vector where the local angular velocity will be stored.
   */
  export function GetClientLocalAngVelocity(number playerSlot): Vector3;

  /**
   * Retrieves the angular velocity of an client.
   * @param playerSlot - The index of the player's slot whose angular velocity is to be retrieved.
   * @returns A vector where the angular velocity will be stored.
   */
  export function GetClientAngVelocity(number playerSlot): Vector3;

  /**
   * Sets the angular velocity of an client.
   * @param playerSlot - The index of the player's slot whose angular velocity is to be set.
   * @param velocity - The new angular velocity to set for the client.
   */
  export function SetClientAngVelocity(number playerSlot, Vector3 velocity): void;

  /**
   * Retrieves the local velocity of an client.
   * @param playerSlot - The index of the player's slot whose local velocity is to be retrieved.
   * @returns A vector where the local velocity will be stored.
   */
  export function GetClientLocalVelocity(number playerSlot): Vector3;

  /**
   * Retrieves the angular rotation of an client.
   * @param playerSlot - The index of the player's slot whose angular rotation is to be retrieved.
   * @returns A vector where the angular rotation will be stored.
   */
  export function GetClientAngRotation(number playerSlot): Vector3;

  /**
   * Sets the angular rotation of an client.
   * @param playerSlot - The index of the player's slot whose angular rotation is to be set.
   * @param rotation - The new angular rotation to set for the client.
   */
  export function SetClientAngRotation(number playerSlot, Vector3 rotation): void;

  /**
   * Returns the input Vector transformed from client to world space.
   * @param playerSlot - The index of the player's slot
   * @param point - Point in client local space to transform
   * @returns The point transformed to world space coordinates
   */
  export function TransformPointClientToWorld(number playerSlot, Vector3 point): Vector3;

  /**
   * Returns the input Vector transformed from world to client space.
   * @param playerSlot - The index of the player's slot
   * @param point - Point in world space to transform
   * @returns The point transformed to client local space coordinates
   */
  export function TransformPointWorldToClient(number playerSlot, Vector3 point): Vector3;

  /**
   * Get vector to eye position - absolute coords.
   * @param playerSlot - The index of the player's slot
   * @returns Eye position in absolute/world coordinates
   */
  export function GetClientEyePosition(number playerSlot): Vector3;

  /**
   * Get the qangles that this client is looking at.
   * @param playerSlot - The index of the player's slot
   * @returns Eye angles as a vector (pitch, yaw, roll)
   */
  export function GetClientEyeAngles(number playerSlot): Vector3;

  /**
   * Sets the forward velocity of an client.
   * @param playerSlot - The index of the player's slot whose forward velocity is to be set.
   * @param forward
   */
  export function SetClientForwardVector(number playerSlot, Vector3 forward): void;

  /**
   * Get the forward vector of the client.
   * @param playerSlot - The index of the player's slot to query
   * @returns Forward-facing direction vector of the client
   */
  export function GetClientForwardVector(number playerSlot): Vector3;

  /**
   * Get the left vector of the client.
   * @param playerSlot - The index of the player's slot to query
   * @returns Left-facing direction vector of the client (aligned with the y axis)
   */
  export function GetClientLeftVector(number playerSlot): Vector3;

  /**
   * Get the right vector of the client.
   * @param playerSlot - The index of the player's slot to query
   * @returns Right-facing direction vector of the client
   */
  export function GetClientRightVector(number playerSlot): Vector3;

  /**
   * Get the up vector of the client.
   * @param playerSlot - The index of the player's slot to query
   * @returns Up-facing direction vector of the client
   */
  export function GetClientUpVector(number playerSlot): Vector3;

  /**
   * Get the client-to-world transformation matrix.
   * @param playerSlot - The index of the player's slot to query
   * @returns 4x4 transformation matrix representing client's position, rotation, and scale in world space
   */
  export function GetClientTransform(number playerSlot): Matrix4x4;

  /**
   * Retrieves the model name of an client.
   * @param playerSlot - The index of the player's slot whose model name is to be retrieved.
   * @returns A string where the model name will be stored.
   */
  export function GetClientModel(number playerSlot): string;

  /**
   * Sets the model name of an client.
   * @param playerSlot - The index of the player's slot whose model name is to be set.
   * @param model - The new model name to set for the client.
   */
  export function SetClientModel(number playerSlot, string model): void;

  /**
   * Retrieves the water level of an client.
   * @param playerSlot - The index of the player's slot whose water level is to be retrieved.
   * @returns The water level of the client, or 0.0f if the client is invalid.
   */
  export function GetClientWaterLevel(number playerSlot): number;

  /**
   * Retrieves the ground client of an client.
   * @param playerSlot - The index of the player's slot whose ground client is to be retrieved.
   * @returns The handle of the ground client, or INVALID_EHANDLE_INDEX if the client is invalid.
   */
  export function GetClientGroundEntity(number playerSlot): number;

  /**
   * Retrieves the effects of an client.
   * @param playerSlot - The index of the player's slot whose effects are to be retrieved.
   * @returns The effect flags of the client, or 0 if the client is invalid.
   */
  export function GetClientEffects(number playerSlot): number;

  /**
   * Adds the render effect flag to an client.
   * @param playerSlot - The index of the player's slot to modify
   * @param effects - Render effect flags to add
   */
  export function AddClientEffects(number playerSlot, number effects): void;

  /**
   * Removes the render effect flag from an client.
   * @param playerSlot - The index of the player's slot to modify
   * @param effects - Render effect flags to remove
   */
  export function RemoveClientEffects(number playerSlot, number effects): void;

  /**
   * Get a vector containing max bounds, centered on object.
   * @param playerSlot - The index of the player's slot to query
   * @returns Vector containing the maximum bounds of the client's bounding box
   */
  export function GetClientBoundingMaxs(number playerSlot): Vector3;

  /**
   * Get a vector containing min bounds, centered on object.
   * @param playerSlot - The index of the player's slot to query
   * @returns Vector containing the minimum bounds of the client's bounding box
   */
  export function GetClientBoundingMins(number playerSlot): Vector3;

  /**
   * Get vector to center of object - absolute coords.
   * @param playerSlot - The index of the player's slot to query
   * @returns Vector pointing to the center of the client in absolute/world coordinates
   */
  export function GetClientCenter(number playerSlot): Vector3;

  /**
   * Teleports an client to a specified location and orientation.
   * @param playerSlot - The index of the player's slot to teleport.
   * @param origin - A pointer to a Vector representing the new absolute position. Can be nullptr.
   * @param angles - A pointer to a QAngle representing the new orientation. Can be nullptr.
   * @param velocity - A pointer to a Vector representing the new velocity. Can be nullptr.
   */
  export function TeleportClient(number playerSlot, bigint origin, bigint angles, bigint velocity): void;

  /**
   * Apply an absolute velocity impulse to an client.
   * @param playerSlot - The index of the player's slot to apply impulse to
   * @param vecImpulse - Velocity impulse vector to apply
   */
  export function ApplyAbsVelocityImpulseToClient(number playerSlot, Vector3 vecImpulse): void;

  /**
   * Apply a local angular velocity impulse to an client.
   * @param playerSlot - The index of the player's slot to apply impulse to
   * @param angImpulse - Angular velocity impulse vector to apply
   */
  export function ApplyLocalAngularVelocityImpulseToClient(number playerSlot, Vector3 angImpulse): void;

  /**
   * Invokes a named input method on a specified client.
   * @param playerSlot - The handle of the target client that will receive the input.
   * @param inputName - The name of the input action to invoke.
   * @param activatorHandle - The index of the player's slot that initiated the sequence of actions.
   * @param callerHandle - The index of the player's slot sending this event. Use -1 to specify
   * @param value - The value associated with the input action.
   * @param type - The type or classification of the value.
   * @param outputId - An identifier for tracking the output of this operation.
   */
  export function AcceptClientInput(number playerSlot, string inputName, number activatorHandle, number callerHandle, any value, number type, number outputId): void;

  /**
   * Connects a script function to an player output.
   * @param playerSlot - The handle of the player.
   * @param output - The name of the output to connect to.
   * @param functionName - The name of the script function to call.
   */
  export function ConnectClientOutput(number playerSlot, string output, string functionName): void;

  /**
   * Disconnects a script function from an player output.
   * @param playerSlot - The handle of the player.
   * @param output - The name of the output.
   * @param functionName - The name of the script function to disconnect.
   */
  export function DisconnectClientOutput(number playerSlot, string output, string functionName): void;

  /**
   * Disconnects a script function from an I/O event on a different player.
   * @param playerSlot - The handle of the calling player.
   * @param output - The name of the output.
   * @param functionName - The function name to disconnect.
   * @param targetHandle - The handle of the entity whose output is being disconnected.
   */
  export function DisconnectClientRedirectedOutput(number playerSlot, string output, string functionName, number targetHandle): void;

  /**
   * Fires an player output.
   * @param playerSlot - The handle of the player firing the output.
   * @param outputName - The name of the output to fire.
   * @param activatorHandle - The entity activating the output.
   * @param callerHandle - The entity that called the output.
   * @param value - The value associated with the input action.
   * @param type - The type or classification of the value.
   * @param delay - Delay in seconds before firing the output.
   */
  export function FireClientOutput(number playerSlot, string outputName, number activatorHandle, number callerHandle, any value, number type, number delay): void;

  /**
   * Redirects an player output to call a function on another player.
   * @param playerSlot - The handle of the player whose output is being redirected.
   * @param output - The name of the output to redirect.
   * @param functionName - The function name to call on the target player.
   * @param targetHandle - The handle of the entity that will receive the output call.
   */
  export function RedirectClientOutput(number playerSlot, string output, string functionName, number targetHandle): void;

  /**
   * Makes an client follow another client with optional bone merging.
   * @param playerSlot - The index of the player's slot that will follow
   * @param attachmentHandle - The index of the player's slot to follow
   * @param boneMerge - If true, bones will be merged between entities
   */
  export function FollowClient(number playerSlot, number attachmentHandle, boolean boneMerge): void;

  /**
   * Makes an client follow another client and merge with a specific bone or attachment.
   * @param playerSlot - The index of the player's slot that will follow
   * @param attachmentHandle - The index of the player's slot to follow
   * @param boneOrAttachName - Name of the bone or attachment point to merge with
   */
  export function FollowClientMerge(number playerSlot, number attachmentHandle, string boneOrAttachName): void;

  /**
   * Apply damage to an client.
   * @param playerSlot - The index of the player's slot receiving damage
   * @param inflictorSlot - The index of the player's slot inflicting damage (e.g., projectile)
   * @param attackerSlot - The index of the attacking client
   * @param force - Direction and magnitude of force to apply
   * @param hitPos - Position where the damage hit occurred
   * @param damage - Amount of damage to apply
   * @param damageTypes - Bitfield of damage type flags
   * @returns Amount of damage actually applied to the client
   */
  export function TakeClientDamage(number playerSlot, number inflictorSlot, number attackerSlot, Vector3 force, Vector3 hitPos, number damage, number damageTypes): number;

  /**
   * Retrieves the pawn entity pointer associated with a client.
   * @param playerSlot - The index of the player's slot.
   * @returns A pointer to the client's pawn entity, or nullptr if the client or controller is invalid.
   */
  export function GetClientPawn(number playerSlot): bigint;

  /**
   * Processes the target string to determine if one user can target another.
   * @param caller - The index of the player's slot making the target request.
   * @param target - The target string specifying the player or players to be targeted.
   * @returns A vector where the result of the targeting operation will be stored.
   */
  export function ProcessTargetString(number caller, string target): number[];

  /**
   * Switches the player's team.
   * @param playerSlot - The index of the player's slot.
   * @param team - The team index to switch the client to.
   */
  export function SwitchClientTeam(number playerSlot, number team): void;

  /**
   * Respawns a player.
   * @param playerSlot - The index of the player's slot to respawn.
   */
  export function RespawnClient(number playerSlot): void;

  /**
   * Forces a player to commit suicide.
   * @param playerSlot - The index of the player's slot.
   * @param explode - If true, the client will explode upon death.
   * @param force - If true, the suicide will be forced.
   */
  export function ForcePlayerSuicide(number playerSlot, boolean explode, boolean force): void;

  /**
   * Disconnects a client from the server as soon as the next frame starts.
   * @param playerSlot - The index of the player's slot to be kicked.
   */
  export function KickClient(number playerSlot): void;

  /**
   * Bans a client for a specified duration.
   * @param playerSlot - The index of the player's slot to be banned.
   * @param duration - Duration of the ban in seconds.
   * @param kick - If true, the client will be kicked immediately after being banned.
   */
  export function BanClient(number playerSlot, number duration, boolean kick): void;

  /**
   * Bans an identity (either an IP address or a Steam authentication string).
   * @param steamId - The Steam ID to ban.
   * @param duration - Duration of the ban in seconds.
   * @param kick - If true, the client will be kicked immediately after being banned.
   */
  export function BanIdentity(bigint steamId, number duration, boolean kick): void;

  /**
   * Retrieves the handle of the client's currently active weapon.
   * @param playerSlot - The index of the player's slot.
   * @returns The entity handle of the active weapon, or INVALID_EHANDLE_INDEX if the client is invalid or has no active weapon.
   */
  export function GetClientActiveWeapon(number playerSlot): number;

  /**
   * Retrieves a list of weapon handles owned by the client.
   * @param playerSlot - The index of the player's slot.
   * @returns A vector of entity handles for the client's weapons, or an empty vector if the client is invalid or has no weapons.
   */
  export function GetClientWeapons(number playerSlot): number[];

  /**
   * Removes all weapons from a client, with an option to remove the suit as well.
   * @param playerSlot - The index of the player's slot.
   * @param removeSuit - A boolean indicating whether to also remove the client's suit.
   */
  export function RemoveWeapons(number playerSlot, boolean removeSuit): void;

  /**
   * Forces a player to drop their weapon.
   * @param playerSlot - The index of the player's slot.
   * @param weaponHandle - The handle of weapon to drop.
   * @param target - Target direction.
   * @param velocity - Velocity to toss weapon or zero to just drop weapon.
   */
  export function DropWeapon(number playerSlot, number weaponHandle, Vector3 target, Vector3 velocity): void;

  /**
   * Selects a player's weapon.
   * @param playerSlot - The index of the player's slot.
   * @param weaponHandle - The handle of weapon to bump.
   */
  export function SelectWeapon(number playerSlot, number weaponHandle): void;

  /**
   * Switches a player's weapon.
   * @param playerSlot - The index of the player's slot.
   * @param weaponHandle - The handle of weapon to switch.
   */
  export function SwitchWeapon(number playerSlot, number weaponHandle): void;

  /**
   * Removes a player's weapon.
   * @param playerSlot - The index of the player's slot.
   * @param weaponHandle - The handle of weapon to remove.
   */
  export function RemoveWeapon(number playerSlot, number weaponHandle): void;

  /**
   * Gives a named item (e.g., weapon) to a client.
   * @param playerSlot - The index of the player's slot.
   * @param itemName - The name of the item to give.
   * @returns The entity handle of the created item, or INVALID_EHANDLE_INDEX if the client or item is invalid.
   */
  export function GiveNamedItem(number playerSlot, string itemName): number;

  /**
   * Retrieves the state of a specific button for a client.
   * @param playerSlot - The index of the player's slot.
   * @param buttonIndex - The index of the button (0-2).
   * @returns uint64_t The state of the specified button, or 0 if the client or button index is invalid.
   */
  export function GetClientButtons(number playerSlot, number buttonIndex): bigint;

  /**
   * Returns the client's armor value.
   * @param playerSlot - The index of the player's slot.
   * @returns The armor value of the client.
   */
  export function GetClientArmor(number playerSlot): number;

  /**
   * Sets the client's armor value.
   * @param playerSlot - The index of the player's slot.
   * @param armor - The armor value to set.
   */
  export function SetClientArmor(number playerSlot, number armor): void;

  /**
   * Returns the client's speed value.
   * @param playerSlot - The index of the player's slot.
   * @returns The speed value of the client.
   */
  export function GetClientSpeed(number playerSlot): number;

  /**
   * Sets the client's speed value.
   * @param playerSlot - The index of the player's slot.
   * @param speed - The speed value to set.
   */
  export function SetClientSpeed(number playerSlot, number speed): void;

  /**
   * Retrieves the amount of money a client has.
   * @param playerSlot - The index of the player's slot.
   * @returns The amount of money the client has, or 0 if the player slot is invalid.
   */
  export function GetClientMoney(number playerSlot): number;

  /**
   * Sets the amount of money for a client.
   * @param playerSlot - The index of the player's slot.
   * @param money - The amount of money to set.
   */
  export function SetClientMoney(number playerSlot, number money): void;

  /**
   * Retrieves the number of kills for a client.
   * @param playerSlot - The index of the player's slot.
   * @returns The number of kills the client has, or 0 if the player slot is invalid.
   */
  export function GetClientKills(number playerSlot): number;

  /**
   * Sets the number of kills for a client.
   * @param playerSlot - The index of the player's slot.
   * @param kills - The number of kills to set.
   */
  export function SetClientKills(number playerSlot, number kills): void;

  /**
   * Retrieves the number of deaths for a client.
   * @param playerSlot - The index of the player's slot.
   * @returns The number of deaths the client has, or 0 if the player slot is invalid.
   */
  export function GetClientDeaths(number playerSlot): number;

  /**
   * Sets the number of deaths for a client.
   * @param playerSlot - The index of the player's slot.
   * @param deaths - The number of deaths to set.
   */
  export function SetClientDeaths(number playerSlot, number deaths): void;

  /**
   * Retrieves the number of assists for a client.
   * @param playerSlot - The index of the player's slot.
   * @returns The number of assists the client has, or 0 if the player slot is invalid.
   */
  export function GetClientAssists(number playerSlot): number;

  /**
   * Sets the number of assists for a client.
   * @param playerSlot - The index of the player's slot.
   * @param assists - The number of assists to set.
   */
  export function SetClientAssists(number playerSlot, number assists): void;

  /**
   * Retrieves the total damage dealt by a client.
   * @param playerSlot - The index of the player's slot.
   * @returns The total damage dealt by the client, or 0 if the player slot is invalid.
   */
  export function GetClientDamage(number playerSlot): number;

  /**
   * Sets the total damage dealt by a client.
   * @param playerSlot - The index of the player's slot.
   * @param damage - The amount of damage to set.
   */
  export function SetClientDamage(number playerSlot, number damage): void;

  /**
   * Creates a console command as an administrative command.
   * @param name - The name of the console command.
   * @param adminFlags - The admin flags that indicate which admin level can use this command.
   * @param description - A brief description of what the command does.
   * @param flags - Command flags that define the behavior of the command.
   * @param callback - A callback function that is invoked when the command is executed.
   * @param type - Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @returns true if the command was successfully created; otherwise, false.
   */
  export function AddAdminCommand(string name, number adminFlags, string description, number flags, function callback, number type): boolean;

  /**
   * Creates a console command or hooks an already existing one.
   * @param name - The name of the console command.
   * @param description - A brief description of what the command does.
   * @param flags - Command flags that define the behavior of the command.
   * @param callback - A callback function that is invoked when the command is executed.
   * @param type - Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @returns true if the command was successfully created; otherwise, false.
   */
  export function AddConsoleCommand(string name, string description, number flags, function callback, number type): boolean;

  /**
   * Removes a console command from the system.
   * @param name - The name of the command to be removed.
   * @param callback - The callback function associated with the command to be removed.
   * @returns true if the command was successfully removed; otherwise, false.
   */
  export function RemoveCommand(string name, function callback): boolean;

  /**
   * Adds a callback that will fire when a command is sent to the server.
   * @param name - The name of the command.
   * @param callback - The callback function that will be invoked when the command is executed.
   * @param type - Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @returns Returns true if the callback was successfully added, false otherwise.
   */
  export function AddCommandListener(string name, function callback, number type): boolean;

  /**
   * Removes a callback that fires when a command is sent to the server.
   * @param name - The name of the command.
   * @param callback - The callback function to be removed.
   * @param type - Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @returns Returns true if the callback was successfully removed, false otherwise.
   */
  export function RemoveCommandListener(string name, function callback, number type): boolean;

  /**
   * Executes a server command as if it were run on the server console or through RCON.
   * @param command - The command to execute on the server.
   */
  export function ServerCommand(string command): void;

  /**
   * Executes a server command as if it were on the server console (or RCON) and stores the printed text into buffer.
   * @param command - The command to execute on the server.
   * @returns String to store command result into.
   */
  export function ServerCommandEx(string command): string;

  /**
   * Executes a client command.
   * @param playerSlot - The index of the client executing the command.
   * @param command - The command to execute on the client.
   */
  export function ClientCommand(number playerSlot, string command): void;

  /**
   * Executes a client command on the server without network communication.
   * @param playerSlot - The index of the client.
   * @param command - The command to be executed by the client.
   */
  export function FakeClientCommand(number playerSlot, string command): void;

  /**
   * Sends a message to the server console.
   * @param msg - The message to be sent to the server console.
   */
  export function PrintToServer(string msg): void;

  /**
   * Sends a message to a client's console.
   * @param playerSlot - The index of the player's slot to whom the message will be sent.
   * @param message - The message to be sent to the client's console.
   */
  export function PrintToConsole(number playerSlot, string message): void;

  /**
   * Prints a message to a specific client in the chat area.
   * @param playerSlot - The index of the player's slot to whom the message will be sent.
   * @param message - The message to be printed in the chat area.
   */
  export function PrintToChat(number playerSlot, string message): void;

  /**
   * Prints a message to a specific client in the center of the screen.
   * @param playerSlot - The index of the player's slot to whom the message will be sent.
   * @param message - The message to be printed in the center of the screen.
   */
  export function PrintCenterText(number playerSlot, string message): void;

  /**
   * Prints a message to a specific client with an alert box.
   * @param playerSlot - The index of the player's slot to whom the message will be sent.
   * @param message - The message to be printed in the alert box.
   */
  export function PrintAlertText(number playerSlot, string message): void;

  /**
   * Prints a html message to a specific client in the center of the screen.
   * @param playerSlot - The index of the player's slot to whom the message will be sent.
   * @param message - The HTML-formatted message to be printed.
   * @param duration - The duration of the message in seconds.
   */
  export function PrintCentreHtml(number playerSlot, string message, number duration): void;

  /**
   * Sends a message to every client's console.
   * @param message - The message to be sent to all clients' consoles.
   */
  export function PrintToConsoleAll(string message): void;

  /**
   * Prints a message to all clients in the chat area.
   * @param message - The message to be printed in the chat area for all clients.
   */
  export function PrintToChatAll(string message): void;

  /**
   * Prints a message to all clients in the center of the screen.
   * @param message - The message to be printed in the center of the screen for all clients.
   */
  export function PrintCenterTextAll(string message): void;

  /**
   * Prints a message to all clients with an alert box.
   * @param message - The message to be printed in an alert box for all clients.
   */
  export function PrintAlertTextAll(string message): void;

  /**
   * Prints a html message to all clients in the center of the screen.
   * @param message - The HTML-formatted message to be printed in the center of the screen for all clients.
   * @param duration - The duration of the message in seconds.
   */
  export function PrintCentreHtmlAll(string message, number duration): void;

  /**
   * Prints a colored message to a specific client in the chat area.
   * @param playerSlot - The index of the player's slot to whom the message will be sent.
   * @param message - The message to be printed in the chat area with color.
   */
  export function PrintToChatColored(number playerSlot, string message): void;

  /**
   * Prints a colored message to all clients in the chat area.
   * @param message - The colored message to be printed in the chat area for all clients.
   */
  export function PrintToChatColoredAll(string message): void;

  /**
   * Sends a reply message to a player or to the server console depending on the command context.
   * @param context - The context from which the command was called (e.g., Console or Chat).
   * @param playerSlot - The slot/index of the player receiving the message.
   * @param message - The message string to be sent as a reply.
   */
  export function ReplyToCommand(number context, number playerSlot, string message): void;

  /**
   * Creates a new console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value of the console variable.
   * @param description - A description of the console variable's purpose.
   * @param flags - Additional flags for the console variable.
   * @returns A handle to the created console variable.
   */
  export function CreateConVar(string name, any defaultValue, string description, number flags): bigint;

  /**
   * Creates a new boolean console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarBool(string name, boolean defaultValue, string description, number flags, boolean hasMin, boolean min, boolean hasMax, boolean max): bigint;

  /**
   * Creates a new 16-bit signed integer console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarInt16(string name, number defaultValue, string description, number flags, boolean hasMin, number min, boolean hasMax, number max): bigint;

  /**
   * Creates a new 16-bit unsigned integer console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarUInt16(string name, number defaultValue, string description, number flags, boolean hasMin, number min, boolean hasMax, number max): bigint;

  /**
   * Creates a new 32-bit signed integer console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarInt32(string name, number defaultValue, string description, number flags, boolean hasMin, number min, boolean hasMax, number max): bigint;

  /**
   * Creates a new 32-bit unsigned integer console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarUInt32(string name, number defaultValue, string description, number flags, boolean hasMin, number min, boolean hasMax, number max): bigint;

  /**
   * Creates a new 64-bit signed integer console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarInt64(string name, number defaultValue, string description, number flags, boolean hasMin, number min, boolean hasMax, number max): bigint;

  /**
   * Creates a new 64-bit unsigned integer console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarUInt64(string name, bigint defaultValue, string description, number flags, boolean hasMin, bigint min, boolean hasMax, bigint max): bigint;

  /**
   * Creates a new floating-point console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarFloat(string name, number defaultValue, string description, number flags, boolean hasMin, number min, boolean hasMax, number max): bigint;

  /**
   * Creates a new double-precision console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarDouble(string name, number defaultValue, string description, number flags, boolean hasMin, number min, boolean hasMax, number max): bigint;

  /**
   * Creates a new color console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default color value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum color value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum color value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarColor(string name, number defaultValue, string description, number flags, boolean hasMin, number min, boolean hasMax, number max): bigint;

  /**
   * Creates a new 2D vector console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarVector2(string name, Vector2 defaultValue, string description, number flags, boolean hasMin, Vector2 min, boolean hasMax, Vector2 max): bigint;

  /**
   * Creates a new 3D vector console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarVector3(string name, Vector3 defaultValue, string description, number flags, boolean hasMin, Vector3 min, boolean hasMax, Vector3 max): bigint;

  /**
   * Creates a new 4D vector console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarVector4(string name, Vector4 defaultValue, string description, number flags, boolean hasMin, Vector4 min, boolean hasMax, Vector4 max): bigint;

  /**
   * Creates a new quaternion angle console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value for the console variable.
   * @param description - A brief description of the console variable.
   * @param flags - Flags that define the behavior of the console variable.
   * @param hasMin - Indicates if a minimum value is provided.
   * @param min - The minimum value if hasMin is true.
   * @param hasMax - Indicates if a maximum value is provided.
   * @param max - The maximum value if hasMax is true.
   * @returns A handle to the created console variable data.
   */
  export function CreateConVarQAngle(string name, Vector3 defaultValue, string description, number flags, boolean hasMin, Vector3 min, boolean hasMax, Vector3 max): bigint;

  /**
   * Creates a new string console variable.
   * @param name - The name of the console variable.
   * @param defaultValue - The default value of the console variable.
   * @param description - A description of the console variable's purpose.
   * @param flags - Additional flags for the console variable.
   * @returns A handle to the created console variable.
   */
  export function CreateConVarString(string name, string defaultValue, string description, number flags): bigint;

  /**
   * Searches for a console variable.
   * @param name - The name of the console variable to search for.
   * @returns A handle to the console variable data if found; otherwise, nullptr.
   */
  export function FindConVar(string name): bigint;

  /**
   * Searches for a console variable of a specific type.
   * @param name - The name of the console variable to search for.
   * @param type - The type of the console variable to search for.
   * @returns A handle to the console variable data if found; otherwise, nullptr.
   */
  export function FindConVar2(string name, number type): bigint;

  /**
   * Creates a hook for when a console variable's value is changed.
   * @param conVarHandle - TThe handle to the console variable data.
   * @param callback - The callback function to be executed when the variable's value changes.
   */
  export function HookConVarChange(bigint conVarHandle, function callback): void;

  /**
   * Removes a hook for when a console variable's value is changed.
   * @param uint64 - The handle to the console variable data.
   * @param callback - The callback function to be removed.
   */
  export function UnhookConVarChange(string uint64, function callback): void;

  /**
   * Checks if a specific flag is set for a console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param flag - The flag to check against the console variable.
   * @returns True if the flag is set; otherwise, false.
   */
  export function IsConVarFlagSet(bigint conVarHandle, number flag): boolean;

  /**
   * Adds flags to a console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param flags - The flags to be added.
   */
  export function AddConVarFlags(bigint conVarHandle, number flags): void;

  /**
   * Removes flags from a console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param flags - The flags to be removed.
   */
  export function RemoveConVarFlags(bigint conVarHandle, number flags): void;

  /**
   * Retrieves the current flags of a console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current flags set on the console variable.
   */
  export function GetConVarFlags(bigint conVarHandle): number;

  /**
   * Gets the specified bound (max or min) of a console variable and stores it in the output string.
   * @param conVarHandle - The handle to the console variable data.
   * @param max - Indicates whether to get the maximum (true) or minimum (false) bound.
   * @returns The bound value.
   */
  export function GetConVarBounds(bigint conVarHandle, boolean max): string;

  /**
   * Sets the specified bound (max or min) for a console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param max - Indicates whether to set the maximum (true) or minimum (false) bound.
   * @param value - The value to set as the bound.
   */
  export function SetConVarBounds(bigint conVarHandle, boolean max, string value): void;

  /**
   * Retrieves the default value of a console variable and stores it in the output string.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The output value in string format.
   */
  export function GetConVarDefault(bigint conVarHandle): string;

  /**
   * Retrieves the current value of a console variable and stores it in the output string.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The output value in string format.
   */
  export function GetConVarValue(bigint conVarHandle): string;

  /**
   * Retrieves the current value of a console variable and stores it in the output.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The output value.
   */
  export function GetConVar(bigint conVarHandle): any;

  /**
   * Retrieves the current value of a boolean console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current boolean value of the console variable.
   */
  export function GetConVarBool(bigint conVarHandle): boolean;

  /**
   * Retrieves the current value of a signed 16-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current int16_t value of the console variable.
   */
  export function GetConVarInt16(bigint conVarHandle): number;

  /**
   * Retrieves the current value of an unsigned 16-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current uint16_t value of the console variable.
   */
  export function GetConVarUInt16(bigint conVarHandle): number;

  /**
   * Retrieves the current value of a signed 32-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current int32_t value of the console variable.
   */
  export function GetConVarInt32(bigint conVarHandle): number;

  /**
   * Retrieves the current value of an unsigned 32-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current uint32_t value of the console variable.
   */
  export function GetConVarUInt32(bigint conVarHandle): number;

  /**
   * Retrieves the current value of a signed 64-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current int64_t value of the console variable.
   */
  export function GetConVarInt64(bigint conVarHandle): number;

  /**
   * Retrieves the current value of an unsigned 64-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current uint64_t value of the console variable.
   */
  export function GetConVarUInt64(bigint conVarHandle): bigint;

  /**
   * Retrieves the current value of a float console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current float value of the console variable.
   */
  export function GetConVarFloat(bigint conVarHandle): number;

  /**
   * Retrieves the current value of a double console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current double value of the console variable.
   */
  export function GetConVarDouble(bigint conVarHandle): number;

  /**
   * Retrieves the current value of a string console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current string value of the console variable.
   */
  export function GetConVarString(bigint conVarHandle): string;

  /**
   * Retrieves the current value of a Color console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current Color value of the console variable.
   */
  export function GetConVarColor(bigint conVarHandle): number;

  /**
   * Retrieves the current value of a Vector2D console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current Vector2D value of the console variable.
   */
  export function GetConVarVector2(bigint conVarHandle): Vector2;

  /**
   * Retrieves the current value of a Vector console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current Vector value of the console variable.
   */
  export function GetConVarVector(bigint conVarHandle): Vector3;

  /**
   * Retrieves the current value of a Vector4D console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current Vector4D value of the console variable.
   */
  export function GetConVarVector4(bigint conVarHandle): Vector4;

  /**
   * Retrieves the current value of a QAngle console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @returns The current QAngle value of the console variable.
   */
  export function GetConVarQAngle(bigint conVarHandle): Vector3;

  /**
   * Sets the value of a console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The string value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarValue(bigint conVarHandle, string value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVar(bigint conVarHandle, any value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a boolean console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarBool(bigint conVarHandle, boolean value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a signed 16-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarInt16(bigint conVarHandle, number value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of an unsigned 16-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarUInt16(bigint conVarHandle, number value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a signed 32-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarInt32(bigint conVarHandle, number value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of an unsigned 32-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarUInt32(bigint conVarHandle, number value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a signed 64-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarInt64(bigint conVarHandle, number value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of an unsigned 64-bit integer console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarUInt64(bigint conVarHandle, bigint value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a floating-point console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarFloat(bigint conVarHandle, number value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a double-precision floating-point console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarDouble(bigint conVarHandle, number value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a string console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarString(bigint conVarHandle, string value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a color console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarColor(bigint conVarHandle, number value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a 2D vector console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarVector2(bigint conVarHandle, Vector2 value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a 3D vector console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarVector3(bigint conVarHandle, Vector3 value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a 4D vector console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarVector4(bigint conVarHandle, Vector4 value, boolean replicate, boolean notify): void;

  /**
   * Sets the value of a quaternion angle console variable.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to set for the console variable.
   * @param replicate - If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify - If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  export function SetConVarQAngle(bigint conVarHandle, Vector3 value, boolean replicate, boolean notify): void;

  /**
   * Replicates a console variable value to a specific client. This does not change the actual console variable value.
   * @param playerSlot - The index of the client to replicate the value to.
   * @param conVarHandle - The handle to the console variable data.
   * @param value - The value to send to the client.
   */
  export function SendConVarValue(number playerSlot, bigint conVarHandle, string value): void;

  /**
   * Retrieves the value of a client's console variable and stores it in the output string.
   * @param playerSlot - The index of the client whose console variable value is being retrieved.
   * @param convarName - The name of the console variable to retrieve.
   * @returns The output string to store the client's console variable value.
   */
  export function GetClientConVarValue(number playerSlot, string convarName): string;

  /**
   * Replicates a console variable value to a specific fake client. This does not change the actual console variable value.
   * @param playerSlot - The index of the fake client to replicate the value to.
   * @param convarName - The name of the console variable.
   * @param convarValue - The value to set for the console variable.
   */
  export function SetFakeClientConVarValue(number playerSlot, string convarName, string convarValue): void;

  /**
   * Starts a query to retrieve the value of a client's console variable.
   * @param playerSlot - The index of the player's slot to query the value from.
   * @param convarName - The name of client convar to query.
   * @param callback - A function to use as a callback when the query has finished.
   * @param data - Optional values to pass to the callback function.
   * @returns A cookie that uniquely identifies the query. Returns -1 on failure, such as when used on a bot.
   */
  export function QueryClientConVar(number playerSlot, string convarName, function callback, any[] data): number;

  /**
   *  Specifies that the given config file should be executed.
   * @param conVarHandles - List of handles to the console variable data.
   * @param autoCreate - If true, and the config file does not exist, such a config file will be automatically created and populated with information from the plugin's registered cvars.
   * @param name - Name of the config file, excluding the .cfg extension. Cannot be empty.
   * @param folder - Folder under cfg/ to use. By default this is "plugify." Can be empty.
   * @returns True on success, false otherwise.
   */
  export function AutoExecConfig(bigint[] conVarHandles, boolean autoCreate, string name, string folder): boolean;

  /**
   * Returns the current server language.
   * @returns The server language as a string.
   */
  export function GetServerLanguage(): string;

  /**
   * Finds a module by name.
   * @param name - The name of the module to find.
   * @returns A pointer to the specified module.
   */
  export function FindModule(string name): bigint;

  /**
   * Finds an interface by name.
   * @param name - The name of the interface to find.
   * @returns A pointer to the interface.
   */
  export function FindInterface(string name): bigint;

  /**
   * Queries an interface from a specified module.
   * @param module - The name of the module to query the interface from.
   * @param name - The name of the interface to find.
   * @returns A pointer to the queried interface.
   */
  export function QueryInterface(string module, string name): bigint;

  /**
   * Returns the path of the game's directory.
   * @returns A reference to a string where the game directory path will be stored.
   */
  export function GetGameDirectory(): string;

  /**
   * Returns the current map name.
   * @returns A reference to a string where the current map name will be stored.
   */
  export function GetCurrentMap(): string;

  /**
   * Returns whether a specified map is valid or not.
   * @param mapname - The name of the map to check for validity.
   * @returns True if the map is valid, false otherwise.
   */
  export function IsMapValid(string mapname): boolean;

  /**
   * Returns the game time based on the game tick.
   * @returns The current game time.
   */
  export function GetGameTime(): number;

  /**
   * Returns the game's internal tick count.
   * @returns The current tick count of the game.
   */
  export function GetGameTickCount(): number;

  /**
   * Returns the time the game took processing the last frame.
   * @returns The frame time of the last processed frame.
   */
  export function GetGameFrameTime(): number;

  /**
   * Returns a high-precision time value for profiling the engine.
   * @returns A high-precision time value.
   */
  export function GetEngineTime(): number;

  /**
   * Returns the maximum number of clients that can connect to the server.
   * @returns The maximum client count, or -1 if global variables are not initialized.
   */
  export function GetMaxClients(): number;

  /**
   * Precaches a given file.
   * @param resource - The name of the resource to be precached.
   */
  export function Precache(string resource): void;

  /**
   * Checks if a specified file is precached.
   * @param resource - The name of the file to check.
   */
  export function IsPrecached(string resource): boolean;

  /**
   * Returns a pointer to the Economy Item System.
   * @returns A pointer to the Econ Item System.
   */
  export function GetEconItemSystem(): bigint;

  /**
   * Checks if the server is currently paused.
   * @returns True if the server is paused, false otherwise.
   */
  export function IsServerPaused(): boolean;

  /**
   * Queues a task to be executed on the next frame.
   * @param callback - A callback function to be executed on the next frame.
   * @param userData - An array intended to hold user-related data, allowing for elements of any type.
   */
  export function QueueTaskForNextFrame(function callback, any[] userData): void;

  /**
   * Queues a task to be executed on the next world update.
   * @param callback - A callback function to be executed on the next world update.
   * @param userData - An array intended to hold user-related data, allowing for elements of any type.
   */
  export function QueueTaskForNextWorldUpdate(function callback, any[] userData): void;

  /**
   * Returns the duration of a specified sound.
   * @param name - The name of the sound to check.
   * @returns The duration of the sound in seconds.
   */
  export function GetSoundDuration(string name): number;

  /**
   * Emits a sound from a specified entity.
   * @param entityHandle - The handle of the entity that will emit the sound.
   * @param sound - The name of the sound to emit.
   * @param pitch - The pitch of the sound.
   * @param volume - The volume of the sound.
   * @param delay - The delay before the sound is played.
   */
  export function EmitSound(number entityHandle, string sound, number pitch, number volume, number delay): void;

  /**
   * Stops a sound from a specified entity.
   * @param entityHandle - The handle of the entity that will stop the sound.
   * @param sound - The name of the sound to stop.
   */
  export function StopSound(number entityHandle, string sound): void;

  /**
   * Emits a sound to a specific client.
   * @param playerSlot - The index of the client to whom the sound will be emitted.
   * @param channel - The channel through which the sound will be played.
   * @param sound - The name of the sound to emit.
   * @param volume - The volume of the sound.
   * @param soundLevel - The level of the sound.
   * @param flags - Additional flags for sound playback.
   * @param pitch - The pitch of the sound.
   * @param origin - The origin of the sound in 3D space.
   * @param soundTime - The time at which the sound should be played.
   */
  export function EmitSoundToClient(number playerSlot, number channel, string sound, number volume, number soundLevel, number flags, number pitch, Vector3 origin, number soundTime): void;

  /**
   * Converts an entity index into an entity pointer.
   * @param entityIndex - The index of the entity to convert.
   * @returns A pointer to the entity instance, or nullptr if the entity does not exist.
   */
  export function EntIndexToEntPointer(number entityIndex): bigint;

  /**
   * Retrieves the entity index from an entity pointer.
   * @param entity - A pointer to the entity whose index is to be retrieved.
   * @returns The index of the entity, or -1 if the entity is nullptr.
   */
  export function EntPointerToEntIndex(bigint entity): number;

  /**
   * Converts an entity pointer into an entity handle.
   * @param entity - A pointer to the entity to convert.
   * @returns The entity handle as an integer, or INVALID_EHANDLE_INDEX if the entity is nullptr.
   */
  export function EntPointerToEntHandle(bigint entity): number;

  /**
   * Retrieves the entity pointer from an entity handle.
   * @param entityHandle - The entity handle to convert.
   * @returns A pointer to the entity instance, or nullptr if the handle is invalid.
   */
  export function EntHandleToEntPointer(number entityHandle): bigint;

  /**
   * Converts an entity index into an entity handle.
   * @param entityIndex - The index of the entity to convert.
   * @returns The entity handle as an integer, or -1 if the entity index is invalid.
   */
  export function EntIndexToEntHandle(number entityIndex): number;

  /**
   * Retrieves the entity index from an entity handle.
   * @param entityHandle - The entity handle from which to retrieve the index.
   * @returns The index of the entity, or -1 if the handle is invalid.
   */
  export function EntHandleToEntIndex(number entityHandle): number;

  /**
   * Checks if the provided entity handle is valid.
   * @param entityHandle - The entity handle to check.
   * @returns True if the entity handle is valid, false otherwise.
   */
  export function IsValidEntHandle(number entityHandle): boolean;

  /**
   * Checks if the provided entity pointer is valid.
   * @param entity - The entity pointer to check.
   * @returns True if the entity pointer is valid, false otherwise.
   */
  export function IsValidEntPointer(bigint entity): boolean;

  /**
   * Retrieves the pointer to the first active entity.
   * @returns A pointer to the first active entity.
   */
  export function GetFirstActiveEntity(): bigint;

  /**
   * Retrieves a pointer to the concrete entity list.
   * @returns A pointer to the entity list structure.
   */
  export function GetConcreteEntityListPointer(): bigint;

  /**
   * Adds an entity output hook on a specified entity class name.
   * @param classname - The class name of the entity to hook the output for.
   * @param output - The output event name to hook.
   * @param callback - The callback function to invoke when the output is fired.
   * @param type - Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @returns True if the hook was successfully added, false otherwise.
   */
  export function HookEntityOutput(string classname, string output, function callback, number type): boolean;

  /**
   * Removes an entity output hook.
   * @param classname - The class name of the entity from which to unhook the output.
   * @param output - The output event name to unhook.
   * @param callback - The callback function that was previously hooked.
   * @param type - Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @returns True if the hook was successfully removed, false otherwise.
   */
  export function UnhookEntityOutput(string classname, string output, function callback, number type): boolean;

  /**
   * Finds an entity by classname within a radius with iteration.
   * @param startFrom - The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param classname - The class name to search for.
   * @param origin - The center of the search sphere.
   * @param radius - The search radius.
   * @returns The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  export function FindEntityByClassnameWithin(number startFrom, string classname, Vector3 origin, number radius): number;

  /**
   * Finds an entity by name with iteration.
   * @param startFrom - The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param name - The targetname to search for.
   * @returns The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  export function FindEntityByName(number startFrom, string name): number;

  /**
   * Finds the nearest entity by name to a point.
   * @param name - The targetname to search for.
   * @param origin - The point to search around.
   * @param maxRadius - Maximum search radius.
   * @returns The handle of the nearest entity, or INVALID_EHANDLE_INDEX if none found.
   */
  export function FindEntityByNameNearest(string name, Vector3 origin, number maxRadius): number;

  /**
   * Finds an entity by name within a radius with iteration.
   * @param startFrom - The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param name - The targetname to search for.
   * @param origin - The center of the search sphere.
   * @param radius - The search radius.
   * @returns The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  export function FindEntityByNameWithin(number startFrom, string name, Vector3 origin, number radius): number;

  /**
   * Finds an entity by targetname with iteration.
   * @param startFrom - The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param name - The targetname to search for.
   * @returns The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  export function FindEntityByTarget(number startFrom, string name): number;

  /**
   * Finds an entity within a sphere with iteration.
   * @param startFrom - The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param origin - The center of the search sphere.
   * @param radius - The search radius.
   * @returns The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  export function FindEntityInSphere(number startFrom, Vector3 origin, number radius): number;

  /**
   * Creates an entity by classname.
   * @param className - The class name of the entity to create.
   * @returns The handle of the created entity, or INVALID_EHANDLE_INDEX if creation failed.
   */
  export function SpawnEntityByName(string className): number;

  /**
   * Creates an entity by string name but does not spawn it.
   * @param className - The class name of the entity to create.
   * @returns The entity handle of the created entity, or INVALID_EHANDLE_INDEX if the entity could not be created.
   */
  export function CreateEntityByName(string className): number;

  /**
   * Spawns an entity into the game.
   * @param entityHandle - The handle of the entity to spawn.
   */
  export function DispatchSpawn(number entityHandle): void;

  /**
   * Spawns an entity into the game with key-value properties.
   * @param entityHandle - The handle of the entity to spawn.
   * @param keys - A vector of keys representing the property names to set on the entity.
   * @param values - A vector of values corresponding to the keys, representing the property values to set on the entity.
   */
  export function DispatchSpawn2(number entityHandle, string[] keys, any[] values): void;

  /**
   * Marks an entity for deletion.
   * @param entityHandle - The handle of the entity to be deleted.
   */
  export function RemoveEntity(number entityHandle): void;

  /**
   * Checks if an entity is a player controller.
   * @param entityHandle - The handle of the entity.
   * @returns True if the entity is a player controller, false otherwise.
   */
  export function IsEntityPlayerController(number entityHandle): boolean;

  /**
   * Checks if an entity is a player pawn.
   * @param entityHandle - The handle of the entity.
   * @returns True if the entity is a player pawn, false otherwise.
   */
  export function IsEntityPlayerPawn(number entityHandle): boolean;

  /**
   * Retrieves the class name of an entity.
   * @param entityHandle - The handle of the entity whose class name is to be retrieved.
   * @returns A string where the class name will be stored.
   */
  export function GetEntityClassname(number entityHandle): string;

  /**
   * Retrieves the name of an entity.
   * @param entityHandle - The handle of the entity whose name is to be retrieved.
   */
  export function GetEntityName(number entityHandle): string;

  /**
   * Sets the name of an entity.
   * @param entityHandle - The handle of the entity whose name is to be set.
   * @param name - The new name to set for the entity.
   */
  export function SetEntityName(number entityHandle, string name): void;

  /**
   * Retrieves the movement type of an entity.
   * @param entityHandle - The handle of the entity whose movement type is to be retrieved.
   * @returns The movement type of the entity, or 0 if the entity is invalid.
   */
  export function GetEntityMoveType(number entityHandle): number;

  /**
   * Sets the movement type of an entity.
   * @param entityHandle - The handle of the entity whose movement type is to be set.
   * @param moveType - The new movement type to set for the entity.
   */
  export function SetEntityMoveType(number entityHandle, number moveType): void;

  /**
   * Retrieves the gravity scale of an entity.
   * @param entityHandle - The handle of the entity whose gravity scale is to be retrieved.
   * @returns The gravity scale of the entity, or 0.0f if the entity is invalid.
   */
  export function GetEntityGravity(number entityHandle): number;

  /**
   * Sets the gravity scale of an entity.
   * @param entityHandle - The handle of the entity whose gravity scale is to be set.
   * @param gravity - The new gravity scale to set for the entity.
   */
  export function SetEntityGravity(number entityHandle, number gravity): void;

  /**
   * Retrieves the flags of an entity.
   * @param entityHandle - The handle of the entity whose flags are to be retrieved.
   * @returns The flags of the entity, or 0 if the entity is invalid.
   */
  export function GetEntityFlags(number entityHandle): number;

  /**
   * Sets the flags of an entity.
   * @param entityHandle - The handle of the entity whose flags are to be set.
   * @param flags - The new flags to set for the entity.
   */
  export function SetEntityFlags(number entityHandle, number flags): void;

  /**
   * Retrieves the render color of an entity.
   * @param entityHandle - The handle of the entity whose render color is to be retrieved.
   * @returns The raw color value of the entity's render color, or 0 if the entity is invalid.
   */
  export function GetEntityRenderColor(number entityHandle): number;

  /**
   * Sets the render color of an entity.
   * @param entityHandle - The handle of the entity whose render color is to be set.
   * @param color - The new raw color value to set for the entity's render color.
   */
  export function SetEntityRenderColor(number entityHandle, number color): void;

  /**
   * Retrieves the render mode of an entity.
   * @param entityHandle - The handle of the entity whose render mode is to be retrieved.
   * @returns The render mode of the entity, or 0 if the entity is invalid.
   */
  export function GetEntityRenderMode(number entityHandle): number;

  /**
   * Sets the render mode of an entity.
   * @param entityHandle - The handle of the entity whose render mode is to be set.
   * @param renderMode - The new render mode to set for the entity.
   */
  export function SetEntityRenderMode(number entityHandle, number renderMode): void;

  /**
   * Retrieves the mass of an entity.
   * @param entityHandle - The handle of the entity whose mass is to be retrieved.
   * @returns The mass of the entity, or 0 if the entity is invalid.
   */
  export function GetEntityMass(number entityHandle): number;

  /**
   * Sets the mass of an entity.
   * @param entityHandle - The handle of the entity whose mass is to be set.
   * @param mass - The new mass value to set for the entity.
   */
  export function SetEntityMass(number entityHandle, number mass): void;

  /**
   * Retrieves the friction of an entity.
   * @param entityHandle - The handle of the entity whose friction is to be retrieved.
   * @returns The friction of the entity, or 0 if the entity is invalid.
   */
  export function GetEntityFriction(number entityHandle): number;

  /**
   * Sets the friction of an entity.
   * @param entityHandle - The handle of the entity whose friction is to be set.
   * @param friction - The new friction value to set for the entity.
   */
  export function SetEntityFriction(number entityHandle, number friction): void;

  /**
   * Retrieves the health of an entity.
   * @param entityHandle - The handle of the entity whose health is to be retrieved.
   * @returns The health of the entity, or 0 if the entity is invalid.
   */
  export function GetEntityHealth(number entityHandle): number;

  /**
   * Sets the health of an entity.
   * @param entityHandle - The handle of the entity whose health is to be set.
   * @param health - The new health value to set for the entity.
   */
  export function SetEntityHealth(number entityHandle, number health): void;

  /**
   * Retrieves the max health of an entity.
   * @param entityHandle - The handle of the entity whose max health is to be retrieved.
   * @returns The max health of the entity, or 0 if the entity is invalid.
   */
  export function GetEntityMaxHealth(number entityHandle): number;

  /**
   * Sets the max health of an entity.
   * @param entityHandle - The handle of the entity whose max health is to be set.
   * @param maxHealth - The new max health value to set for the entity.
   */
  export function SetEntityMaxHealth(number entityHandle, number maxHealth): void;

  /**
   * Retrieves the team number of an entity.
   * @param entityHandle - The handle of the entity whose team number is to be retrieved.
   * @returns The team number of the entity, or 0 if the entity is invalid.
   */
  export function GetEntityTeam(number entityHandle): number;

  /**
   * Sets the team number of an entity.
   * @param entityHandle - The handle of the entity whose team number is to be set.
   * @param team - The new team number to set for the entity.
   */
  export function SetEntityTeam(number entityHandle, number team): void;

  /**
   * Retrieves the owner of an entity.
   * @param entityHandle - The handle of the entity whose owner is to be retrieved.
   * @returns The handle of the owner entity, or INVALID_EHANDLE_INDEX if the entity is invalid.
   */
  export function GetEntityOwner(number entityHandle): number;

  /**
   * Sets the owner of an entity.
   * @param entityHandle - The handle of the entity whose owner is to be set.
   * @param ownerHandle - The handle of the new owner entity.
   */
  export function SetEntityOwner(number entityHandle, number ownerHandle): void;

  /**
   * Retrieves the parent of an entity.
   * @param entityHandle - The handle of the entity whose parent is to be retrieved.
   * @returns The handle of the parent entity, or INVALID_EHANDLE_INDEX if the entity is invalid.
   */
  export function GetEntityParent(number entityHandle): number;

  /**
   * Sets the parent of an entity.
   * @param entityHandle - The handle of the entity whose parent is to be set.
   * @param parentHandle - The handle of the new parent entity.
   * @param attachmentName - The name of the entity's attachment.
   */
  export function SetEntityParent(number entityHandle, number parentHandle, string attachmentName): void;

  /**
   * Retrieves the absolute origin of an entity.
   * @param entityHandle - The handle of the entity whose absolute origin is to be retrieved.
   * @returns A vector where the absolute origin will be stored.
   */
  export function GetEntityAbsOrigin(number entityHandle): Vector3;

  /**
   * Sets the absolute origin of an entity.
   * @param entityHandle - The handle of the entity whose absolute origin is to be set.
   * @param origin - The new absolute origin to set for the entity.
   */
  export function SetEntityAbsOrigin(number entityHandle, Vector3 origin): void;

  /**
   * Retrieves the absolute scale of an entity.
   * @param entityHandle - The handle of the entity whose absolute scale is to be retrieved.
   * @returns A vector where the absolute scale will be stored.
   */
  export function GetEntityAbsScale(number entityHandle): number;

  /**
   * Sets the absolute scale of an entity.
   * @param entityHandle - The handle of the entity whose absolute scale is to be set.
   * @param scale - The new absolute scale to set for the entity.
   */
  export function SetEntityAbsScale(number entityHandle, number scale): void;

  /**
   * Retrieves the angular rotation of an entity.
   * @param entityHandle - The handle of the entity whose angular rotation is to be retrieved.
   * @returns A QAngle where the angular rotation will be stored.
   */
  export function GetEntityAbsAngles(number entityHandle): Vector3;

  /**
   * Sets the angular rotation of an entity.
   * @param entityHandle - The handle of the entity whose angular rotation is to be set.
   * @param angle - The new angular rotation to set for the entity.
   */
  export function SetEntityAbsAngles(number entityHandle, Vector3 angle): void;

  /**
   * Retrieves the local origin of an entity.
   * @param entityHandle - The handle of the entity whose local origin is to be retrieved.
   * @returns A vector where the local origin will be stored.
   */
  export function GetEntityLocalOrigin(number entityHandle): Vector3;

  /**
   * Sets the local origin of an entity.
   * @param entityHandle - The handle of the entity whose local origin is to be set.
   * @param origin - The new local origin to set for the entity.
   */
  export function SetEntityLocalOrigin(number entityHandle, Vector3 origin): void;

  /**
   * Retrieves the local scale of an entity.
   * @param entityHandle - The handle of the entity whose local scale is to be retrieved.
   * @returns A vector where the local scale will be stored.
   */
  export function GetEntityLocalScale(number entityHandle): number;

  /**
   * Sets the local scale of an entity.
   * @param entityHandle - The handle of the entity whose local scale is to be set.
   * @param scale - The new local scale to set for the entity.
   */
  export function SetEntityLocalScale(number entityHandle, number scale): void;

  /**
   * Retrieves the angular rotation of an entity.
   * @param entityHandle - The handle of the entity whose angular rotation is to be retrieved.
   * @returns A QAngle where the angular rotation will be stored.
   */
  export function GetEntityLocalAngles(number entityHandle): Vector3;

  /**
   * Sets the angular rotation of an entity.
   * @param entityHandle - The handle of the entity whose angular rotation is to be set.
   * @param angle - The new angular rotation to set for the entity.
   */
  export function SetEntityLocalAngles(number entityHandle, Vector3 angle): void;

  /**
   * Retrieves the absolute velocity of an entity.
   * @param entityHandle - The handle of the entity whose absolute velocity is to be retrieved.
   * @returns A vector where the absolute velocity will be stored.
   */
  export function GetEntityAbsVelocity(number entityHandle): Vector3;

  /**
   * Sets the absolute velocity of an entity.
   * @param entityHandle - The handle of the entity whose absolute velocity is to be set.
   * @param velocity - The new absolute velocity to set for the entity.
   */
  export function SetEntityAbsVelocity(number entityHandle, Vector3 velocity): void;

  /**
   * Retrieves the base velocity of an entity.
   * @param entityHandle - The handle of the entity whose base velocity is to be retrieved.
   * @returns A vector where the base velocity will be stored.
   */
  export function GetEntityBaseVelocity(number entityHandle): Vector3;

  /**
   * Retrieves the local angular velocity of an entity.
   * @param entityHandle - The handle of the entity whose local angular velocity is to be retrieved.
   * @returns A vector where the local angular velocity will be stored.
   */
  export function GetEntityLocalAngVelocity(number entityHandle): Vector3;

  /**
   * Retrieves the angular velocity of an entity.
   * @param entityHandle - The handle of the entity whose angular velocity is to be retrieved.
   * @returns A vector where the angular velocity will be stored.
   */
  export function GetEntityAngVelocity(number entityHandle): Vector3;

  /**
   * Sets the angular velocity of an entity.
   * @param entityHandle - The handle of the entity whose angular velocity is to be set.
   * @param velocity - The new angular velocity to set for the entity.
   */
  export function SetEntityAngVelocity(number entityHandle, Vector3 velocity): void;

  /**
   * Retrieves the local velocity of an entity.
   * @param entityHandle - The handle of the entity whose local velocity is to be retrieved.
   * @returns A vector where the local velocity will be stored.
   */
  export function GetEntityLocalVelocity(number entityHandle): Vector3;

  /**
   * Retrieves the angular rotation of an entity.
   * @param entityHandle - The handle of the entity whose angular rotation is to be retrieved.
   * @returns A vector where the angular rotation will be stored.
   */
  export function GetEntityAngRotation(number entityHandle): Vector3;

  /**
   * Sets the angular rotation of an entity.
   * @param entityHandle - The handle of the entity whose angular rotation is to be set.
   * @param rotation - The new angular rotation to set for the entity.
   */
  export function SetEntityAngRotation(number entityHandle, Vector3 rotation): void;

  /**
   * Returns the input Vector transformed from entity to world space.
   * @param entityHandle - The handle of the entity
   * @param point - Point in entity local space to transform
   * @returns The point transformed to world space coordinates
   */
  export function TransformPointEntityToWorld(number entityHandle, Vector3 point): Vector3;

  /**
   * Returns the input Vector transformed from world to entity space.
   * @param entityHandle - The handle of the entity
   * @param point - Point in world space to transform
   * @returns The point transformed to entity local space coordinates
   */
  export function TransformPointWorldToEntity(number entityHandle, Vector3 point): Vector3;

  /**
   * Get vector to eye position - absolute coords.
   * @param entityHandle - The handle of the entity
   * @returns Eye position in absolute/world coordinates
   */
  export function GetEntityEyePosition(number entityHandle): Vector3;

  /**
   * Get the qangles that this entity is looking at.
   * @param entityHandle - The handle of the entity
   * @returns Eye angles as a vector (pitch, yaw, roll)
   */
  export function GetEntityEyeAngles(number entityHandle): Vector3;

  /**
   * Sets the forward velocity of an entity.
   * @param entityHandle - The handle of the entity whose forward velocity is to be set.
   * @param forward
   */
  export function SetEntityForwardVector(number entityHandle, Vector3 forward): void;

  /**
   * Get the forward vector of the entity.
   * @param entityHandle - The handle of the entity to query
   * @returns Forward-facing direction vector of the entity
   */
  export function GetEntityForwardVector(number entityHandle): Vector3;

  /**
   * Get the left vector of the entity.
   * @param entityHandle - The handle of the entity to query
   * @returns Left-facing direction vector of the entity (aligned with the y axis)
   */
  export function GetEntityLeftVector(number entityHandle): Vector3;

  /**
   * Get the right vector of the entity.
   * @param entityHandle - The handle of the entity to query
   * @returns Right-facing direction vector of the entity
   */
  export function GetEntityRightVector(number entityHandle): Vector3;

  /**
   * Get the up vector of the entity.
   * @param entityHandle - The handle of the entity to query
   * @returns Up-facing direction vector of the entity
   */
  export function GetEntityUpVector(number entityHandle): Vector3;

  /**
   * Get the entity-to-world transformation matrix.
   * @param entityHandle - The handle of the entity to query
   * @returns 4x4 transformation matrix representing entity's position, rotation, and scale in world space
   */
  export function GetEntityTransform(number entityHandle): Matrix4x4;

  /**
   * Retrieves the model name of an entity.
   * @param entityHandle - The handle of the entity whose model name is to be retrieved.
   * @returns A string where the model name will be stored.
   */
  export function GetEntityModel(number entityHandle): string;

  /**
   * Sets the model name of an entity.
   * @param entityHandle - The handle of the entity whose model name is to be set.
   * @param model - The new model name to set for the entity.
   */
  export function SetEntityModel(number entityHandle, string model): void;

  /**
   * Retrieves the water level of an entity.
   * @param entityHandle - The handle of the entity whose water level is to be retrieved.
   * @returns The water level of the entity, or 0.0f if the entity is invalid.
   */
  export function GetEntityWaterLevel(number entityHandle): number;

  /**
   * Retrieves the ground entity of an entity.
   * @param entityHandle - The handle of the entity whose ground entity is to be retrieved.
   * @returns The handle of the ground entity, or INVALID_EHANDLE_INDEX if the entity is invalid.
   */
  export function GetEntityGroundEntity(number entityHandle): number;

  /**
   * Retrieves the effects of an entity.
   * @param entityHandle - The handle of the entity whose effects are to be retrieved.
   * @returns The effect flags of the entity, or 0 if the entity is invalid.
   */
  export function GetEntityEffects(number entityHandle): number;

  /**
   * Adds the render effect flag to an entity.
   * @param entityHandle - The handle of the entity to modify
   * @param effects - Render effect flags to add
   */
  export function AddEntityEffects(number entityHandle, number effects): void;

  /**
   * Removes the render effect flag from an entity.
   * @param entityHandle - The handle of the entity to modify
   * @param effects - Render effect flags to remove
   */
  export function RemoveEntityEffects(number entityHandle, number effects): void;

  /**
   * Get a vector containing max bounds, centered on object.
   * @param entityHandle - The handle of the entity to query
   * @returns Vector containing the maximum bounds of the entity's bounding box
   */
  export function GetEntityBoundingMaxs(number entityHandle): Vector3;

  /**
   * Get a vector containing min bounds, centered on object.
   * @param entityHandle - The handle of the entity to query
   * @returns Vector containing the minimum bounds of the entity's bounding box
   */
  export function GetEntityBoundingMins(number entityHandle): Vector3;

  /**
   * Get vector to center of object - absolute coords.
   * @param entityHandle - The handle of the entity to query
   * @returns Vector pointing to the center of the entity in absolute/world coordinates
   */
  export function GetEntityCenter(number entityHandle): Vector3;

  /**
   * Teleports an entity to a specified location and orientation.
   * @param entityHandle - The handle of the entity to teleport.
   * @param origin - A pointer to a Vector representing the new absolute position. Can be nullptr.
   * @param angles - A pointer to a QAngle representing the new orientation. Can be nullptr.
   * @param velocity - A pointer to a Vector representing the new velocity. Can be nullptr.
   */
  export function TeleportEntity(number entityHandle, bigint origin, bigint angles, bigint velocity): void;

  /**
   * Apply an absolute velocity impulse to an entity.
   * @param entityHandle - The handle of the entity to apply impulse to
   * @param vecImpulse - Velocity impulse vector to apply
   */
  export function ApplyAbsVelocityImpulseToEntity(number entityHandle, Vector3 vecImpulse): void;

  /**
   * Apply a local angular velocity impulse to an entity.
   * @param entityHandle - The handle of the entity to apply impulse to
   * @param angImpulse - Angular velocity impulse vector to apply
   */
  export function ApplyLocalAngularVelocityImpulseToEntity(number entityHandle, Vector3 angImpulse): void;

  /**
   * Invokes a named input method on a specified entity.
   * @param entityHandle - The handle of the target entity that will receive the input.
   * @param inputName - The name of the input action to invoke.
   * @param activatorHandle - The handle of the entity that initiated the sequence of actions.
   * @param callerHandle - The handle of the entity sending this event.
   * @param value - The value associated with the input action.
   * @param type - The type or classification of the value.
   * @param outputId - An identifier for tracking the output of this operation.
   */
  export function AcceptEntityInput(number entityHandle, string inputName, number activatorHandle, number callerHandle, any value, number type, number outputId): void;

  /**
   * Connects a script function to an entity output.
   * @param entityHandle - The handle of the entity.
   * @param output - The name of the output to connect to.
   * @param functionName - The name of the script function to call.
   */
  export function ConnectEntityOutput(number entityHandle, string output, string functionName): void;

  /**
   * Disconnects a script function from an entity output.
   * @param entityHandle - The handle of the entity.
   * @param output - The name of the output.
   * @param functionName - The name of the script function to disconnect.
   */
  export function DisconnectEntityOutput(number entityHandle, string output, string functionName): void;

  /**
   * Disconnects a script function from an I/O event on a different entity.
   * @param entityHandle - The handle of the calling entity.
   * @param output - The name of the output.
   * @param functionName - The function name to disconnect.
   * @param targetHandle - The handle of the entity whose output is being disconnected.
   */
  export function DisconnectEntityRedirectedOutput(number entityHandle, string output, string functionName, number targetHandle): void;

  /**
   * Fires an entity output.
   * @param entityHandle - The handle of the entity firing the output.
   * @param outputName - The name of the output to fire.
   * @param activatorHandle - The entity activating the output.
   * @param callerHandle - The entity that called the output.
   * @param value - The value associated with the input action.
   * @param type - The type or classification of the value.
   * @param delay - Delay in seconds before firing the output.
   */
  export function FireEntityOutput(number entityHandle, string outputName, number activatorHandle, number callerHandle, any value, number type, number delay): void;

  /**
   * Redirects an entity output to call a function on another entity.
   * @param entityHandle - The handle of the entity whose output is being redirected.
   * @param output - The name of the output to redirect.
   * @param functionName - The function name to call on the target entity.
   * @param targetHandle - The handle of the entity that will receive the output call.
   */
  export function RedirectEntityOutput(number entityHandle, string output, string functionName, number targetHandle): void;

  /**
   * Makes an entity follow another entity with optional bone merging.
   * @param entityHandle - The handle of the entity that will follow
   * @param attachmentHandle - The handle of the entity to follow
   * @param boneMerge - If true, bones will be merged between entities
   */
  export function FollowEntity(number entityHandle, number attachmentHandle, boolean boneMerge): void;

  /**
   * Makes an entity follow another entity and merge with a specific bone or attachment.
   * @param entityHandle - The handle of the entity that will follow
   * @param attachmentHandle - The handle of the entity to follow
   * @param boneOrAttachName - Name of the bone or attachment point to merge with
   */
  export function FollowEntityMerge(number entityHandle, number attachmentHandle, string boneOrAttachName): void;

  /**
   * Apply damage to an entity.
   * @param entityHandle - The handle of the entity receiving damage
   * @param inflictorHandle - The handle of the entity inflicting damage (e.g., projectile)
   * @param attackerHandle - The handle of the attacking entity
   * @param force - Direction and magnitude of force to apply
   * @param hitPos - Position where the damage hit occurred
   * @param damage - Amount of damage to apply
   * @param damageTypes - Bitfield of damage type flags
   * @returns Amount of damage actually applied to the entity
   */
  export function TakeEntityDamage(number entityHandle, number inflictorHandle, number attackerHandle, Vector3 force, Vector3 hitPos, number damage, number damageTypes): number;

  /**
   * Retrieves a float attribute value from an entity.
   * @param entityHandle - The handle of the entity.
   * @param name - The name of the attribute.
   * @param defaultValue - The default value to return if the attribute does not exist.
   * @returns The float value of the attribute, or the default value if missing or invalid.
   */
  export function GetEntityAttributeFloatValue(number entityHandle, string name, number defaultValue): number;

  /**
   * Retrieves an integer attribute value from an entity.
   * @param entityHandle - The handle of the entity.
   * @param name - The name of the attribute.
   * @param defaultValue - The default value to return if the attribute does not exist.
   * @returns The integer value of the attribute, or the default value if missing or invalid.
   */
  export function GetEntityAttributeIntValue(number entityHandle, string name, number defaultValue): number;

  /**
   * Sets a float attribute value on an entity.
   * @param entityHandle - The handle of the entity.
   * @param name - The name of the attribute.
   * @param value - The float value to assign to the attribute.
   */
  export function SetEntityAttributeFloatValue(number entityHandle, string name, number value): void;

  /**
   * Sets an integer attribute value on an entity.
   * @param entityHandle - The handle of the entity.
   * @param name - The name of the attribute.
   * @param value - The integer value to assign to the attribute.
   */
  export function SetEntityAttributeIntValue(number entityHandle, string name, number value): void;

  /**
   * Deletes an attribute from an entity.
   * @param entityHandle - The handle of the entity.
   * @param name - The name of the attribute to delete.
   */
  export function DeleteEntityAttribute(number entityHandle, string name): void;

  /**
   * Checks if an entity has a specific attribute.
   * @param entityHandle - The handle of the entity.
   * @param name - The name of the attribute to check.
   * @returns True if the attribute exists, false otherwise.
   */
  export function HasEntityAttribute(number entityHandle, string name): boolean;

  /**
   * Creates a hook for when a game event is fired.
   * @param name - The name of the event to hook.
   * @param callback - The callback function to call when the event is fired.
   * @param type - Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @returns An integer indicating the result of the hook operation.
   */
  export function HookEvent(string name, function callback, number type): number;

  /**
   * Removes a hook for when a game event is fired.
   * @param name - The name of the event to unhook.
   * @param callback - The callback function to remove.
   * @param type - Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @returns An integer indicating the result of the unhook operation.
   */
  export function UnhookEvent(string name, function callback, number type): number;

  /**
   * Creates a game event to be fired later.
   * @param name - The name of the event to create.
   * @param force - A boolean indicating whether to force the creation of the event.
   * @returns A pointer to the created EventInfo structure.
   */
  export function CreateEvent(string name, boolean force): bigint;

  /**
   * Fires a game event.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param dontBroadcast - A boolean indicating whether to broadcast the event.
   */
  export function FireEvent(bigint info, boolean dontBroadcast): void;

  /**
   * Fires a game event to a specific client.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param playerSlot - The index of the client to fire the event to.
   */
  export function FireEventToClient(bigint info, number playerSlot): void;

  /**
   * Cancels a previously created game event that has not been fired.
   * @param info - A pointer to the EventInfo structure of the event to cancel.
   */
  export function CancelCreatedEvent(bigint info): void;

  /**
   * Retrieves the boolean value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the boolean value.
   * @returns The boolean value associated with the key.
   */
  export function GetEventBool(bigint info, string key): boolean;

  /**
   * Retrieves the float value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the float value.
   * @returns The float value associated with the key.
   */
  export function GetEventFloat(bigint info, string key): number;

  /**
   * Retrieves the integer value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the integer value.
   * @returns The integer value associated with the key.
   */
  export function GetEventInt(bigint info, string key): number;

  /**
   * Retrieves the long integer value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the long integer value.
   * @returns The long integer value associated with the key.
   */
  export function GetEventUInt64(bigint info, string key): bigint;

  /**
   * Retrieves the string value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the string value.
   * @returns A string where the result will be stored.
   */
  export function GetEventString(bigint info, string key): string;

  /**
   * Retrieves the pointer value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the pointer value.
   * @returns The pointer value associated with the key.
   */
  export function GetEventPtr(bigint info, string key): bigint;

  /**
   * Retrieves the player controller address of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the player controller address.
   * @returns A pointer to the player controller associated with the key.
   */
  export function GetEventPlayerController(bigint info, string key): bigint;

  /**
   * Retrieves the player index of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the player index.
   * @returns The player index associated with the key.
   */
  export function GetEventPlayerIndex(bigint info, string key): number;

  /**
   * Retrieves the player pawn address of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the player pawn address.
   * @returns A pointer to the player pawn associated with the key.
   */
  export function GetEventPlayerPawn(bigint info, string key): bigint;

  /**
   * Retrieves the entity address of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the entity address.
   * @returns A pointer to the entity associated with the key.
   */
  export function GetEventEntity(bigint info, string key): bigint;

  /**
   * Retrieves the entity index of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the entity index.
   * @returns The entity index associated with the key.
   */
  export function GetEventEntityIndex(bigint info, string key): number;

  /**
   * Retrieves the entity handle of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to retrieve the entity handle.
   * @returns The entity handle associated with the key.
   */
  export function GetEventEntityHandle(bigint info, string key): number;

  /**
   * Retrieves the name of a game event.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @returns A string where the result will be stored.
   */
  export function GetEventName(bigint info): string;

  /**
   * Sets the boolean value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the boolean value.
   * @param value - The boolean value to set.
   */
  export function SetEventBool(bigint info, string key, boolean value): void;

  /**
   * Sets the floating point value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the float value.
   * @param value - The float value to set.
   */
  export function SetEventFloat(bigint info, string key, number value): void;

  /**
   * Sets the integer value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the integer value.
   * @param value - The integer value to set.
   */
  export function SetEventInt(bigint info, string key, number value): void;

  /**
   * Sets the long integer value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the long integer value.
   * @param value - The long integer value to set.
   */
  export function SetEventUInt64(bigint info, string key, bigint value): void;

  /**
   * Sets the string value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the string value.
   * @param value - The string value to set.
   */
  export function SetEventString(bigint info, string key, string value): void;

  /**
   * Sets the pointer value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the pointer value.
   * @param value - The pointer value to set.
   */
  export function SetEventPtr(bigint info, string key, bigint value): void;

  /**
   * Sets the player controller address of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the player controller address.
   * @param value - A pointer to the player controller to set.
   */
  export function SetEventPlayerController(bigint info, string key, bigint value): void;

  /**
   * Sets the player index value of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the player index value.
   * @param value - The player index value to set.
   */
  export function SetEventPlayerIndex(bigint info, string key, number value): void;

  /**
   * Sets the entity address of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the entity address.
   * @param value - A pointer to the entity to set.
   */
  export function SetEventEntity(bigint info, string key, bigint value): void;

  /**
   * Sets the entity index of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the entity index.
   * @param value - The entity index value to set.
   */
  export function SetEventEntityIndex(bigint info, string key, number value): void;

  /**
   * Sets the entity handle of a game event's key.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param key - The key for which to set the entity handle.
   * @param value - The entity handle value to set.
   */
  export function SetEventEntityHandle(bigint info, string key, number value): void;

  /**
   * Sets whether an event's broadcasting will be disabled or not.
   * @param info - A pointer to the EventInfo structure containing event data.
   * @param dontBroadcast - A boolean indicating whether to disable broadcasting.
   */
  export function SetEventBroadcast(bigint info, boolean dontBroadcast): void;

  /**
   * Load game event descriptions from a file (e.g., "resource/gameevents.res").
   * @param path - The path to the file containing event descriptions.
   * @param searchAll - A boolean indicating whether to search all paths for the file.
   * @returns An integer indicating the result of the loading operation.
   */
  export function LoadEventsFromFile(string path, boolean searchAll): number;

  /**
   * Closes a game configuration file.
   * @param id - An id to the game configuration to be closed.
   */
  export function CloseGameConfigFile(number id): void;

  /**
   * Loads a game configuration file.
   * @param paths - The paths to the game configuration file to be loaded.
   * @returns A id to the loaded game configuration object, or -1 if loading fails.
   */
  export function LoadGameConfigFile(string[] paths): number;

  /**
   * Retrieves a patch associated with the game configuration.
   * @param id - An id to the game configuration from which to retrieve the patch.
   * @param name - The name of the patch to be retrieved.
   * @returns A string where the patch will be stored.
   */
  export function GetGameConfigPatch(number id, string name): string;

  /**
   * Retrieves the offset associated with a name from the game configuration.
   * @param id - An id to the game configuration from which to retrieve the offset.
   * @param name - The name whose offset is to be retrieved.
   * @returns The offset associated with the specified name.
   */
  export function GetGameConfigOffset(number id, string name): number;

  /**
   * Retrieves the address associated with a name from the game configuration.
   * @param id - An id to the game configuration from which to retrieve the address.
   * @param name - The name whose address is to be retrieved.
   * @returns A pointer to the address associated with the specified name.
   */
  export function GetGameConfigAddress(number id, string name): bigint;

  /**
   * Retrieves a vtable associated with the game configuration.
   * @param id - An id to the game configuration from which to retrieve the vtable.
   * @param name - The name of the vtable to be retrieved.
   * @returns A pointer to the vtable associated with the specified name
   */
  export function GetGameConfigVTable(number id, string name): bigint;

  /**
   * Retrieves the signature associated with a name from the game configuration.
   * @param id - An id to the game configuration from which to retrieve the signature.
   * @param name - The name whose signature is to be resolved and retrieved.
   * @returns A pointer to the signature associated with the specified name.
   */
  export function GetGameConfigSignature(number id, string name): bigint;

  /**
   * Registers a new logging channel with specified properties.
   * @param name - The name of the logging channel.
   * @param iFlags - Flags associated with the logging channel.
   * @param verbosity - The verbosity level for the logging channel.
   * @param color - The color for messages logged to this channel.
   * @returns The ID of the newly created logging channel.
   */
  export function RegisterLoggingChannel(string name, number iFlags, number verbosity, number color): number;

  /**
   * Adds a tag to a specified logging channel.
   * @param channelID - The ID of the logging channel to which the tag will be added.
   * @param tagName - The name of the tag to add to the channel.
   */
  export function AddLoggerTagToChannel(number channelID, string tagName): void;

  /**
   * Checks if a specified tag exists in a logging channel.
   * @param channelID - The ID of the logging channel.
   * @param tag - The name of the tag to check for.
   * @returns True if the tag exists in the channel, otherwise false.
   */
  export function HasLoggerTag(number channelID, string tag): boolean;

  /**
   * Checks if a logging channel is enabled based on severity.
   * @param channelID - The ID of the logging channel.
   * @param severity - The severity of a logging operation.
   * @returns True if the channel is enabled for the specified severity, otherwise false.
   */
  export function IsLoggerChannelEnabledBySeverity(number channelID, number severity): boolean;

  /**
   * Checks if a logging channel is enabled based on verbosity.
   * @param channelID - The ID of the logging channel.
   * @param verbosity - The verbosity level to check.
   * @returns True if the channel is enabled for the specified verbosity, otherwise false.
   */
  export function IsLoggerChannelEnabledByVerbosity(number channelID, number verbosity): boolean;

  /**
   * Retrieves the verbosity level of a logging channel.
   * @param channelID - The ID of the logging channel.
   * @returns The verbosity level of the specified logging channel.
   */
  export function GetLoggerChannelVerbosity(number channelID): number;

  /**
   * Sets the verbosity level of a logging channel.
   * @param channelID - The ID of the logging channel.
   * @param verbosity - The new verbosity level to set.
   */
  export function SetLoggerChannelVerbosity(number channelID, number verbosity): void;

  /**
   * Sets the verbosity level of a logging channel by name.
   * @param channelID - The ID of the logging channel.
   * @param name - The name of the logging channel.
   * @param verbosity - The new verbosity level to set.
   */
  export function SetLoggerChannelVerbosityByName(number channelID, string name, number verbosity): void;

  /**
   * Sets the verbosity level of a logging channel by tag.
   * @param channelID - The ID of the logging channel.
   * @param tag - The name of the tag.
   * @param verbosity - The new verbosity level to set.
   */
  export function SetLoggerChannelVerbosityByTag(number channelID, string tag, number verbosity): void;

  /**
   * Retrieves the color setting of a logging channel.
   * @param channelID - The ID of the logging channel.
   * @returns The color value of the specified logging channel.
   */
  export function GetLoggerChannelColor(number channelID): number;

  /**
   * Sets the color setting of a logging channel.
   * @param channelID - The ID of the logging channel.
   * @param color - The new color value to set for the channel.
   */
  export function SetLoggerChannelColor(number channelID, number color): void;

  /**
   * Retrieves the flags of a logging channel.
   * @param channelID - The ID of the logging channel.
   * @returns The flags of the specified logging channel.
   */
  export function GetLoggerChannelFlags(number channelID): number;

  /**
   * Sets the flags of a logging channel.
   * @param channelID - The ID of the logging channel.
   * @param eFlags - The new flags to set for the channel.
   */
  export function SetLoggerChannelFlags(number channelID, number eFlags): void;

  /**
   * Logs a message to a specified channel with a severity level.
   * @param channelID - The ID of the logging channel.
   * @param severity - The severity level for the log message.
   * @param message - The message to log.
   * @returns An integer indicating the result of the logging operation.
   */
  export function Log(number channelID, number severity, string message): number;

  /**
   * Logs a colored message to a specified channel with a severity level.
   * @param channelID - The ID of the logging channel.
   * @param severity - The severity level for the log message.
   * @param color - The color for the log message.
   * @param message - The message to log.
   * @returns An integer indicating the result of the logging operation.
   */
  export function LogColored(number channelID, number severity, number color, string message): number;

  /**
   * Logs a detailed message to a specified channel, including source code info.
   * @param channelID - The ID of the logging channel.
   * @param severity - The severity level for the log message.
   * @param file - The file name where the log call occurred.
   * @param line - The line number where the log call occurred.
   * @param function_ - The name of the function where the log call occurred.
   * @param message - The message to log.
   * @returns An integer indicating the result of the logging operation.
   */
  export function LogFull(number channelID, number severity, string file, number line, string function_, string message): number;

  /**
   * Logs a detailed colored message to a specified channel, including source code info.
   * @param channelID - The ID of the logging channel.
   * @param severity - The severity level for the log message.
   * @param file - The file name where the log call occurred.
   * @param line - The line number where the log call occurred.
   * @param function_ - The name of the function where the log call occurred.
   * @param color - The color for the log message.
   * @param message - The message to log.
   * @returns An integer indicating the result of the logging operation.
   */
  export function LogFullColored(number channelID, number severity, string file, number line, string function_, number color, string message): number;

  /**
   * Retrieves the attachment angles of an entity.
   * @param entityHandle - The handle of the entity whose attachment angles are to be retrieved.
   * @param attachmentIndex - The attachment index.
   * @returns A vector representing the attachment angles (pitch, yaw, roll).
   */
  export function GetEntityAttachmentAngles(number entityHandle, number attachmentIndex): Vector3;

  /**
   * Retrieves the forward vector of an entity's attachment.
   * @param entityHandle - The handle of the entity whose attachment forward vector is to be retrieved.
   * @param attachmentIndex - The attachment index.
   * @returns A vector representing the forward direction of the attachment.
   */
  export function GetEntityAttachmentForward(number entityHandle, number attachmentIndex): Vector3;

  /**
   * Retrieves the origin vector of an entity's attachment.
   * @param entityHandle - The handle of the entity whose attachment origin is to be retrieved.
   * @param attachmentIndex - The attachment index.
   * @returns A vector representing the origin of the attachment.
   */
  export function GetEntityAttachmentOrigin(number entityHandle, number attachmentIndex): Vector3;

  /**
   * Retrieves the material group hash of an entity.
   * @param entityHandle - The handle of the entity.
   * @returns The material group hash.
   */
  export function GetEntityMaterialGroupHash(number entityHandle): number;

  /**
   * Retrieves the material group mask of an entity.
   * @param entityHandle - The handle of the entity.
   * @returns The mesh group mask.
   */
  export function GetEntityMaterialGroupMask(number entityHandle): bigint;

  /**
   * Retrieves the model scale of an entity.
   * @param entityHandle - The handle of the entity.
   * @returns The model scale factor.
   */
  export function GetEntityModelScale(number entityHandle): number;

  /**
   * Retrieves the render alpha of an entity.
   * @param entityHandle - The handle of the entity.
   * @returns The alpha modulation value.
   */
  export function GetEntityRenderAlpha(number entityHandle): number;

  /**
   * Retrieves the render color of an entity.
   * @param entityHandle - The handle of the entity.
   * @returns A vector representing the render color (R, G, B).
   */
  export function GetEntityRenderColor2(number entityHandle): Vector3;

  /**
   * Retrieves an attachment index by name.
   * @param entityHandle - The handle of the entity.
   * @param attachmentName - The name of the attachment.
   * @returns The attachment index, or -1 if not found.
   */
  export function ScriptLookupAttachment(number entityHandle, string attachmentName): number;

  /**
   * Sets a bodygroup value by index.
   * @param entityHandle - The handle of the entity.
   * @param group - The bodygroup index.
   * @param value - The new value to set for the bodygroup.
   */
  export function SetEntityBodygroup(number entityHandle, number group, number value): void;

  /**
   * Sets a bodygroup value by name.
   * @param entityHandle - The handle of the entity.
   * @param name - The bodygroup name.
   * @param value - The new value to set for the bodygroup.
   */
  export function SetEntityBodygroupByName(number entityHandle, string name, number value): void;

  /**
   * Sets the light group of an entity.
   * @param entityHandle - The handle of the entity.
   * @param lightGroup - The light group name.
   */
  export function SetEntityLightGroup(number entityHandle, string lightGroup): void;

  /**
   * Sets the material group of an entity.
   * @param entityHandle - The handle of the entity.
   * @param materialGroup - The material group name.
   */
  export function SetEntityMaterialGroup(number entityHandle, string materialGroup): void;

  /**
   * Sets the material group hash of an entity.
   * @param entityHandle - The handle of the entity.
   * @param hash - The new material group hash to set.
   */
  export function SetEntityMaterialGroupHash(number entityHandle, number hash): void;

  /**
   * Sets the material group mask of an entity.
   * @param entityHandle - The handle of the entity.
   * @param mask - The new mesh group mask to set.
   */
  export function SetEntityMaterialGroupMask(number entityHandle, bigint mask): void;

  /**
   * Sets the model scale of an entity.
   * @param entityHandle - The handle of the entity.
   * @param scale - The new scale factor.
   */
  export function SetEntityModelScale(number entityHandle, number scale): void;

  /**
   * Sets the render alpha of an entity.
   * @param entityHandle - The handle of the entity.
   * @param alpha - The new alpha value (0â€“255).
   */
  export function SetEntityRenderAlpha(number entityHandle, number alpha): void;

  /**
   * Sets the render color of an entity.
   * @param entityHandle - The handle of the entity.
   * @param r - The red component (0â€“255).
   * @param g - The green component (0â€“255).
   * @param b - The blue component (0â€“255).
   */
  export function SetEntityRenderColor2(number entityHandle, number r, number g, number b): void;

  /**
   * Sets the render mode of an entity.
   * @param entityHandle - The handle of the entity.
   * @param mode - The new render mode value.
   */
  export function SetEntityRenderMode2(number entityHandle, number mode): void;

  /**
   * Sets a single mesh group for an entity.
   * @param entityHandle - The handle of the entity.
   * @param meshGroupName - The name of the mesh group.
   */
  export function SetEntitySingleMeshGroup(number entityHandle, string meshGroupName): void;

  /**
   * Sets the size (bounding box) of an entity.
   * @param entityHandle - The handle of the entity.
   * @param mins - The minimum bounding box vector.
   * @param maxs - The maximum bounding box vector.
   */
  export function SetEntitySize(number entityHandle, Vector3 mins, Vector3 maxs): void;

  /**
   * Sets the skin of an entity.
   * @param entityHandle - The handle of the entity.
   * @param skin - The new skin index.
   */
  export function SetEntitySkin(number entityHandle, number skin): void;

  /**
   * Start a new Yes/No vote
   * @param duration - Maximum time to leave vote active for
   * @param caller - Player slot of the vote caller. Use VOTE_CALLER_SERVER for 'Server'.
   * @param voteTitle - Translation string to use as the vote message. (Only '#SFUI_vote' or '#Panorama_vote' strings)
   * @param detailStr - Extra string used in some vote translation strings.
   * @param votePassTitle - Translation string to use as the vote message. (Only '#SFUI_vote' or '#Panorama_vote' strings)
   * @param detailPassStr - Extra string used in some vote translation strings when the vote passes.
   * @param failReason - Reason for the vote to fail, used in some translation strings.
   * @param filter - Recipient filter with all the clients who are allowed to participate in the vote.
   * @param result - Called when a menu action is completed.
   * @param handler - Called when the vote has finished.
   */
  export function PanoramaSendYesNoVote(number duration, number caller, string voteTitle, string detailStr, string votePassTitle, string detailPassStr, number failReason, bigint filter, function result, function handler): boolean;

  /**
   * Start a new Yes/No vote with all players included
   * @param duration - Maximum time to leave vote active for
   * @param caller - Player slot of the vote caller. Use VOTE_CALLER_SERVER for 'Server'.
   * @param voteTitle - Translation string to use as the vote message. (Only '#SFUI_vote' or '#Panorama_vote' strings)
   * @param detailStr - Extra string used in some vote translation strings.
   * @param votePassTitle - Translation string to use as the vote message. (Only '#SFUI_vote' or '#Panorama_vote' strings)
   * @param detailPassStr - Extra string used in some vote translation strings when the vote passes.
   * @param failReason - Reason for the vote to fail, used in some translation strings.
   * @param result - Called when a menu action is completed.
   * @param handler - Called when the vote has finished.
   */
  export function PanoramaSendYesNoVoteToAll(number duration, number caller, string voteTitle, string detailStr, string votePassTitle, string detailPassStr, number failReason, function result, function handler): boolean;

  /**
   * Removes a player from the current vote.
   * @param playerSlot - The slot/index of the player to remove from the vote.
   */
  export function PanoramaRemovePlayerFromVote(number playerSlot): void;

  /**
   * Checks if a player is in the vote pool.
   * @param playerSlot - The slot/index of the player to check.
   * @returns true if the player is in the vote pool, false otherwise.
   */
  export function PanoramaIsPlayerInVotePool(number playerSlot): boolean;

  /**
   * Redraws the vote UI to a specific player client.
   * @param playerSlot - The slot/index of the player to update.
   * @returns true if the vote UI was successfully redrawn, false otherwise.
   */
  export function PanoramaRedrawVoteToClient(number playerSlot): boolean;

  /**
   * Checks if a vote is currently in progress.
   * @returns true if a vote is active, false otherwise.
   */
  export function PanoramaIsVoteInProgress(): boolean;

  /**
   * Ends the current vote with a specified reason.
   * @param reason - The reason for ending the vote.
   */
  export function PanoramaEndVote(number reason): void;

  /**
   * Get the offset of a member in a given schema class.
   * @param className - The name of the class.
   * @param memberName - The name of the member whose offset is to be retrieved.
   * @returns The offset of the member in the class.
   */
  export function GetSchemaOffset(string className, string memberName): number;

  /**
   * Get the offset of a chain in a given schema class.
   * @param className - The name of the class.
   * @returns The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function GetSchemaChainOffset(string className): number;

  /**
   * Check if a schema field is networked.
   * @param className - The name of the class.
   * @param memberName - The name of the member to check.
   * @returns True if the member is networked, false otherwise.
   */
  export function IsSchemaFieldNetworked(string className, string memberName): boolean;

  /**
   * Get the size of a schema class.
   * @param className - The name of the class.
   * @returns The size of the class in bytes, or -1 if the class is not found.
   */
  export function GetSchemaClassSize(string className): number;

  /**
   * Peeks into an entity's object schema and retrieves the integer value at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @param size - Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @returns The integer value at the given memory location.
   */
  export function GetEntData2(bigint entity, number offset, number size): number;

  /**
   * Peeks into an entity's object data and sets the integer value at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @param value - The integer value to set.
   * @param size - Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntData2(bigint entity, number offset, number value, number size, boolean changeState, number chainOffset): void;

  /**
   * Peeks into an entity's object schema and retrieves the float value at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @param size - Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @returns The float value at the given memory location.
   */
  export function GetEntDataFloat2(bigint entity, number offset, number size): number;

  /**
   * Peeks into an entity's object data and sets the float value at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @param value - The float value to set.
   * @param size - Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntDataFloat2(bigint entity, number offset, number value, number size, boolean changeState, number chainOffset): void;

  /**
   * Peeks into an entity's object schema and retrieves the string value at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @returns The string value at the given memory location.
   */
  export function GetEntDataString2(bigint entity, number offset): string;

  /**
   * Peeks into an entity's object data and sets the string at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @param value - The string value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntDataString2(bigint entity, number offset, string value, boolean changeState, number chainOffset): void;

  /**
   * Peeks into an entity's object schema and retrieves the vector value at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @returns The vector value at the given memory location.
   */
  export function GetEntDataVector2(bigint entity, number offset): Vector3;

  /**
   * Peeks into an entity's object data and sets the vector at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @param value - The vector value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntDataVector2(bigint entity, number offset, Vector3 value, boolean changeState, number chainOffset): void;

  /**
   * Peeks into an entity's object data and retrieves the entity handle at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @returns The entity handle at the given memory location.
   */
  export function GetEntDataEnt2(bigint entity, number offset): number;

  /**
   * Peeks into an entity's object data and sets the entity handle at the given offset.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @param value - The entity handle to set.
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntDataEnt2(bigint entity, number offset, number value, boolean changeState, number chainOffset): void;

  /**
   * Updates the networked state of a schema field for a given entity pointer.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param offset - The offset of the schema to use.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function ChangeEntityState2(bigint entity, number offset, number chainOffset): void;

  /**
   * Peeks into an entity's object schema and retrieves the integer value at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @param size - Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @returns The integer value at the given memory location.
   */
  export function GetEntData(number entityHandle, number offset, number size): number;

  /**
   * Peeks into an entity's object data and sets the integer value at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @param value - The integer value to set.
   * @param size - Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntData(number entityHandle, number offset, number value, number size, boolean changeState, number chainOffset): void;

  /**
   * Peeks into an entity's object schema and retrieves the float value at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @param size - Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @returns The float value at the given memory location.
   */
  export function GetEntDataFloat(number entityHandle, number offset, number size): number;

  /**
   * Peeks into an entity's object data and sets the float value at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @param value - The float value to set.
   * @param size - Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntDataFloat(number entityHandle, number offset, number value, number size, boolean changeState, number chainOffset): void;

  /**
   * Peeks into an entity's object schema and retrieves the string value at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @returns The string value at the given memory location.
   */
  export function GetEntDataString(number entityHandle, number offset): string;

  /**
   * Peeks into an entity's object data and sets the string at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @param value - The string value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntDataString(number entityHandle, number offset, string value, boolean changeState, number chainOffset): void;

  /**
   * Peeks into an entity's object schema and retrieves the vector value at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @returns The vector value at the given memory location.
   */
  export function GetEntDataVector(number entityHandle, number offset): Vector3;

  /**
   * Peeks into an entity's object data and sets the vector at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @param value - The vector value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntDataVector(number entityHandle, number offset, Vector3 value, boolean changeState, number chainOffset): void;

  /**
   * Peeks into an entity's object data and retrieves the entity handle at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @returns The entity handle at the given memory location.
   */
  export function GetEntDataEnt(number entityHandle, number offset): number;

  /**
   * Peeks into an entity's object data and sets the entity handle at the given offset.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @param value - The entity handle to set.
   * @param changeState - If true, change will be sent over the network.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function SetEntDataEnt(number entityHandle, number offset, number value, boolean changeState, number chainOffset): void;

  /**
   * Updates the networked state of a schema field for a given entity handle.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param offset - The offset of the schema to use.
   * @param chainOffset - The offset of the chain entity in the class (-1 for non-entity classes).
   */
  export function ChangeEntityState(number entityHandle, number offset, number chainOffset): void;

  /**
   * Retrieves the count of values that an entity schema's array can store.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @returns Size of array (in elements) or 0 if schema is not an array.
   */
  export function GetEntSchemaArraySize2(bigint entity, string className, string memberName): number;

  /**
   * Retrieves an integer value from an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns An integer value at the given schema offset.
   */
  export function GetEntSchema2(bigint entity, string className, string memberName, number element): number;

  /**
   * Sets an integer value in an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The integer value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchema2(bigint entity, string className, string memberName, number value, boolean changeState, number element): void;

  /**
   * Retrieves a float value from an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A float value at the given schema offset.
   */
  export function GetEntSchemaFloat2(bigint entity, string className, string memberName, number element): number;

  /**
   * Sets a float value in an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The float value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaFloat2(bigint entity, string className, string memberName, number value, boolean changeState, number element): void;

  /**
   * Retrieves a string value from an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A string value at the given schema offset.
   */
  export function GetEntSchemaString2(bigint entity, string className, string memberName, number element): string;

  /**
   * Sets a string value in an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The string value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaString2(bigint entity, string className, string memberName, string value, boolean changeState, number element): void;

  /**
   * Retrieves a vector value from an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A vector value at the given schema offset.
   */
  export function GetEntSchemaVector3D2(bigint entity, string className, string memberName, number element): Vector3;

  /**
   * Sets a vector value in an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The vector value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaVector3D2(bigint entity, string className, string memberName, Vector3 value, boolean changeState, number element): void;

  /**
   * Retrieves a vector value from an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A vector value at the given schema offset.
   */
  export function GetEntSchemaVector2D2(bigint entity, string className, string memberName, number element): Vector2;

  /**
   * Sets a vector value in an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The vector value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaVector2D2(bigint entity, string className, string memberName, Vector2 value, boolean changeState, number element): void;

  /**
   * Retrieves a vector value from an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A vector value at the given schema offset.
   */
  export function GetEntSchemaVector4D2(bigint entity, string className, string memberName, number element): Vector4;

  /**
   * Sets a vector value in an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The vector value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaVector4D2(bigint entity, string className, string memberName, Vector4 value, boolean changeState, number element): void;

  /**
   * Retrieves an entity handle from an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A string value at the given schema offset.
   */
  export function GetEntSchemaEnt2(bigint entity, string className, string memberName, number element): number;

  /**
   * Sets an entity handle in an entity's schema.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The entity handle to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaEnt2(bigint entity, string className, string memberName, number value, boolean changeState, number element): void;

  /**
   * Updates the networked state of a schema field for a given entity pointer.
   * @param entity - Pointer to the instance of the class where the value is to be set.
   * @param className - The name of the class that contains the member.
   * @param memberName - The name of the member to be set.
   */
  export function NetworkStateChanged2(bigint entity, string className, string memberName): void;

  /**
   * Retrieves the count of values that an entity schema's array can store.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @returns Size of array (in elements) or 0 if schema is not an array.
   */
  export function GetEntSchemaArraySize(number entityHandle, string className, string memberName): number;

  /**
   * Retrieves an integer value from an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns An integer value at the given schema offset.
   */
  export function GetEntSchema(number entityHandle, string className, string memberName, number element): number;

  /**
   * Sets an integer value in an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The integer value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchema(number entityHandle, string className, string memberName, number value, boolean changeState, number element): void;

  /**
   * Retrieves a float value from an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A float value at the given schema offset.
   */
  export function GetEntSchemaFloat(number entityHandle, string className, string memberName, number element): number;

  /**
   * Sets a float value in an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The float value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaFloat(number entityHandle, string className, string memberName, number value, boolean changeState, number element): void;

  /**
   * Retrieves a string value from an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A string value at the given schema offset.
   */
  export function GetEntSchemaString(number entityHandle, string className, string memberName, number element): string;

  /**
   * Sets a string value in an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The string value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaString(number entityHandle, string className, string memberName, string value, boolean changeState, number element): void;

  /**
   * Retrieves a vector value from an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A string value at the given schema offset.
   */
  export function GetEntSchemaVector3D(number entityHandle, string className, string memberName, number element): Vector3;

  /**
   * Sets a vector value in an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The vector value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaVector3D(number entityHandle, string className, string memberName, Vector3 value, boolean changeState, number element): void;

  /**
   * Retrieves a vector value from an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A string value at the given schema offset.
   */
  export function GetEntSchemaVector2D(number entityHandle, string className, string memberName, number element): Vector2;

  /**
   * Sets a vector value in an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The vector value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaVector2D(number entityHandle, string className, string memberName, Vector2 value, boolean changeState, number element): void;

  /**
   * Retrieves a vector value from an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A string value at the given schema offset.
   */
  export function GetEntSchemaVector4D(number entityHandle, string className, string memberName, number element): Vector4;

  /**
   * Sets a vector value in an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The vector value to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaVector4D(number entityHandle, string className, string memberName, Vector4 value, boolean changeState, number element): void;

  /**
   * Retrieves an entity handle from an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param element - Element # (starting from 0) if schema is an array.
   * @returns A string value at the given schema offset.
   */
  export function GetEntSchemaEnt(number entityHandle, string className, string memberName, number element): number;

  /**
   * Sets an entity handle in an entity's schema.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class.
   * @param memberName - The name of the schema member.
   * @param value - The entity handle to set.
   * @param changeState - If true, change will be sent over the network.
   * @param element - Element # (starting from 0) if schema is an array.
   */
  export function SetEntSchemaEnt(number entityHandle, string className, string memberName, number value, boolean changeState, number element): void;

  /**
   * Updates the networked state of a schema field for a given entity handle.
   * @param entityHandle - The handle of the entity from which the value is to be retrieved.
   * @param className - The name of the class that contains the member.
   * @param memberName - The name of the member to be set.
   */
  export function NetworkStateChanged(number entityHandle, string className, string memberName): void;

  /**
   * Creates a new timer that executes a callback function at specified delays.
   * @param delay - The time delay in seconds between each callback execution.
   * @param callback - The function to be called when the timer expires.
   * @param flags - Flags that modify the behavior of the timer (e.g., no-map change, repeating).
   * @param userData - An array intended to hold user-related data, allowing for elements of any type.
   * @returns A id to the newly created Timer object, or -1 if the timer could not be created.
   */
  export function CreateTimer(number delay, function callback, number flags, any[] userData): number;

  /**
   * Stops and removes an existing timer.
   * @param timer - A id of the Timer object to be stopped and removed.
   */
  export function KillsTimer(number timer): void;

  /**
   * Reschedules an existing timer with a new delay.
   * @param timer - A id of the Timer object to be stopped and removed.
   * @param newDaly - The new delay in seconds between each callback execution.
   */
  export function RescheduleTimer(number timer, number newDaly): void;

  /**
   * Returns the number of seconds in between game server ticks.
   * @returns The tick interval value.
   */
  export function GetTickInterval(): number;

  /**
   * Returns the simulated game time.
   * @returns The ticked time value.
   */
  export function GetTickedTime(): number;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientConnect_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientConnect_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientConnect_Post_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientConnect_Post_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientConnected_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientConnected_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientPutInServer_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientPutInServer_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientDisconnect_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientDisconnect_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientDisconnect_Post_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientDisconnect_Post_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientActive_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientActive_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientFullyConnect_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientFullyConnect_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientSettingsChanged_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientSettingsChanged_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnClientAuthenticated_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnClientAuthenticated_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnRoundTerminated_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnRoundTerminated_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnEntityCreated_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnEntityCreated_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnEntityDeleted_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnEntityDeleted_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnEntityParentChanged_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnEntityParentChanged_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnServerStartup_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnServerStartup_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnServerActivate_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnServerActivate_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnServerSpawn_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnServerSpawn_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnServerStarted_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnServerStarted_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnMapStart_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnMapStart_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnMapEnd_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnMapEnd_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnGameFrame_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnGameFrame_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnUpdateWhenNotInGame_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnUpdateWhenNotInGame_Unregister(function callback): void;

  /**
   * Register callback to event.
   * @param callback - Function callback.
   */
  export function OnPreWorldUpdate_Register(function callback): void;

  /**
   * Unregister callback to event.
   * @param callback - Function callback.
   */
  export function OnPreWorldUpdate_Unregister(function callback): void;

  /**
   * Retrieves the pointer to the current game rules proxy instance.
   * @returns A pointer to the game rules entity instance.
   */
  export function GetGameRulesProxy(): bigint;

  /**
   * Retrieves the pointer to the current game rules instance.
   * @returns A pointer to the game rules object.
   */
  export function GetGameRules(): bigint;

  /**
   * Retrieves the team manager instance for a specified team.
   * @param team - The numeric identifier of the team.
   * @returns A pointer to the corresponding CTeam instance, or nullptr if the team was not found.
   */
  export function GetGameTeamManager(number team): bigint;

  /**
   * Retrieves the current score of a specified team.
   * @param team - The numeric identifier of the team.
   * @returns The current score of the team, or -1 if the team could not be found.
   */
  export function GetGameTeamScore(number team): number;

  /**
   * Retrieves the number of players on a specified team.
   * @param team - The numeric identifier of the team (e.g., CS_TEAM_T, CS_TEAM_CT, CS_TEAM_SPECTATOR).
   * @returns The number of players on the team, or -1 if game rules are unavailable.
   */
  export function GetGamePlayerCount(number team): number;

  /**
   * Returns the total number of rounds played in the current match.
   * @returns The total number of rounds played, or -1 if the game rules are unavailable.
   */
  export function GetGameTotalRoundsPlayed(): number;

  /**
   * Forces the round to end with a specified reason and delay.
   * @param delay - Time (in seconds) to delay before the next round starts.
   * @param reason - The reason for ending the round, defined by the CSRoundEndReason enum.
   */
  export function TerminateRound(number delay, number reason): void;

  /**
   * Hooks a user message with a callback.
   * @param messageId - The ID of the message to hook.
   * @param callback - The callback function to invoke when the message is received.
   * @param mode - Whether to hook the message in the post mode (after processing) or pre mode (before processing).
   * @returns True if the hook was successfully added, false otherwise.
   */
  export function HookUserMessage(number messageId, function callback, number mode): boolean;

  /**
   * Unhooks a previously hooked user message.
   * @param messageId - The ID of the message to unhook.
   * @param callback - The callback function to remove.
   * @param mode - Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @returns True if the hook was successfully removed, false otherwise.
   */
  export function UnhookUserMessage(number messageId, function callback, number mode): boolean;

  /**
   * Creates a UserMessage from a serializable message.
   * @param msgSerializable - The serializable message.
   * @param message - The network message.
   * @param recipientMask - The recipient mask.
   * @returns A pointer to the newly created UserMessage.
   */
  export function UserMessageCreateFromSerializable(bigint msgSerializable, bigint message, bigint recipientMask): bigint;

  /**
   * Creates a UserMessage from a message name.
   * @param messageName - The name of the message.
   * @returns A pointer to the newly created UserMessage.
   */
  export function UserMessageCreateFromName(string messageName): bigint;

  /**
   * Creates a UserMessage from a message ID.
   * @param messageId - The ID of the message.
   * @returns A pointer to the newly created UserMessage.
   */
  export function UserMessageCreateFromId(number messageId): bigint;

  /**
   * Destroys a UserMessage and frees its memory.
   * @param userMessage - The UserMessage to destroy.
   */
  export function UserMessageDestroy(bigint userMessage): void;

  /**
   * Sends a UserMessage to the specified recipients.
   * @param userMessage - The UserMessage to send.
   */
  export function UserMessageSend(bigint userMessage): void;

  /**
   * Gets the name of the message.
   * @param userMessage - The UserMessage instance.
   * @returns The name of the message as a string.
   */
  export function UserMessageGetMessageName(bigint userMessage): string;

  /**
   * Gets the ID of the message.
   * @param userMessage - The UserMessage instance.
   * @returns The ID of the message.
   */
  export function UserMessageGetMessageID(bigint userMessage): number;

  /**
   * Checks if the message has a specific field.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field to check.
   * @returns True if the field exists, false otherwise.
   */
  export function UserMessageHasField(bigint userMessage, string fieldName): boolean;

  /**
   * Gets the protobuf message associated with the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @returns A pointer to the protobuf message.
   */
  export function UserMessageGetProtobufMessage(bigint userMessage): bigint;

  /**
   * Gets the serializable message associated with the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @returns A pointer to the serializable message.
   */
  export function UserMessageGetSerializableMessage(bigint userMessage): bigint;

  /**
   * Finds a message ID by its name.
   * @param messageName - The name of the message.
   * @returns The ID of the message, or 0 if the message was not found.
   */
  export function UserMessageFindMessageIdByName(string messageName): number;

  /**
   * Gets the recipient mask for the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @returns The recipient mask.
   */
  export function UserMessageGetRecipientMask(bigint userMessage): bigint;

  /**
   * Adds a single recipient (player) to the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param playerSlot - The slot index of the player to add as a recipient.
   */
  export function UserMessageAddRecipient(bigint userMessage, number playerSlot): void;

  /**
   * Adds all connected players as recipients to the UserMessage.
   * @param userMessage - The UserMessage instance.
   */
  export function UserMessageAddAllPlayers(bigint userMessage): void;

  /**
   * Sets the recipient mask for the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param mask - The recipient mask to set.
   */
  export function UserMessageSetRecipientMask(bigint userMessage, bigint mask): void;

  /**
   * Gets a nested message from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param message - A pointer to store the retrieved message.
   * @returns True if the message was successfully retrieved, false otherwise.
   */
  export function UserMessageGetMessage(bigint userMessage, string fieldName, bigint message): boolean;

  /**
   * Gets a repeated nested message from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param message - A pointer to store the retrieved message.
   * @returns True if the message was successfully retrieved, false otherwise.
   */
  export function UserMessageGetRepeatedMessage(bigint userMessage, string fieldName, number index, bigint message): boolean;

  /**
   * Adds a nested message to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param message - A pointer to the message to add.
   * @returns True if the message was successfully added, false otherwise.
   */
  export function UserMessageAddMessage(bigint userMessage, string fieldName, bigint message): boolean;

  /**
   * Gets the count of repeated fields in a field of the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @returns The count of repeated fields, or -1 if the field is not repeated or does not exist.
   */
  export function UserMessageGetRepeatedFieldCount(bigint userMessage, string fieldName): number;

  /**
   * Removes a value from a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the value to remove.
   * @returns True if the value was successfully removed, false otherwise.
   */
  export function UserMessageRemoveRepeatedFieldValue(bigint userMessage, string fieldName, number index): boolean;

  /**
   * Gets the debug string representation of the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @returns The debug string as a string.
   */
  export function UserMessageGetDebugString(bigint userMessage): string;

  /**
   * Reads an enum value from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The integer representation of the enum value, or 0 if invalid.
   */
  export function PbReadEnum(bigint userMessage, string fieldName, number index): number;

  /**
   * Reads a 32-bit integer from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The int32_t value read, or 0 if invalid.
   */
  export function PbReadInt32(bigint userMessage, string fieldName, number index): number;

  /**
   * Reads a 64-bit integer from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The int64_t value read, or 0 if invalid.
   */
  export function PbReadInt64(bigint userMessage, string fieldName, number index): number;

  /**
   * Reads an unsigned 32-bit integer from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The uint32_t value read, or 0 if invalid.
   */
  export function PbReadUInt32(bigint userMessage, string fieldName, number index): number;

  /**
   * Reads an unsigned 64-bit integer from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The uint64_t value read, or 0 if invalid.
   */
  export function PbReadUInt64(bigint userMessage, string fieldName, number index): bigint;

  /**
   * Reads a floating-point value from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The float value read, or 0.0 if invalid.
   */
  export function PbReadFloat(bigint userMessage, string fieldName, number index): number;

  /**
   * Reads a double-precision floating-point value from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The double value read, or 0.0 if invalid.
   */
  export function PbReadDouble(bigint userMessage, string fieldName, number index): number;

  /**
   * Reads a boolean value from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The boolean value read, or false if invalid.
   */
  export function PbReadBool(bigint userMessage, string fieldName, number index): boolean;

  /**
   * Reads a string from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The string value read, or an empty string if invalid.
   */
  export function PbReadString(bigint userMessage, string fieldName, number index): string;

  /**
   * Reads a color value from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The color value read, or an empty value if invalid.
   */
  export function PbReadColor(bigint userMessage, string fieldName, number index): number;

  /**
   * Reads a 2D vector from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The 2D vector value read, or an empty value if invalid.
   */
  export function PbReadVector2(bigint userMessage, string fieldName, number index): Vector2;

  /**
   * Reads a 3D vector from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The 3D vector value read, or an empty value if invalid.
   */
  export function PbReadVector3(bigint userMessage, string fieldName, number index): Vector3;

  /**
   * Reads a QAngle (rotation vector) from a UserMessage.
   * @param userMessage - Pointer to the UserMessage object.
   * @param fieldName - Name of the field to read.
   * @param index - Index of the repeated field (use -1 for non-repeated fields).
   * @returns The QAngle value read, or an empty value if invalid.
   */
  export function PbReadQAngle(bigint userMessage, string fieldName, number index): Vector3;

  /**
   * Gets a enum value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetEnum(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a enum value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetEnum(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a 32-bit integer value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetInt32(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a 32-bit integer value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetInt32(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a 64-bit integer value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetInt64(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a 64-bit integer value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetInt64(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets an unsigned 32-bit integer value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetUInt32(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets an unsigned 32-bit integer value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetUInt32(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets an unsigned 64-bit integer value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetUInt64(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets an unsigned 64-bit integer value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetUInt64(bigint userMessage, string fieldName, bigint value): boolean;

  /**
   * Gets a bool value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetBool(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a bool value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetBool(bigint userMessage, string fieldName, boolean value): boolean;

  /**
   * Gets a float value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetFloat(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a float value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetFloat(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a double value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetDouble(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a double value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetDouble(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a string value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output string.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetString(bigint userMessage, string fieldName, string out): boolean;

  /**
   * Sets a string value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetString(bigint userMessage, string fieldName, string value): boolean;

  /**
   * Gets a color value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output string.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetColor(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a color value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetColor(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a Vector2 value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output string.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetVector2(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a Vector2 value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetVector2(bigint userMessage, string fieldName, Vector2 value): boolean;

  /**
   * Gets a Vector3 value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output string.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetVector3(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a Vector3 value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetVector3(bigint userMessage, string fieldName, Vector3 value): boolean;

  /**
   * Gets a QAngle value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param out - The output string.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetQAngle(bigint userMessage, string fieldName, bigint out): boolean;

  /**
   * Sets a QAngle value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetQAngle(bigint userMessage, string fieldName, Vector3 value): boolean;

  /**
   * Gets a repeated enum value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedEnum(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated enum value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedEnum(bigint userMessage, string fieldName, number index, number value): boolean;

  /**
   * Adds a enum value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddEnum(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a repeated int32_t value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedInt32(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated int32_t value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedInt32(bigint userMessage, string fieldName, number index, number value): boolean;

  /**
   * Adds a 32-bit integer value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddInt32(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a repeated int64_t value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedInt64(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated int64_t value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedInt64(bigint userMessage, string fieldName, number index, number value): boolean;

  /**
   * Adds a 64-bit integer value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddInt64(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a repeated uint32_t value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedUInt32(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated uint32_t value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedUInt32(bigint userMessage, string fieldName, number index, number value): boolean;

  /**
   * Adds an unsigned 32-bit integer value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddUInt32(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a repeated uint64_t value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedUInt64(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated uint64_t value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedUInt64(bigint userMessage, string fieldName, number index, bigint value): boolean;

  /**
   * Adds an unsigned 64-bit integer value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddUInt64(bigint userMessage, string fieldName, bigint value): boolean;

  /**
   * Gets a repeated bool value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedBool(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated bool value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedBool(bigint userMessage, string fieldName, number index, boolean value): boolean;

  /**
   * Adds a bool value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddBool(bigint userMessage, string fieldName, boolean value): boolean;

  /**
   * Gets a repeated float value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedFloat(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated float value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedFloat(bigint userMessage, string fieldName, number index, number value): boolean;

  /**
   * Adds a float value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddFloat(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a repeated double value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output value.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedDouble(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated double value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedDouble(bigint userMessage, string fieldName, number index, number value): boolean;

  /**
   * Adds a double value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddDouble(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a repeated string value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output string.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedString(bigint userMessage, string fieldName, number index, string out): boolean;

  /**
   * Sets a repeated string value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedString(bigint userMessage, string fieldName, number index, string value): boolean;

  /**
   * Adds a string value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddString(bigint userMessage, string fieldName, string value): boolean;

  /**
   * Gets a repeated color value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output color.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedColor(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated color value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedColor(bigint userMessage, string fieldName, number index, number value): boolean;

  /**
   * Adds a color value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddColor(bigint userMessage, string fieldName, number value): boolean;

  /**
   * Gets a repeated Vector2 value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output vector2.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedVector2(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated Vector2 value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedVector2(bigint userMessage, string fieldName, number index, Vector2 value): boolean;

  /**
   * Adds a Vector2 value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddVector2(bigint userMessage, string fieldName, Vector2 value): boolean;

  /**
   * Gets a repeated Vector3 value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output vector2.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedVector3(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated Vector3 value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedVector3(bigint userMessage, string fieldName, number index, Vector3 value): boolean;

  /**
   * Adds a Vector3 value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddVector3(bigint userMessage, string fieldName, Vector3 value): boolean;

  /**
   * Gets a repeated QAngle value from a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param out - The output vector2.
   * @returns True if the field was successfully retrieved, false otherwise.
   */
  export function PbGetRepeatedQAngle(bigint userMessage, string fieldName, number index, bigint out): boolean;

  /**
   * Sets a repeated QAngle value for a field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param index - The index of the repeated field.
   * @param value - The value to set.
   * @returns True if the field was successfully set, false otherwise.
   */
  export function PbSetRepeatedQAngle(bigint userMessage, string fieldName, number index, Vector3 value): boolean;

  /**
   * Adds a QAngle value to a repeated field in the UserMessage.
   * @param userMessage - The UserMessage instance.
   * @param fieldName - The name of the field.
   * @param value - The value to add.
   * @returns True if the value was successfully added, false otherwise.
   */
  export function PbAddQAngle(bigint userMessage, string fieldName, Vector3 value): boolean;

  /**
   * Retrieves the weapon VData for a given weapon name.
   * @param name - The name of the weapon.
   * @returns A pointer to the `CCSWeaponBaseVData` if the entity handle is valid and represents a player weapon; otherwise, nullptr.
   */
  export function GetWeaponVDataFromKey(string name): bigint;

  /**
   * Retrieves the weapon VData for a given weapon.
   * @param entityHandle - The handle of the entity from which to retrieve the weapon VData.
   * @returns A pointer to the `CCSWeaponBaseVData` if the entity handle is valid and represents a player weapon; otherwise, nullptr.
   */
  export function GetWeaponVData(number entityHandle): bigint;

  /**
   * Retrieves the weapon type of a given entity.
   * @param entityHandle - The handle of the entity (weapon).
   * @returns The type of the weapon, or WEAPONTYPE_UNKNOWN if the entity is invalid.
   */
  export function GetWeaponType(number entityHandle): number;

  /**
   * Retrieves the weapon category of a given entity.
   * @param entityHandle - The handle of the entity (weapon).
   * @returns The category of the weapon, or WEAPONCATEGORY_OTHER if the entity is invalid.
   */
  export function GetWeaponCategory(number entityHandle): number;

  /**
   * Retrieves the gear slot of a given weapon entity.
   * @param entityHandle - The handle of the entity (weapon).
   * @returns The gear slot of the weapon, or GEAR_SLOT_INVALID if the entity is invalid.
   */
  export function GetWeaponGearSlot(number entityHandle): number;

  /**
   * Retrieves the weapon definition index for a given entity handle.
   * @param entityHandle - The handle of the entity from which to retrieve the weapon def index.
   * @returns The weapon definition index as a `uint16_t`, or 0 if the entity handle is invalid.
   */
  export function GetWeaponItemDefinition(number entityHandle): number;

  /**
   * Retrieves the item definition index associated with a given item name.
   * @param itemName - The name of the item.
   * @returns The weapon definition index as a `uint16_t`, or 0 if the entity handle is invalid.
   */
  export function GetWeaponItemDefinitionByName(string itemName): number;

}
