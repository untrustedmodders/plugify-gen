#pragma once

#include <plg/plugin.hpp>
#include <plg/any.hpp>
#include <cstdint>

// Generated from ${S2SDK_PACKAGE}.pplugin

namespace ${S2SDK_PACKAGE} {

  // Enum representing various movement types for entities.
  enum class MoveType : int32_t {
    None = 0, // Never moves.
    Isometric = 1, // Previously isometric movement type.
    Walk = 2, // Player only - moving on the ground.
    Fly = 3, // No gravity, but still collides with stuff.
    Flygravity = 4, // Flies through the air and is affected by gravity.
    Vphysics = 5, // Uses VPHYSICS for simulation.
    Push = 6, // No clip to world, push and crush.
    Noclip = 7, // No gravity, no collisions, still has velocity/avelocity.
    Ladder = 8, // Used by players only when going onto a ladder.
    Observer = 9, // Observer movement, depends on player's observer mode.
    Custom = 10, // Allows the entity to describe its own physics.
  };

  // Enum representing rendering modes for materials.
  enum class RenderMode : int32_t {
    Normal = 0, // Standard rendering mode (src).
    TransColor = 1, // Composite: c*a + dest*(1-a).
    TransTexture = 2, // Composite: src*a + dest*(1-a).
    Glow = 3, // Composite: src*a + dest -- No Z buffer checks -- Fixed size in screen space.
    TransAlpha = 4, // Composite: src*srca + dest*(1-srca).
    TransAdd = 5, // Composite: src*a + dest.
    Environmental = 6, // Not drawn, used for environmental effects.
    TransAddFrameBlend = 7, // Uses a fractional frame value to blend between animation frames.
    TransAlphaAdd = 8, // Composite: src + dest*(1-a).
    WorldGlow = 9, // Same as Glow but not fixed size in screen space.
    None = 10, // No rendering.
    DevVisualizer = 11, // Developer visualizer rendering mode.
  };

  // Enum representing the possible teams in Counter-Strike.
  enum class CSTeam : int32_t {
    None = 0, // No team.
    Spectator = 1, // Spectator team.
    T = 2, // Terrorist team.
    CT = 3, // Counter-Terrorist team.
  };

  // Represents the possible types of data that can be passed as a value in input actions.
  enum class FieldType : int32_t {
    Auto = 0, // Automatically detect the type of the value.
    Float32 = 1, // A 32-bit floating-point number.
    Float64 = 2, // A 64-bit floating-point number.
    Int32 = 3, // A 32-bit signed integer.
    UInt32 = 4, // A 32-bit unsigned integer.
    Int64 = 5, // A 64-bit signed integer.
    UInt64 = 6, // A 64-bit unsigned integer.
    Boolean = 7, // A boolean value (true or false).
    Character = 8, // A single character.
    String = 9, // A managed string object.
    CString = 10, // A null-terminated C-style string.
    HScript = 11, // A script handle, typically for scripting integration.
    EHandle = 12, // An entity handle, used to reference an entity within the system.
    Resource = 13, // A resource handle, such as a file or asset reference.
    Vector3d = 14, // A 3D vector, typically representing position or direction.
    Vector2d = 15, // A 2D vector, for planar data or coordinates.
    Vector4d = 16, // A 4D vector, often used for advanced mathematical representations.
    Color32 = 17, // A 32-bit color value (RGBA).
    QAngle = 18, // A quaternion-based angle representation.
    Quaternion = 19, // A quaternion, used for rotation and orientation calculations.
  };

  // Enum representing various damage types.
  enum class DamageTypes : int32_t {
    DMG_GENERIC = 0, // Generic damage.
    DMG_CRUSH = 1, // Crush damage.
    DMG_BULLET = 2, // Bullet damage.
    DMG_SLASH = 4, // Slash damage.
    DMG_BURN = 8, // Burn damage.
    DMG_VEHICLE = 16, // Vehicle damage.
    DMG_FALL = 32, // Fall damage.
    DMG_BLAST = 64, // Blast damage.
    DMG_CLUB = 128, // Club damage.
    DMG_SHOCK = 256, // Shock damage.
    DMG_SONIC = 512, // Sonic damage.
    DMG_ENERGYBEAM = 1024, // Energy beam damage.
    DMG_DROWN = 16384, // Drowning damage.
    DMG_POISON = 32768, // Poison damage.
    DMG_RADIATION = 65536, // Radiation damage.
    DMG_DROWNRECOVER = 131072, // Recovering from drowning damage.
    DMG_ACID = 262144, // Acid damage.
    DMG_PHYSGUN = 1048576, // Physgun damage.
    DMG_DISSOLVE = 2097152, // Dissolve damage.
    DMG_BLAST_SURFACE = 4194304, // Surface blast damage.
    DMG_BUCKSHOT = 16777216, // Buckshot damage.
    DMG_LASTGENERICFLAG = 16777216, // Last generic flag damage.
    DMG_HEADSHOT = 33554432, // Headshot damage.
    DMG_DANGERZONE = 67108864, // Danger zone damage.
  };

  // Enum representing various flags for ConVars and ConCommands.
  enum class ConVarFlag : int32_t {
    None = 0, // The default, no flags at all.
    LinkedConcommand = 1, // Linked to a ConCommand.
    DevelopmentOnly = 2, // Hidden in released products. Automatically removed if ALLOW_DEVELOPMENT_CVARS is defined.
    GameDll = 4, // Defined by the game DLL.
    ClientDll = 8, // Defined by the client DLL.
    Hidden = 16, // Hidden. Doesn't appear in find or auto-complete. Like DEVELOPMENTONLY but cannot be compiled out.
    Protected = 32, // Server cvar; data is not sent since it's sensitive (e.g., passwords).
    SpOnly = 64, // This cvar cannot be changed by clients connected to a multiplayer server.
    Archive = 128, // Saved to vars.rc.
    Notify = 256, // Notifies players when changed.
    UserInfo = 512, // Changes the client's info string.
    Missing0 = 1024, // Hides the cvar from lookups.
    Unlogged = 2048, // If this is a server cvar, changes are not logged to the file or console.
    Missing1 = 4096, // Hides the cvar from lookups.
    Replicated = 8192, // Server-enforced setting on clients.
    Cheat = 16384, // Only usable in singleplayer/debug or multiplayer with sv_cheats.
    PerUser = 32768, // Causes auto-generated varnameN for splitscreen slots.
    Demo = 65536, // Records this cvar when starting a demo file.
    DontRecord = 131072, // Excluded from demo files.
    Missing2 = 262144, // Reserved for future use.
    Release = 524288, // Cvars tagged with this are available to customers.
    MenuBarItem = 1048576, // Marks the cvar as a menu bar item.
    Missing3 = 2097152, // Reserved for future use.
    NotConnected = 4194304, // Cannot be changed by a client connected to a server.
    VconsoleFuzzyMatching = 8388608, // Enables fuzzy matching for vconsole.
    ServerCanExecute = 16777216, // The server can execute this command on clients.
    ClientCanExecute = 33554432, // Allows clients to execute this command.
    ServerCannotQuery = 67108864, // The server cannot query this cvar's value.
    VconsoleSetFocus = 134217728, // Sets focus in the vconsole.
    ClientCmdCanExecute = 268435456, // IVEngineClient::ClientCmd can execute this command.
    ExecutePerTick = 536870912, // Executes the cvar every tick.
  };

  // Handles the execution of a command triggered by a caller. This function processes the command, interprets its context, and handles any provided arguments.
  using CommandCallback = int32_t (*)(int32_t, int32_t, plg::vector<plg::string>);

  // Enum representing the type of callback.
  enum class HookMode : int32_t {
    Pre = 0, // Callback will be executed before the original function
    Post = 1, // Callback will be executed after the original function
  };

  // The command execution context.
  enum class CommandCallingContext : int32_t {
    Console = 0, // The command execute from the client's console.
    Chat = 1, // The command execute from the client's chat.
  };

  enum class ConVarType : int32_t {
    Invalid = -1, // Invalid type
    Bool = 0, // Boolean type
    Int16 = 1, // 16-bit signed integer
    UInt16 = 2, // 16-bit unsigned integer
    Int32 = 3, // 32-bit signed integer
    UInt32 = 4, // 32-bit unsigned integer
    Int64 = 5, // 64-bit signed integer
    UInt64 = 6, // 64-bit unsigned integer
    Float32 = 7, // 32-bit floating point
    Float64 = 8, // 64-bit floating point (double)
    String = 9, // String type
    Color = 10, // Color type
    Vector2 = 11, // 2D vector
    Vector3 = 12, // 3D vector
    Vector4 = 13, // 4D vector
    Qangle = 14, // Quaternion angle
    Max = 15, // Maximum value (used for bounds checking)
  };

  // Handles changes to a console variable's value. This function is called whenever the value of a specific console variable is modified.
  using ChangeCallback = void (*)(uint64_t, plg::string, plg::string);

  // Handles changes to a console variable's value. This function is called whenever the value of a specific console variable is modified.
  using CvarValueCallback = void (*)(int32_t, int32_t, int32_t, plg::string, plg::string, plg::vector<plg::any>);

  // Defines a QueueTask Callback.
  using TaskCallback = void (*)(plg::vector<plg::any>);

  // This function is a callback handler for entity output events. It is triggered when a specific output event is activated, and it handles the process by passing the activator, the caller, and a delay parameter for the output.
  using HookEntityOutputCallback = int32_t (*)(int32_t, int32_t, float);

  // Enum representing the type of callback.
  enum class EventHookError : int32_t {
    Okay = 0, // Indicates that the event hook was successfully created.
    InvalidEvent = 1, // Indicates that the event name provided is invalid or does not exist.
    NotActive = 2, // Indicates that the event system is not currently active or initialized.
    InvalidCallback = 3, // Indicates that the callback function provided is invalid or not compatible with the event system.
  };

  // Handles events triggered by the game event system. This function processes the event data, determines the necessary action, and optionally prevents event broadcasting.
  using EventCallback = int32_t (*)(plg::string, void*, bool);

  // Enum representing the possible verbosity of a logger.
  enum class LoggingVerbosity : int32_t {
    Off = 0, // Turns off all spew.
    Essential = 1, // Turns on vital logs.
    Default = 2, // Turns on most messages.
    Detailed = 3, // Allows for walls of text that are usually useful.
    Max = 4, // Allows everything.
  };

  // Enum representing the possible verbosity of a logger.
  enum class LoggingSeverity : int32_t {
    Off = 0, // Turns off all spew.
    Detailed = 1, // A debug message.
    Message = 2, // An informative logging message.
    Warning = 3, // A warning, typically non-fatal.
    Assert = 4, // A message caused by an Assert**() operation.
    Error = 5, // An error, typically fatal/unrecoverable.
  };

  // Logging channel behavior flags, set on channel creation.
  enum class LoggingChannelFlags : int32_t {
    ConsoleOnly = 1, // Indicates that the spew is only relevant to interactive consoles.
    DoNotEcho = 2, // Indicates that spew should not be echoed to any output devices.
  };

  // Enum representing the possible reasons a vote creation or processing has failed.
  enum class VoteCreateFailed : int32_t {
    Generic = 0, // Generic vote failure.
    TransitioningPlayers = 1, // Vote failed due to players transitioning.
    RateExceeded = 2, // Vote failed because vote rate limit was exceeded.
    YesMustExceedNo = 3, // Vote failed because Yes votes must exceed No votes.
    QuorumFailure = 4, // Vote failed due to quorum not being met.
    IssueDisabled = 5, // Vote failed because the issue is disabled.
    MapNotFound = 6, // Vote failed because the map was not found.
    MapNameRequired = 7, // Vote failed because map name is required.
    FailedRecently = 8, // Vote failed because a similar vote failed recently.
    FailedRecentKick = 9, // Vote to kick failed recently.
    FailedRecentChangeMap = 10, // Vote to change map failed recently.
    FailedRecentSwapTeams = 11, // Vote to swap teams failed recently.
    FailedRecentScrambleTeams = 12, // Vote to scramble teams failed recently.
    FailedRecentRestart = 13, // Vote to restart failed recently.
    TeamCantCall = 14, // Team is not allowed to call vote.
    WaitingForPlayers = 15, // Vote failed because game is waiting for players.
    PlayerNotFound = 16, // Target player was not found.
    CannotKickAdmin = 17, // Cannot kick an admin.
    ScrambleInProgress = 18, // Scramble is currently in progress.
    SwapInProgress = 19, // Swap is currently in progress.
    Spectator = 20, // Spectators are not allowed to vote.
    Disabled = 21, // Voting is disabled.
    NextLevelSet = 22, // Next level is already set.
    Rematch = 23, // Rematch vote failed.
    TooEarlySurrender = 24, // Vote to surrender failed due to being too early.
    Continue = 25, // Vote to continue failed.
    MatchPaused = 26, // Vote failed because match is already paused.
    MatchNotPaused = 27, // Vote failed because match is not paused.
    NotInWarmup = 28, // Vote failed because game is not in warmup.
    Not10Players = 29, // Vote failed because there are not 10 players.
    TimeoutActive = 30, // Vote failed due to an active timeout.
    TimeoutInactive = 31, // Vote failed because timeout is inactive.
    TimeoutExhausted = 32, // Vote failed because timeout has been exhausted.
    CantRoundEnd = 33, // Vote failed because the round can't end now.
    Max = 34, // Sentinel value. Not a real failure reason.
  };

  // Handles the final result of a Yes/No vote. This function is called when a vote concludes, and is responsible for determining whether the vote passed based on the number of 'yes' and 'no' votes. Also receives context about the clients who participated in the vote.
  using YesNoVoteResult = bool (*)(int32_t, int32_t, int32_t, int32_t, plg::vector<int32_t>, plg::vector<int32_t>);

  using YesNoVoteHandler = void (*)(int32_t, int32_t, int32_t);

  // Enum representing the possible types of a vote.
  enum class VoteEndReason : int32_t {
    AllVotes = 0, // All possible votes were cast.
    TimeUp = 1, // Time ran out.
    Cancelled = 2, // The vote got cancelled.
  };

  // This function is invoked when a timer event occurs. It handles the timer-related logic and performs necessary actions based on the event.
  using TimerCallback = void (*)(uint32_t, plg::vector<plg::any>);

  // Enum representing the possible flags of a timer.
  enum class TimerFlag : int32_t {
    Default = 0, // Timer with no unique properties.
    Repeat = 1, // Timer will repeat until stopped.
    NoMapChange = 2, // Timer will not carry over mapchanges.
  };

  // Called on client connection. If you return true, the client will be allowed in the server. If you return false (or return nothing), the client will be rejected. If the client is rejected by this forward or any other, OnClientDisconnect will not be called.<br>Note: Do not write to rejectmsg if you plan on returning true. If multiple plugins write to the string buffer, it is not defined which plugin's string will be shown to the client, but it is guaranteed one of them will.
  using OnClientConnectCallback = bool (*)(int32_t, plg::string, plg::string);

  // Called on client connection.
  using OnClientConnect_PostCallback = void (*)(int32_t);

  // Called once a client successfully connects. This callback is paired with OnClientDisconnect.
  using OnClientConnectedCallback = void (*)(int32_t);

  // Called when a client is entering the game.
  using OnClientPutInServerCallback = void (*)(int32_t);

  // Called when a client is disconnecting from the server.
  using OnClientDisconnectCallback = void (*)(int32_t);

  // Called when a client is disconnected from the server.
  using OnClientDisconnect_PostCallback = void (*)(int32_t, int32_t);

  // Called when a client is activated by the game.
  using OnClientActiveCallback = void (*)(int32_t, bool);

  // Called when a client is fully connected to the game.
  using OnClientFullyConnectCallback = void (*)(int32_t);

  // Called whenever the client's settings are changed.
  using OnClientSettingsChangedCallback = void (*)(int32_t);

  // Called when a client is fully connected to the game.
  using OnClientAuthenticatedCallback = void (*)(int32_t, uint64_t);

  // Called right before a round terminates.
  using OnRoundTerminatedCallback = void (*)(float, int32_t);

  // Called when an entity is created.
  using OnEntityCreatedCallback = void (*)(int32_t);

  // Called when when an entity is destroyed.
  using OnEntityDeletedCallback = void (*)(int32_t);

  // When an entity is reparented to another entity.
  using OnEntityParentChangedCallback = void (*)(int32_t, int32_t);

  // Called on every server startup.
  using OnServerStartupCallback = void (*)();

  // Called on every server activate.
  using OnServerActivateCallback = void (*)();

  // Called on every server spawn.
  using OnServerSpawnCallback = void (*)();

  // Called on every server started only once.
  using OnServerStartedCallback = void (*)();

  // Called on every map start.
  using OnMapStartCallback = void (*)();

  // Called on every map end.
  using OnMapEndCallback = void (*)();

  // Called before every server frame. Note that you should avoid doing expensive computations or declaring large local arrays.
  using OnGameFrameCallback = void (*)(bool, bool, bool);

  // Called when the server is not in game.
  using OnUpdateWhenNotInGameCallback = void (*)(float);

  // Called before every server frame, before entities are updated.
  using OnPreWorldUpdateCallback = void (*)(bool);

  // Enum representing the possible reasons for a round ending in Counter-Strike.
  enum class CSRoundEndReason : int32_t {
    TargetBombed = 1, // Target successfully bombed.
    VIPEscaped = 2, // The VIP has escaped (not present in CS:GO).
    VIPKilled = 3, // VIP has been assassinated (not present in CS:GO).
    TerroristsEscaped = 4, // The terrorists have escaped.
    CTStoppedEscape = 5, // The CTs have prevented most of the terrorists from escaping.
    TerroristsStopped = 6, // Escaping terrorists have all been neutralized.
    BombDefused = 7, // The bomb has been defused.
    CTWin = 8, // Counter-Terrorists win.
    TerroristWin = 9, // Terrorists win.
    Draw = 10, // Round draw.
    HostagesRescued = 11, // All hostages have been rescued.
    TargetSaved = 12, // Target has been saved.
    HostagesNotRescued = 13, // Hostages have not been rescued.
    TerroristsNotEscaped = 14, // Terrorists have not escaped.
    VIPNotEscaped = 15, // VIP has not escaped (not present in CS:GO).
    GameStart = 16, // Game commencing.
    TerroristsSurrender = 17, // Terrorists surrender.
    CTSurrender = 18, // CTs surrender.
    TerroristsPlanted = 19, // Terrorists planted the bomb.
    CTsReachedHostage = 20, // CTs reached the hostage.
    SurvivalWin = 21, // Survival mode win.
    SurvivalDraw = 22, // Survival mode draw.
  };

  // Callback function for user messages.
  using UserMessageCallback = int32_t (*)(void*);

  // Enum representing different weapon types.
  enum class CSWeaponType : int32_t {
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
  };

  // Enum representing different weapon categories.
  enum class CSWeaponCategory : int32_t {
    Other = 0,
    Melee = 1,
    Secondary = 2,
    SMG = 3,
    Rifle = 4,
    Heavy = 5,
    Count = 6
  };

  // Enum representing different gear slots.
  enum class GearSlot : int32_t {
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
  };

  // Enum representing different weapon definition indices.
  enum class WeaponDefIndex : int32_t {
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
  };


  /**
   * @brief Creates a new KeyValues instance
   * @function Kv1Create
   * @param setName (string): The name to assign to this KeyValues instance
   * @return ptr64: Pointer to the newly created KeyValues object
   */
  inline void* Kv1Create(plg::string setName) {
    using Kv1CreateFn = void* (*)(plg::string);
    static Kv1CreateFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1Create", reinterpret_cast<void**>(&__func));
    return __func(setName);
  }

  /**
   * @brief Destroys a KeyValues instance
   * @function Kv1Destroy
   * @param kv (ptr64): Pointer to the KeyValues object to destroy
   */
  inline void Kv1Destroy(void* kv) {
    using Kv1DestroyFn = void (*)(void*);
    static Kv1DestroyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1Destroy", reinterpret_cast<void**>(&__func));
    __func(kv);
  }

  /**
   * @brief Gets the section name of a KeyValues instance
   * @function Kv1GetName
   * @param kv (ptr64): Pointer to the KeyValues object
   * @return string: The name of the KeyValues section
   */
  inline plg::string Kv1GetName(void* kv) {
    using Kv1GetNameFn = plg::string (*)(void*);
    static Kv1GetNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1GetName", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Sets the section name of a KeyValues instance
   * @function Kv1SetName
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param name (string): The new name to assign to this KeyValues section
   */
  inline void Kv1SetName(void* kv, plg::string name) {
    using Kv1SetNameFn = void (*)(void*, plg::string);
    static Kv1SetNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1SetName", reinterpret_cast<void**>(&__func));
    __func(kv, name);
  }

  /**
   * @brief Finds a key by name
   * @function Kv1FindKey
   * @param kv (ptr64): Pointer to the KeyValues object to search in
   * @param keyName (string): The name of the key to find
   * @return ptr64: Pointer to the found KeyValues subkey, or NULL if not found
   */
  inline void* Kv1FindKey(void* kv, plg::string keyName) {
    using Kv1FindKeyFn = void* (*)(void*, plg::string);
    static Kv1FindKeyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1FindKey", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName);
  }

  /**
   * @brief Finds a key by name or creates it if it doesn't exist
   * @function Kv1FindKeyOrCreate
   * @param kv (ptr64): Pointer to the KeyValues object to search in
   * @param keyName (string): The name of the key to find or create
   * @return ptr64: Pointer to the found or newly created KeyValues subkey (never NULL)
   */
  inline void* Kv1FindKeyOrCreate(void* kv, plg::string keyName) {
    using Kv1FindKeyOrCreateFn = void* (*)(void*, plg::string);
    static Kv1FindKeyOrCreateFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1FindKeyOrCreate", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName);
  }

  /**
   * @brief Creates a new subkey with the specified name
   * @function Kv1CreateKey
   * @param kv (ptr64): Pointer to the parent KeyValues object
   * @param keyName (string): The name for the new key
   * @return ptr64: Pointer to the newly created KeyValues subkey
   */
  inline void* Kv1CreateKey(void* kv, plg::string keyName) {
    using Kv1CreateKeyFn = void* (*)(void*, plg::string);
    static Kv1CreateKeyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1CreateKey", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName);
  }

  /**
   * @brief Creates a new subkey with an autogenerated name
   * @function Kv1CreateNewKey
   * @param kv (ptr64): Pointer to the parent KeyValues object
   * @return ptr64: Pointer to the newly created KeyValues subkey
   */
  inline void* Kv1CreateNewKey(void* kv) {
    using Kv1CreateNewKeyFn = void* (*)(void*);
    static Kv1CreateNewKeyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1CreateNewKey", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Adds a subkey to this KeyValues instance
   * @function Kv1AddSubKey
   * @param kv (ptr64): Pointer to the parent KeyValues object
   * @param subKey (ptr64): Pointer to the KeyValues object to add as a child
   */
  inline void Kv1AddSubKey(void* kv, void* subKey) {
    using Kv1AddSubKeyFn = void (*)(void*, void*);
    static Kv1AddSubKeyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1AddSubKey", reinterpret_cast<void**>(&__func));
    __func(kv, subKey);
  }

  /**
   * @brief Gets the first subkey in the list
   * @function Kv1GetFirstSubKey
   * @param kv (ptr64): Pointer to the parent KeyValues object
   * @return ptr64: Pointer to the first subkey, or NULL if there are no children
   */
  inline void* Kv1GetFirstSubKey(void* kv) {
    using Kv1GetFirstSubKeyFn = void* (*)(void*);
    static Kv1GetFirstSubKeyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1GetFirstSubKey", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Gets the next sibling key in the list
   * @function Kv1GetNextKey
   * @param kv (ptr64): Pointer to the current KeyValues object
   * @return ptr64: Pointer to the next sibling key, or NULL if this is the last sibling
   */
  inline void* Kv1GetNextKey(void* kv) {
    using Kv1GetNextKeyFn = void* (*)(void*);
    static Kv1GetNextKeyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1GetNextKey", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Gets a color value from a key
   * @function Kv1GetColor
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to retrieve the color from
   * @param defaultValue (int32): The default color value to return if the key is not found
   * @return int32: The color value as a 32-bit integer (RGBA)
   */
  inline int32_t Kv1GetColor(void* kv, plg::string keyName, int32_t defaultValue) {
    using Kv1GetColorFn = int32_t (*)(void*, plg::string, int32_t);
    static Kv1GetColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1GetColor", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName, defaultValue);
  }

  /**
   * @brief Sets a color value for a key
   * @function Kv1SetColor
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to set the color for
   * @param value (int32): The color value as a 32-bit integer (RGBA)
   */
  inline void Kv1SetColor(void* kv, plg::string keyName, int32_t value) {
    using Kv1SetColorFn = void (*)(void*, plg::string, int32_t);
    static Kv1SetColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1SetColor", reinterpret_cast<void**>(&__func));
    __func(kv, keyName, value);
  }

  /**
   * @brief Gets an integer value from a key
   * @function Kv1GetInt
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to retrieve the integer from
   * @param defaultValue (int32): The default value to return if the key is not found
   * @return int32: The integer value associated with the key, or defaultValue if not found
   */
  inline int32_t Kv1GetInt(void* kv, plg::string keyName, int32_t defaultValue) {
    using Kv1GetIntFn = int32_t (*)(void*, plg::string, int32_t);
    static Kv1GetIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1GetInt", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName, defaultValue);
  }

  /**
   * @brief Sets an integer value for a key
   * @function Kv1SetInt
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to set the integer for
   * @param value (int32): The integer value to set
   */
  inline void Kv1SetInt(void* kv, plg::string keyName, int32_t value) {
    using Kv1SetIntFn = void (*)(void*, plg::string, int32_t);
    static Kv1SetIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1SetInt", reinterpret_cast<void**>(&__func));
    __func(kv, keyName, value);
  }

  /**
   * @brief Gets a float value from a key
   * @function Kv1GetFloat
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to retrieve the float from
   * @param defaultValue (float): The default value to return if the key is not found
   * @return float: The float value associated with the key, or defaultValue if not found
   */
  inline float Kv1GetFloat(void* kv, plg::string keyName, float defaultValue) {
    using Kv1GetFloatFn = float (*)(void*, plg::string, float);
    static Kv1GetFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1GetFloat", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName, defaultValue);
  }

  /**
   * @brief Sets a float value for a key
   * @function Kv1SetFloat
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to set the float for
   * @param value (float): The float value to set
   */
  inline void Kv1SetFloat(void* kv, plg::string keyName, float value) {
    using Kv1SetFloatFn = void (*)(void*, plg::string, float);
    static Kv1SetFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1SetFloat", reinterpret_cast<void**>(&__func));
    __func(kv, keyName, value);
  }

  /**
   * @brief Gets a string value from a key
   * @function Kv1GetString
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to retrieve the string from
   * @param defaultValue (string): The default string to return if the key is not found
   * @return string: The string value associated with the key, or defaultValue if not found
   */
  inline plg::string Kv1GetString(void* kv, plg::string keyName, plg::string defaultValue) {
    using Kv1GetStringFn = plg::string (*)(void*, plg::string, plg::string);
    static Kv1GetStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1GetString", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName, defaultValue);
  }

  /**
   * @brief Sets a string value for a key
   * @function Kv1SetString
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to set the string for
   * @param value (string): The string value to set
   */
  inline void Kv1SetString(void* kv, plg::string keyName, plg::string value) {
    using Kv1SetStringFn = void (*)(void*, plg::string, plg::string);
    static Kv1SetStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1SetString", reinterpret_cast<void**>(&__func));
    __func(kv, keyName, value);
  }

  /**
   * @brief Gets a pointer value from a key
   * @function Kv1GetPtr
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to retrieve the pointer from
   * @param defaultValue (ptr64): The default pointer to return if the key is not found
   * @return ptr64: The pointer value associated with the key, or defaultValue if not found
   */
  inline void* Kv1GetPtr(void* kv, plg::string keyName, void* defaultValue) {
    using Kv1GetPtrFn = void* (*)(void*, plg::string, void*);
    static Kv1GetPtrFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1GetPtr", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName, defaultValue);
  }

  /**
   * @brief Sets a pointer value for a key
   * @function Kv1SetPtr
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to set the pointer for
   * @param value (ptr64): The pointer value to set
   */
  inline void Kv1SetPtr(void* kv, plg::string keyName, void* value) {
    using Kv1SetPtrFn = void (*)(void*, plg::string, void*);
    static Kv1SetPtrFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1SetPtr", reinterpret_cast<void**>(&__func));
    __func(kv, keyName, value);
  }

  /**
   * @brief Gets a boolean value from a key
   * @function Kv1GetBool
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to retrieve the boolean from
   * @param defaultValue (bool): The default value to return if the key is not found
   * @return bool: The boolean value associated with the key, or defaultValue if not found
   */
  inline bool Kv1GetBool(void* kv, plg::string keyName, bool defaultValue) {
    using Kv1GetBoolFn = bool (*)(void*, plg::string, bool);
    static Kv1GetBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1GetBool", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName, defaultValue);
  }

  /**
   * @brief Sets a boolean value for a key
   * @function Kv1SetBool
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to set the boolean for
   * @param value (bool): The boolean value to set
   */
  inline void Kv1SetBool(void* kv, plg::string keyName, bool value) {
    using Kv1SetBoolFn = void (*)(void*, plg::string, bool);
    static Kv1SetBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1SetBool", reinterpret_cast<void**>(&__func));
    __func(kv, keyName, value);
  }

  /**
   * @brief Makes a deep copy of a KeyValues tree
   * @function Kv1MakeCopy
   * @param kv (ptr64): Pointer to the KeyValues object to copy
   * @return ptr64: Pointer to the newly allocated copy of the KeyValues tree
   */
  inline void* Kv1MakeCopy(void* kv) {
    using Kv1MakeCopyFn = void* (*)(void*);
    static Kv1MakeCopyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1MakeCopy", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Clears all subkeys and the current value
   * @function Kv1Clear
   * @param kv (ptr64): Pointer to the KeyValues object to clear
   */
  inline void Kv1Clear(void* kv) {
    using Kv1ClearFn = void (*)(void*);
    static Kv1ClearFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1Clear", reinterpret_cast<void**>(&__func));
    __func(kv);
  }

  /**
   * @brief Checks if a key exists and has no value or subkeys
   * @function Kv1IsEmpty
   * @param kv (ptr64): Pointer to the KeyValues object
   * @param keyName (string): The name of the key to check
   * @return bool: true if the key exists and is empty, false otherwise
   */
  inline bool Kv1IsEmpty(void* kv, plg::string keyName) {
    using Kv1IsEmptyFn = bool (*)(void*, plg::string);
    static Kv1IsEmptyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv1IsEmpty", reinterpret_cast<void**>(&__func));
    return __func(kv, keyName);
  }

  /**
   * @brief Creates a new KeyValues3 object with specified type and subtype
   * @function Kv3Create
   * @param type (int32): The KV3 type enumeration value
   * @param subtype (int32): The KV3 subtype enumeration value
   * @return ptr64: Pointer to the newly created KeyValues3 object
   */
  inline void* Kv3Create(int32_t type, int32_t subtype) {
    using Kv3CreateFn = void* (*)(int32_t, int32_t);
    static Kv3CreateFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3Create", reinterpret_cast<void**>(&__func));
    return __func(type, subtype);
  }

  /**
   * @brief Creates a new KeyValues3 object with cluster element, type, and subtype
   * @function Kv3CreateWithCluster
   * @param cluster_elem (int32): The cluster element index
   * @param type (int32): The KV3 type enumeration value
   * @param subtype (int32): The KV3 subtype enumeration value
   * @return ptr64: Pointer to the newly created KeyValues3 object
   */
  inline void* Kv3CreateWithCluster(int32_t cluster_elem, int32_t type, int32_t subtype) {
    using Kv3CreateWithClusterFn = void* (*)(int32_t, int32_t, int32_t);
    static Kv3CreateWithClusterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3CreateWithCluster", reinterpret_cast<void**>(&__func));
    return __func(cluster_elem, type, subtype);
  }

  /**
   * @brief Creates a copy of an existing KeyValues3 object
   * @function Kv3CreateCopy
   * @param other (ptr64): Pointer to the KeyValues3 object to copy
   * @return ptr64: Pointer to the newly created copy, or nullptr if other is null
   */
  inline void* Kv3CreateCopy(void* other) {
    using Kv3CreateCopyFn = void* (*)(void*);
    static Kv3CreateCopyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3CreateCopy", reinterpret_cast<void**>(&__func));
    return __func(other);
  }

  /**
   * @brief Destroys a KeyValues3 object and frees its memory
   * @function Kv3Destroy
   * @param kv (ptr64): Pointer to the KeyValues3 object to destroy
   */
  inline void Kv3Destroy(void* kv) {
    using Kv3DestroyFn = void (*)(void*);
    static Kv3DestroyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3Destroy", reinterpret_cast<void**>(&__func));
    __func(kv);
  }

  /**
   * @brief Copies data from another KeyValues3 object
   * @function Kv3CopyFrom
   * @param kv (ptr64): Pointer to the destination KeyValues3 object
   * @param other (ptr64): Pointer to the source KeyValues3 object
   */
  inline void Kv3CopyFrom(void* kv, void* other) {
    using Kv3CopyFromFn = void (*)(void*, void*);
    static Kv3CopyFromFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3CopyFrom", reinterpret_cast<void**>(&__func));
    __func(kv, other);
  }

  /**
   * @brief Overlays keys from another KeyValues3 object
   * @function Kv3OverlayKeysFrom
   * @param kv (ptr64): Pointer to the destination KeyValues3 object
   * @param other (ptr64): Pointer to the source KeyValues3 object
   * @param depth (bool): Whether to perform a deep overlay
   */
  inline void Kv3OverlayKeysFrom(void* kv, void* other, bool depth) {
    using Kv3OverlayKeysFromFn = void (*)(void*, void*, bool);
    static Kv3OverlayKeysFromFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3OverlayKeysFrom", reinterpret_cast<void**>(&__func));
    __func(kv, other, depth);
  }

  /**
   * @brief Gets the context associated with a KeyValues3 object
   * @function Kv3GetContext
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return ptr64: Pointer to the CKeyValues3Context, or nullptr if kv is null
   */
  inline void* Kv3GetContext(void* kv) {
    using Kv3GetContextFn = void* (*)(void*);
    static Kv3GetContextFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetContext", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Gets the metadata associated with a KeyValues3 object
   * @function Kv3GetMetaData
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param ppCtx (ptr64): Pointer to store the context pointer
   * @return ptr64: Pointer to the KV3MetaData_t structure, or nullptr if kv is null
   */
  inline void* Kv3GetMetaData(void* kv, void* ppCtx) {
    using Kv3GetMetaDataFn = void* (*)(void*, void*);
    static Kv3GetMetaDataFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMetaData", reinterpret_cast<void**>(&__func));
    return __func(kv, ppCtx);
  }

  /**
   * @brief Checks if a specific flag is set
   * @function Kv3HasFlag
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param flag (uint8): The flag to check
   * @return bool: true if the flag is set, false otherwise
   */
  inline bool Kv3HasFlag(void* kv, uint8_t flag) {
    using Kv3HasFlagFn = bool (*)(void*, uint8_t);
    static Kv3HasFlagFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3HasFlag", reinterpret_cast<void**>(&__func));
    return __func(kv, flag);
  }

  /**
   * @brief Checks if any flags are set
   * @function Kv3HasAnyFlags
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return bool: true if any flags are set, false otherwise
   */
  inline bool Kv3HasAnyFlags(void* kv) {
    using Kv3HasAnyFlagsFn = bool (*)(void*);
    static Kv3HasAnyFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3HasAnyFlags", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Gets all flags as a bitmask
   * @function Kv3GetAllFlags
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return uint8: Bitmask of all flags, or 0 if kv is null
   */
  inline uint8_t Kv3GetAllFlags(void* kv) {
    using Kv3GetAllFlagsFn = uint8_t (*)(void*);
    static Kv3GetAllFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetAllFlags", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Sets all flags from a bitmask
   * @function Kv3SetAllFlags
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param flags (uint8): Bitmask of flags to set
   */
  inline void Kv3SetAllFlags(void* kv, uint8_t flags) {
    using Kv3SetAllFlagsFn = void (*)(void*, uint8_t);
    static Kv3SetAllFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetAllFlags", reinterpret_cast<void**>(&__func));
    __func(kv, flags);
  }

  /**
   * @brief Sets or clears a specific flag
   * @function Kv3SetFlag
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param flag (uint8): The flag to modify
   * @param state (bool): true to set the flag, false to clear it
   */
  inline void Kv3SetFlag(void* kv, uint8_t flag, bool state) {
    using Kv3SetFlagFn = void (*)(void*, uint8_t, bool);
    static Kv3SetFlagFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetFlag", reinterpret_cast<void**>(&__func));
    __func(kv, flag, state);
  }

  /**
   * @brief Gets the basic type of the KeyValues3 object
   * @function Kv3GetType
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return uint8: The type enumeration value, or 0 if kv is null
   */
  inline uint8_t Kv3GetType(void* kv) {
    using Kv3GetTypeFn = uint8_t (*)(void*);
    static Kv3GetTypeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetType", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Gets the extended type of the KeyValues3 object
   * @function Kv3GetTypeEx
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return uint8: The extended type enumeration value, or 0 if kv is null
   */
  inline uint8_t Kv3GetTypeEx(void* kv) {
    using Kv3GetTypeExFn = uint8_t (*)(void*);
    static Kv3GetTypeExFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetTypeEx", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Gets the subtype of the KeyValues3 object
   * @function Kv3GetSubType
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return uint8: The subtype enumeration value, or 0 if kv is null
   */
  inline uint8_t Kv3GetSubType(void* kv) {
    using Kv3GetSubTypeFn = uint8_t (*)(void*);
    static Kv3GetSubTypeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetSubType", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Checks if the object has invalid member names
   * @function Kv3HasInvalidMemberNames
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return bool: true if invalid member names exist, false otherwise
   */
  inline bool Kv3HasInvalidMemberNames(void* kv) {
    using Kv3HasInvalidMemberNamesFn = bool (*)(void*);
    static Kv3HasInvalidMemberNamesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3HasInvalidMemberNames", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Sets the invalid member names flag
   * @function Kv3SetHasInvalidMemberNames
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param bValue (bool): true to mark as having invalid member names, false otherwise
   */
  inline void Kv3SetHasInvalidMemberNames(void* kv, bool bValue) {
    using Kv3SetHasInvalidMemberNamesFn = void (*)(void*, bool);
    static Kv3SetHasInvalidMemberNamesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetHasInvalidMemberNames", reinterpret_cast<void**>(&__func));
    __func(kv, bValue);
  }

  /**
   * @brief Gets the type as a string representation
   * @function Kv3GetTypeAsString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return string: String representation of the type, or empty string if kv is null
   */
  inline plg::string Kv3GetTypeAsString(void* kv) {
    using Kv3GetTypeAsStringFn = plg::string (*)(void*);
    static Kv3GetTypeAsStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetTypeAsString", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Gets the subtype as a string representation
   * @function Kv3GetSubTypeAsString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return string: String representation of the subtype, or empty string if kv is null
   */
  inline plg::string Kv3GetSubTypeAsString(void* kv) {
    using Kv3GetSubTypeAsStringFn = plg::string (*)(void*);
    static Kv3GetSubTypeAsStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetSubTypeAsString", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Converts the KeyValues3 object to a string representation
   * @function Kv3ToString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param flags (uint32): Formatting flags for the string conversion
   * @return string: String representation of the object, or empty string if kv is null
   */
  inline plg::string Kv3ToString(void* kv, uint32_t flags) {
    using Kv3ToStringFn = plg::string (*)(void*, uint32_t);
    static Kv3ToStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3ToString", reinterpret_cast<void**>(&__func));
    return __func(kv, flags);
  }

  /**
   * @brief Checks if the KeyValues3 object is null
   * @function Kv3IsNull
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return bool: true if the object is null or the pointer is null, false otherwise
   */
  inline bool Kv3IsNull(void* kv) {
    using Kv3IsNullFn = bool (*)(void*);
    static Kv3IsNullFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3IsNull", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Sets the KeyValues3 object to null
   * @function Kv3SetToNull
   * @param kv (ptr64): Pointer to the KeyValues3 object
   */
  inline void Kv3SetToNull(void* kv) {
    using Kv3SetToNullFn = void (*)(void*);
    static Kv3SetToNullFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetToNull", reinterpret_cast<void**>(&__func));
    __func(kv);
  }

  /**
   * @brief Checks if the KeyValues3 object is an array
   * @function Kv3IsArray
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return bool: true if the object is an array, false otherwise
   */
  inline bool Kv3IsArray(void* kv) {
    using Kv3IsArrayFn = bool (*)(void*);
    static Kv3IsArrayFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3IsArray", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Checks if the KeyValues3 object is a KV3 array
   * @function Kv3IsKV3Array
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return bool: true if the object is a KV3 array, false otherwise
   */
  inline bool Kv3IsKV3Array(void* kv) {
    using Kv3IsKV3ArrayFn = bool (*)(void*);
    static Kv3IsKV3ArrayFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3IsKV3Array", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Checks if the KeyValues3 object is a table
   * @function Kv3IsTable
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return bool: true if the object is a table, false otherwise
   */
  inline bool Kv3IsTable(void* kv) {
    using Kv3IsTableFn = bool (*)(void*);
    static Kv3IsTableFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3IsTable", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Checks if the KeyValues3 object is a string
   * @function Kv3IsString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return bool: true if the object is a string, false otherwise
   */
  inline bool Kv3IsString(void* kv) {
    using Kv3IsStringFn = bool (*)(void*);
    static Kv3IsStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3IsString", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Gets the boolean value from the KeyValues3 object
   * @function Kv3GetBool
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (bool): Default value to return if kv is null or conversion fails
   * @return bool: Boolean value or defaultValue
   */
  inline bool Kv3GetBool(void* kv, bool defaultValue) {
    using Kv3GetBoolFn = bool (*)(void*, bool);
    static Kv3GetBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetBool", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the char value from the KeyValues3 object
   * @function Kv3GetChar
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (char8): Default value to return if kv is null or conversion fails
   * @return char8: Char value or defaultValue
   */
  inline char Kv3GetChar(void* kv, char defaultValue) {
    using Kv3GetCharFn = char (*)(void*, char);
    static Kv3GetCharFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetChar", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the 32-bit Unicode character value from the KeyValues3 object
   * @function Kv3GetUChar32
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (uint32): Default value to return if kv is null or conversion fails
   * @return uint32: 32-bit Unicode character value or defaultValue
   */
  inline uint32_t Kv3GetUChar32(void* kv, uint32_t defaultValue) {
    using Kv3GetUChar32Fn = uint32_t (*)(void*, uint32_t);
    static Kv3GetUChar32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetUChar32", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the signed 8-bit integer value from the KeyValues3 object
   * @function Kv3GetInt8
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (int8): Default value to return if kv is null or conversion fails
   * @return int8: int8_t value or defaultValue
   */
  inline int8_t Kv3GetInt8(void* kv, int8_t defaultValue) {
    using Kv3GetInt8Fn = int8_t (*)(void*, int8_t);
    static Kv3GetInt8Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetInt8", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the unsigned 8-bit integer value from the KeyValues3 object
   * @function Kv3GetUInt8
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (uint8): Default value to return if kv is null or conversion fails
   * @return uint8: uint8_t value or defaultValue
   */
  inline uint8_t Kv3GetUInt8(void* kv, uint8_t defaultValue) {
    using Kv3GetUInt8Fn = uint8_t (*)(void*, uint8_t);
    static Kv3GetUInt8Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetUInt8", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the signed 16-bit integer value from the KeyValues3 object
   * @function Kv3GetShort
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (int16): Default value to return if kv is null or conversion fails
   * @return int16: int16_t value or defaultValue
   */
  inline int16_t Kv3GetShort(void* kv, int16_t defaultValue) {
    using Kv3GetShortFn = int16_t (*)(void*, int16_t);
    static Kv3GetShortFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetShort", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the unsigned 16-bit integer value from the KeyValues3 object
   * @function Kv3GetUShort
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (uint16): Default value to return if kv is null or conversion fails
   * @return uint16: uint16_t value or defaultValue
   */
  inline uint16_t Kv3GetUShort(void* kv, uint16_t defaultValue) {
    using Kv3GetUShortFn = uint16_t (*)(void*, uint16_t);
    static Kv3GetUShortFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetUShort", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the signed 32-bit integer value from the KeyValues3 object
   * @function Kv3GetInt
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (int32): Default value to return if kv is null or conversion fails
   * @return int32: int32_t value or defaultValue
   */
  inline int32_t Kv3GetInt(void* kv, int32_t defaultValue) {
    using Kv3GetIntFn = int32_t (*)(void*, int32_t);
    static Kv3GetIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetInt", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the unsigned 32-bit integer value from the KeyValues3 object
   * @function Kv3GetUInt
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (uint32): Default value to return if kv is null or conversion fails
   * @return uint32: uint32_t value or defaultValue
   */
  inline uint32_t Kv3GetUInt(void* kv, uint32_t defaultValue) {
    using Kv3GetUIntFn = uint32_t (*)(void*, uint32_t);
    static Kv3GetUIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetUInt", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the signed 64-bit integer value from the KeyValues3 object
   * @function Kv3GetInt64
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (int64): Default value to return if kv is null or conversion fails
   * @return int64: int64_t value or defaultValue
   */
  inline int64_t Kv3GetInt64(void* kv, int64_t defaultValue) {
    using Kv3GetInt64Fn = int64_t (*)(void*, int64_t);
    static Kv3GetInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetInt64", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the unsigned 64-bit integer value from the KeyValues3 object
   * @function Kv3GetUInt64
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (uint64): Default value to return if kv is null or conversion fails
   * @return uint64: uint64_t value or defaultValue
   */
  inline uint64_t Kv3GetUInt64(void* kv, uint64_t defaultValue) {
    using Kv3GetUInt64Fn = uint64_t (*)(void*, uint64_t);
    static Kv3GetUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetUInt64", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the float value from the KeyValues3 object
   * @function Kv3GetFloat
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (float): Default value to return if kv is null or conversion fails
   * @return float: Float value or defaultValue
   */
  inline float Kv3GetFloat(void* kv, float defaultValue) {
    using Kv3GetFloatFn = float (*)(void*, float);
    static Kv3GetFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetFloat", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the double value from the KeyValues3 object
   * @function Kv3GetDouble
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (double): Default value to return if kv is null or conversion fails
   * @return double: Double value or defaultValue
   */
  inline double Kv3GetDouble(void* kv, double defaultValue) {
    using Kv3GetDoubleFn = double (*)(void*, double);
    static Kv3GetDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetDouble", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Sets the KeyValues3 object to a boolean value
   * @function Kv3SetBool
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (bool): Boolean value to set
   */
  inline void Kv3SetBool(void* kv, bool value) {
    using Kv3SetBoolFn = void (*)(void*, bool);
    static Kv3SetBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetBool", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to a char value
   * @function Kv3SetChar
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (char8): Char value to set
   */
  inline void Kv3SetChar(void* kv, char value) {
    using Kv3SetCharFn = void (*)(void*, char);
    static Kv3SetCharFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetChar", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to a 32-bit Unicode character value
   * @function Kv3SetUChar32
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (uint32): 32-bit Unicode character value to set
   */
  inline void Kv3SetUChar32(void* kv, uint32_t value) {
    using Kv3SetUChar32Fn = void (*)(void*, uint32_t);
    static Kv3SetUChar32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetUChar32", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to a signed 8-bit integer value
   * @function Kv3SetInt8
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (int8): int8_t value to set
   */
  inline void Kv3SetInt8(void* kv, int8_t value) {
    using Kv3SetInt8Fn = void (*)(void*, int8_t);
    static Kv3SetInt8Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetInt8", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to an unsigned 8-bit integer value
   * @function Kv3SetUInt8
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (uint8): uint8_t value to set
   */
  inline void Kv3SetUInt8(void* kv, uint8_t value) {
    using Kv3SetUInt8Fn = void (*)(void*, uint8_t);
    static Kv3SetUInt8Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetUInt8", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to a signed 16-bit integer value
   * @function Kv3SetShort
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (int16): int16_t value to set
   */
  inline void Kv3SetShort(void* kv, int16_t value) {
    using Kv3SetShortFn = void (*)(void*, int16_t);
    static Kv3SetShortFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetShort", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to an unsigned 16-bit integer value
   * @function Kv3SetUShort
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (uint16): uint16_t value to set
   */
  inline void Kv3SetUShort(void* kv, uint16_t value) {
    using Kv3SetUShortFn = void (*)(void*, uint16_t);
    static Kv3SetUShortFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetUShort", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to a signed 32-bit integer value
   * @function Kv3SetInt
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (int32): int32_t value to set
   */
  inline void Kv3SetInt(void* kv, int32_t value) {
    using Kv3SetIntFn = void (*)(void*, int32_t);
    static Kv3SetIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetInt", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to an unsigned 32-bit integer value
   * @function Kv3SetUInt
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (uint32): uint32_t value to set
   */
  inline void Kv3SetUInt(void* kv, uint32_t value) {
    using Kv3SetUIntFn = void (*)(void*, uint32_t);
    static Kv3SetUIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetUInt", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to a signed 64-bit integer value
   * @function Kv3SetInt64
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (int64): int64_t value to set
   */
  inline void Kv3SetInt64(void* kv, int64_t value) {
    using Kv3SetInt64Fn = void (*)(void*, int64_t);
    static Kv3SetInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetInt64", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to an unsigned 64-bit integer value
   * @function Kv3SetUInt64
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (uint64): uint64_t value to set
   */
  inline void Kv3SetUInt64(void* kv, uint64_t value) {
    using Kv3SetUInt64Fn = void (*)(void*, uint64_t);
    static Kv3SetUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetUInt64", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to a float value
   * @function Kv3SetFloat
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (float): Float value to set
   */
  inline void Kv3SetFloat(void* kv, float value) {
    using Kv3SetFloatFn = void (*)(void*, float);
    static Kv3SetFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetFloat", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Sets the KeyValues3 object to a double value
   * @function Kv3SetDouble
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param value (double): Double value to set
   */
  inline void Kv3SetDouble(void* kv, double value) {
    using Kv3SetDoubleFn = void (*)(void*, double);
    static Kv3SetDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetDouble", reinterpret_cast<void**>(&__func));
    __func(kv, value);
  }

  /**
   * @brief Gets the pointer value from the KeyValues3 object
   * @function Kv3GetPointer
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (ptr64): Default value to return if kv is null
   * @return ptr64: Pointer value as uintptr_t or defaultValue
   */
  inline void* Kv3GetPointer(void* kv, void* defaultValue) {
    using Kv3GetPointerFn = void* (*)(void*, void*);
    static Kv3GetPointerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetPointer", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Sets the KeyValues3 object to a pointer value
   * @function Kv3SetPointer
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param ptr (ptr64): Pointer value as uintptr_t to set
   */
  inline void Kv3SetPointer(void* kv, void* ptr) {
    using Kv3SetPointerFn = void (*)(void*, void*);
    static Kv3SetPointerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetPointer", reinterpret_cast<void**>(&__func));
    __func(kv, ptr);
  }

  /**
   * @brief Gets the string token value from the KeyValues3 object
   * @function Kv3GetStringToken
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (uint32): Default token value to return if kv is null
   * @return uint32: String token hash code or defaultValue
   */
  inline uint32_t Kv3GetStringToken(void* kv, uint32_t defaultValue) {
    using Kv3GetStringTokenFn = uint32_t (*)(void*, uint32_t);
    static Kv3GetStringTokenFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetStringToken", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Sets the KeyValues3 object to a string token value
   * @function Kv3SetStringToken
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param token (uint32): String token hash code to set
   */
  inline void Kv3SetStringToken(void* kv, uint32_t token) {
    using Kv3SetStringTokenFn = void (*)(void*, uint32_t);
    static Kv3SetStringTokenFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetStringToken", reinterpret_cast<void**>(&__func));
    __func(kv, token);
  }

  /**
   * @brief Gets the entity handle value from the KeyValues3 object
   * @function Kv3GetEHandle
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (int32): Default entity handle value to return if kv is null
   * @return int32: Entity handle as int32_t or defaultValue
   */
  inline int32_t Kv3GetEHandle(void* kv, int32_t defaultValue) {
    using Kv3GetEHandleFn = int32_t (*)(void*, int32_t);
    static Kv3GetEHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetEHandle", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Sets the KeyValues3 object to an entity handle value
   * @function Kv3SetEHandle
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param ehandle (int32): Entity handle value to set
   */
  inline void Kv3SetEHandle(void* kv, int32_t ehandle) {
    using Kv3SetEHandleFn = void (*)(void*, int32_t);
    static Kv3SetEHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetEHandle", reinterpret_cast<void**>(&__func));
    __func(kv, ehandle);
  }

  /**
   * @brief Gets the string value from the KeyValues3 object
   * @function Kv3GetString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (string): Default string to return if kv is null or value is empty
   * @return string: String value or defaultValue
   */
  inline plg::string Kv3GetString(void* kv, plg::string defaultValue) {
    using Kv3GetStringFn = plg::string (*)(void*, plg::string);
    static Kv3GetStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetString", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Sets the KeyValues3 object to a string value (copies the string)
   * @function Kv3SetString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param str (string): String value to set
   * @param subtype (uint8): String subtype enumeration value
   */
  inline void Kv3SetString(void* kv, plg::string str, uint8_t subtype) {
    using Kv3SetStringFn = void (*)(void*, plg::string, uint8_t);
    static Kv3SetStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetString", reinterpret_cast<void**>(&__func));
    __func(kv, str, subtype);
  }

  /**
   * @brief Sets the KeyValues3 object to an external string value (does not copy)
   * @function Kv3SetStringExternal
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param str (string): External string value to reference
   * @param subtype (uint8): String subtype enumeration value
   */
  inline void Kv3SetStringExternal(void* kv, plg::string str, uint8_t subtype) {
    using Kv3SetStringExternalFn = void (*)(void*, plg::string, uint8_t);
    static Kv3SetStringExternalFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetStringExternal", reinterpret_cast<void**>(&__func));
    __func(kv, str, subtype);
  }

  /**
   * @brief Gets the binary blob from the KeyValues3 object
   * @function Kv3GetBinaryBlob
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return uint8[]: Vector containing the binary blob data, or empty vector if kv is null
   */
  inline plg::vector<uint8_t> Kv3GetBinaryBlob(void* kv) {
    using Kv3GetBinaryBlobFn = plg::vector<uint8_t> (*)(void*);
    static Kv3GetBinaryBlobFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetBinaryBlob", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Gets the size of the binary blob in the KeyValues3 object
   * @function Kv3GetBinaryBlobSize
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return int32: Size of the binary blob in bytes, or 0 if kv is null
   */
  inline int32_t Kv3GetBinaryBlobSize(void* kv) {
    using Kv3GetBinaryBlobSizeFn = int32_t (*)(void*);
    static Kv3GetBinaryBlobSizeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetBinaryBlobSize", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Sets the KeyValues3 object to a binary blob (copies the data)
   * @function Kv3SetToBinaryBlob
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param blob (uint8[]): Vector containing the binary blob data
   */
  inline void Kv3SetToBinaryBlob(void* kv, plg::vector<uint8_t> blob) {
    using Kv3SetToBinaryBlobFn = void (*)(void*, plg::vector<uint8_t>);
    static Kv3SetToBinaryBlobFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetToBinaryBlob", reinterpret_cast<void**>(&__func));
    __func(kv, blob);
  }

  /**
   * @brief Sets the KeyValues3 object to an external binary blob (does not copy)
   * @function Kv3SetToBinaryBlobExternal
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param blob (uint8[]): Vector containing the external binary blob data
   * @param free_mem (bool): Whether to free the memory when the object is destroyed
   */
  inline void Kv3SetToBinaryBlobExternal(void* kv, plg::vector<uint8_t> blob, bool free_mem) {
    using Kv3SetToBinaryBlobExternalFn = void (*)(void*, plg::vector<uint8_t>, bool);
    static Kv3SetToBinaryBlobExternalFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetToBinaryBlobExternal", reinterpret_cast<void**>(&__func));
    __func(kv, blob, free_mem);
  }

  /**
   * @brief Gets the color value from the KeyValues3 object
   * @function Kv3GetColor
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (int32): Default color value to return if kv is null
   * @return int32: Color value as int32_t or defaultValue
   */
  inline int32_t Kv3GetColor(void* kv, int32_t defaultValue) {
    using Kv3GetColorFn = int32_t (*)(void*, int32_t);
    static Kv3GetColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetColor", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Sets the KeyValues3 object to a color value
   * @function Kv3SetColor
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param color (int32): Color value as int32_t to set
   */
  inline void Kv3SetColor(void* kv, int32_t color) {
    using Kv3SetColorFn = void (*)(void*, int32_t);
    static Kv3SetColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetColor", reinterpret_cast<void**>(&__func));
    __func(kv, color);
  }

  /**
   * @brief Gets the 3D vector value from the KeyValues3 object
   * @function Kv3GetVector
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (vec3): Default vector to return if kv is null
   * @return vec3: 3D vector or defaultValue
   */
  inline plg::vec3 Kv3GetVector(void* kv, plg::vec3 defaultValue) {
    using Kv3GetVectorFn = plg::vec3 (*)(void*, plg::vec3);
    static Kv3GetVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetVector", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the 2D vector value from the KeyValues3 object
   * @function Kv3GetVector2D
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (vec2): Default 2D vector to return if kv is null
   * @return vec2: 2D vector or defaultValue
   */
  inline plg::vec2 Kv3GetVector2D(void* kv, plg::vec2 defaultValue) {
    using Kv3GetVector2DFn = plg::vec2 (*)(void*, plg::vec2);
    static Kv3GetVector2DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetVector2D", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the 4D vector value from the KeyValues3 object
   * @function Kv3GetVector4D
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (vec4): Default 4D vector to return if kv is null
   * @return vec4: 4D vector or defaultValue
   */
  inline plg::vec4 Kv3GetVector4D(void* kv, plg::vec4 defaultValue) {
    using Kv3GetVector4DFn = plg::vec4 (*)(void*, plg::vec4);
    static Kv3GetVector4DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetVector4D", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the quaternion value from the KeyValues3 object
   * @function Kv3GetQuaternion
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (vec4): Default quaternion to return if kv is null
   * @return vec4: Quaternion as vec4 or defaultValue
   */
  inline plg::vec4 Kv3GetQuaternion(void* kv, plg::vec4 defaultValue) {
    using Kv3GetQuaternionFn = plg::vec4 (*)(void*, plg::vec4);
    static Kv3GetQuaternionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetQuaternion", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the angle (QAngle) value from the KeyValues3 object
   * @function Kv3GetQAngle
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (vec3): Default angle to return if kv is null
   * @return vec3: QAngle as vec3 or defaultValue
   */
  inline plg::vec3 Kv3GetQAngle(void* kv, plg::vec3 defaultValue) {
    using Kv3GetQAngleFn = plg::vec3 (*)(void*, plg::vec3);
    static Kv3GetQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetQAngle", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Gets the 3x4 matrix value from the KeyValues3 object
   * @function Kv3GetMatrix3x4
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param defaultValue (mat4x4): Default matrix to return if kv is null
   * @return mat4x4: 3x4 matrix as mat4x4 or defaultValue
   */
  inline plg::mat4x4 Kv3GetMatrix3x4(void* kv, plg::mat4x4 defaultValue) {
    using Kv3GetMatrix3x4Fn = plg::mat4x4 (*)(void*, plg::mat4x4);
    static Kv3GetMatrix3x4Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMatrix3x4", reinterpret_cast<void**>(&__func));
    return __func(kv, defaultValue);
  }

  /**
   * @brief Sets the KeyValues3 object to a 3D vector value
   * @function Kv3SetVector
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param vec (vec3): 3D vector to set
   */
  inline void Kv3SetVector(void* kv, plg::vec3 vec) {
    using Kv3SetVectorFn = void (*)(void*, plg::vec3);
    static Kv3SetVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetVector", reinterpret_cast<void**>(&__func));
    __func(kv, vec);
  }

  /**
   * @brief Sets the KeyValues3 object to a 2D vector value
   * @function Kv3SetVector2D
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param vec2d (vec2): 2D vector to set
   */
  inline void Kv3SetVector2D(void* kv, plg::vec2 vec2d) {
    using Kv3SetVector2DFn = void (*)(void*, plg::vec2);
    static Kv3SetVector2DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetVector2D", reinterpret_cast<void**>(&__func));
    __func(kv, vec2d);
  }

  /**
   * @brief Sets the KeyValues3 object to a 4D vector value
   * @function Kv3SetVector4D
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param vec4d (vec4): 4D vector to set
   */
  inline void Kv3SetVector4D(void* kv, plg::vec4 vec4d) {
    using Kv3SetVector4DFn = void (*)(void*, plg::vec4);
    static Kv3SetVector4DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetVector4D", reinterpret_cast<void**>(&__func));
    __func(kv, vec4d);
  }

  /**
   * @brief Sets the KeyValues3 object to a quaternion value
   * @function Kv3SetQuaternion
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param quat (vec4): Quaternion to set (as vec4)
   */
  inline void Kv3SetQuaternion(void* kv, plg::vec4 quat) {
    using Kv3SetQuaternionFn = void (*)(void*, plg::vec4);
    static Kv3SetQuaternionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetQuaternion", reinterpret_cast<void**>(&__func));
    __func(kv, quat);
  }

  /**
   * @brief Sets the KeyValues3 object to an angle (QAngle) value
   * @function Kv3SetQAngle
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param ang (vec3): QAngle to set (as vec3)
   */
  inline void Kv3SetQAngle(void* kv, plg::vec3 ang) {
    using Kv3SetQAngleFn = void (*)(void*, plg::vec3);
    static Kv3SetQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetQAngle", reinterpret_cast<void**>(&__func));
    __func(kv, ang);
  }

  /**
   * @brief Sets the KeyValues3 object to a 3x4 matrix value
   * @function Kv3SetMatrix3x4
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param matrix (mat4x4): 3x4 matrix to set (as mat4x4)
   */
  inline void Kv3SetMatrix3x4(void* kv, plg::mat4x4 matrix) {
    using Kv3SetMatrix3x4Fn = void (*)(void*, plg::mat4x4);
    static Kv3SetMatrix3x4Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMatrix3x4", reinterpret_cast<void**>(&__func));
    __func(kv, matrix);
  }

  /**
   * @brief Gets the number of elements in the array
   * @function Kv3GetArrayElementCount
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return int32: Number of array elements, or 0 if kv is null or not an array
   */
  inline int32_t Kv3GetArrayElementCount(void* kv) {
    using Kv3GetArrayElementCountFn = int32_t (*)(void*);
    static Kv3GetArrayElementCountFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetArrayElementCount", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Sets the number of elements in the array
   * @function Kv3SetArrayElementCount
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param count (int32): Number of elements to set
   * @param type (uint8): Type of array elements
   * @param subtype (uint8): Subtype of array elements
   */
  inline void Kv3SetArrayElementCount(void* kv, int32_t count, uint8_t type, uint8_t subtype) {
    using Kv3SetArrayElementCountFn = void (*)(void*, int32_t, uint8_t, uint8_t);
    static Kv3SetArrayElementCountFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetArrayElementCount", reinterpret_cast<void**>(&__func));
    __func(kv, count, type, subtype);
  }

  /**
   * @brief Sets the KeyValues3 object to an empty KV3 array
   * @function Kv3SetToEmptyKV3Array
   * @param kv (ptr64): Pointer to the KeyValues3 object
   */
  inline void Kv3SetToEmptyKV3Array(void* kv) {
    using Kv3SetToEmptyKV3ArrayFn = void (*)(void*);
    static Kv3SetToEmptyKV3ArrayFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetToEmptyKV3Array", reinterpret_cast<void**>(&__func));
    __func(kv);
  }

  /**
   * @brief Gets an array element at the specified index
   * @function Kv3GetArrayElement
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param elem (int32): Index of the element to get
   * @return ptr64: Pointer to the element KeyValues3 object, or nullptr if invalid
   */
  inline void* Kv3GetArrayElement(void* kv, int32_t elem) {
    using Kv3GetArrayElementFn = void* (*)(void*, int32_t);
    static Kv3GetArrayElementFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetArrayElement", reinterpret_cast<void**>(&__func));
    return __func(kv, elem);
  }

  /**
   * @brief Inserts a new element before the specified index
   * @function Kv3ArrayInsertElementBefore
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param elem (int32): Index before which to insert
   * @return ptr64: Pointer to the newly inserted element, or nullptr if invalid
   */
  inline void* Kv3ArrayInsertElementBefore(void* kv, int32_t elem) {
    using Kv3ArrayInsertElementBeforeFn = void* (*)(void*, int32_t);
    static Kv3ArrayInsertElementBeforeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3ArrayInsertElementBefore", reinterpret_cast<void**>(&__func));
    return __func(kv, elem);
  }

  /**
   * @brief Inserts a new element after the specified index
   * @function Kv3ArrayInsertElementAfter
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param elem (int32): Index after which to insert
   * @return ptr64: Pointer to the newly inserted element, or nullptr if invalid
   */
  inline void* Kv3ArrayInsertElementAfter(void* kv, int32_t elem) {
    using Kv3ArrayInsertElementAfterFn = void* (*)(void*, int32_t);
    static Kv3ArrayInsertElementAfterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3ArrayInsertElementAfter", reinterpret_cast<void**>(&__func));
    return __func(kv, elem);
  }

  /**
   * @brief Adds a new element to the end of the array
   * @function Kv3ArrayAddElementToTail
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return ptr64: Pointer to the newly added element, or nullptr if invalid
   */
  inline void* Kv3ArrayAddElementToTail(void* kv) {
    using Kv3ArrayAddElementToTailFn = void* (*)(void*);
    static Kv3ArrayAddElementToTailFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3ArrayAddElementToTail", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Swaps two array elements
   * @function Kv3ArraySwapItems
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param idx1 (int32): Index of the first element
   * @param idx2 (int32): Index of the second element
   */
  inline void Kv3ArraySwapItems(void* kv, int32_t idx1, int32_t idx2) {
    using Kv3ArraySwapItemsFn = void (*)(void*, int32_t, int32_t);
    static Kv3ArraySwapItemsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3ArraySwapItems", reinterpret_cast<void**>(&__func));
    __func(kv, idx1, idx2);
  }

  /**
   * @brief Removes an element from the array
   * @function Kv3ArrayRemoveElement
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param elem (int32): Index of the element to remove
   */
  inline void Kv3ArrayRemoveElement(void* kv, int32_t elem) {
    using Kv3ArrayRemoveElementFn = void (*)(void*, int32_t);
    static Kv3ArrayRemoveElementFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3ArrayRemoveElement", reinterpret_cast<void**>(&__func));
    __func(kv, elem);
  }

  /**
   * @brief Sets the KeyValues3 object to an empty table
   * @function Kv3SetToEmptyTable
   * @param kv (ptr64): Pointer to the KeyValues3 object
   */
  inline void Kv3SetToEmptyTable(void* kv) {
    using Kv3SetToEmptyTableFn = void (*)(void*);
    static Kv3SetToEmptyTableFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetToEmptyTable", reinterpret_cast<void**>(&__func));
    __func(kv);
  }

  /**
   * @brief Gets the number of members in the table
   * @function Kv3GetMemberCount
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @return int32: Number of table members, or 0 if kv is null or not a table
   */
  inline int32_t Kv3GetMemberCount(void* kv) {
    using Kv3GetMemberCountFn = int32_t (*)(void*);
    static Kv3GetMemberCountFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberCount", reinterpret_cast<void**>(&__func));
    return __func(kv);
  }

  /**
   * @brief Checks if a member with the specified name exists
   * @function Kv3HasMember
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member to check
   * @return bool: true if the member exists, false otherwise
   */
  inline bool Kv3HasMember(void* kv, plg::string name) {
    using Kv3HasMemberFn = bool (*)(void*, plg::string);
    static Kv3HasMemberFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3HasMember", reinterpret_cast<void**>(&__func));
    return __func(kv, name);
  }

  /**
   * @brief Finds a member by name
   * @function Kv3FindMember
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member to find
   * @return ptr64: Pointer to the member KeyValues3 object, or nullptr if not found
   */
  inline void* Kv3FindMember(void* kv, plg::string name) {
    using Kv3FindMemberFn = void* (*)(void*, plg::string);
    static Kv3FindMemberFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3FindMember", reinterpret_cast<void**>(&__func));
    return __func(kv, name);
  }

  /**
   * @brief Finds a member by name, or creates it if it doesn't exist
   * @function Kv3FindOrCreateMember
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member to find or create
   * @return ptr64: Pointer to the member KeyValues3 object, or nullptr if kv is null
   */
  inline void* Kv3FindOrCreateMember(void* kv, plg::string name) {
    using Kv3FindOrCreateMemberFn = void* (*)(void*, plg::string);
    static Kv3FindOrCreateMemberFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3FindOrCreateMember", reinterpret_cast<void**>(&__func));
    return __func(kv, name);
  }

  /**
   * @brief Removes a member from the table
   * @function Kv3RemoveMember
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member to remove
   * @return bool: true if the member was removed, false otherwise
   */
  inline bool Kv3RemoveMember(void* kv, plg::string name) {
    using Kv3RemoveMemberFn = bool (*)(void*, plg::string);
    static Kv3RemoveMemberFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3RemoveMember", reinterpret_cast<void**>(&__func));
    return __func(kv, name);
  }

  /**
   * @brief Gets the name of a member at the specified index
   * @function Kv3GetMemberName
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param index (int32): Index of the member
   * @return string: Name of the member, or empty string if invalid
   */
  inline plg::string Kv3GetMemberName(void* kv, int32_t index) {
    using Kv3GetMemberNameFn = plg::string (*)(void*, int32_t);
    static Kv3GetMemberNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberName", reinterpret_cast<void**>(&__func));
    return __func(kv, index);
  }

  /**
   * @brief Gets a member by index
   * @function Kv3GetMemberByIndex
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param index (int32): Index of the member to get
   * @return ptr64: Pointer to the member KeyValues3 object, or nullptr if invalid
   */
  inline void* Kv3GetMemberByIndex(void* kv, int32_t index) {
    using Kv3GetMemberByIndexFn = void* (*)(void*, int32_t);
    static Kv3GetMemberByIndexFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberByIndex", reinterpret_cast<void**>(&__func));
    return __func(kv, index);
  }

  /**
   * @brief Gets a boolean value from a table member
   * @function Kv3GetMemberBool
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (bool): Default value to return if member not found
   * @return bool: Boolean value or defaultValue
   */
  inline bool Kv3GetMemberBool(void* kv, plg::string name, bool defaultValue) {
    using Kv3GetMemberBoolFn = bool (*)(void*, plg::string, bool);
    static Kv3GetMemberBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberBool", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a char value from a table member
   * @function Kv3GetMemberChar
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (char8): Default value to return if member not found
   * @return char8: Char value or defaultValue
   */
  inline char Kv3GetMemberChar(void* kv, plg::string name, char defaultValue) {
    using Kv3GetMemberCharFn = char (*)(void*, plg::string, char);
    static Kv3GetMemberCharFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberChar", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a 32-bit Unicode character value from a table member
   * @function Kv3GetMemberUChar32
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (uint32): Default value to return if member not found
   * @return uint32: 32-bit Unicode character value or defaultValue
   */
  inline uint32_t Kv3GetMemberUChar32(void* kv, plg::string name, uint32_t defaultValue) {
    using Kv3GetMemberUChar32Fn = uint32_t (*)(void*, plg::string, uint32_t);
    static Kv3GetMemberUChar32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberUChar32", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a signed 8-bit integer value from a table member
   * @function Kv3GetMemberInt8
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (int8): Default value to return if member not found
   * @return int8: int8_t value or defaultValue
   */
  inline int8_t Kv3GetMemberInt8(void* kv, plg::string name, int8_t defaultValue) {
    using Kv3GetMemberInt8Fn = int8_t (*)(void*, plg::string, int8_t);
    static Kv3GetMemberInt8Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberInt8", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets an unsigned 8-bit integer value from a table member
   * @function Kv3GetMemberUInt8
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (uint8): Default value to return if member not found
   * @return uint8: uint8_t value or defaultValue
   */
  inline uint8_t Kv3GetMemberUInt8(void* kv, plg::string name, uint8_t defaultValue) {
    using Kv3GetMemberUInt8Fn = uint8_t (*)(void*, plg::string, uint8_t);
    static Kv3GetMemberUInt8Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberUInt8", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a signed 16-bit integer value from a table member
   * @function Kv3GetMemberShort
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (int16): Default value to return if member not found
   * @return int16: int16_t value or defaultValue
   */
  inline int16_t Kv3GetMemberShort(void* kv, plg::string name, int16_t defaultValue) {
    using Kv3GetMemberShortFn = int16_t (*)(void*, plg::string, int16_t);
    static Kv3GetMemberShortFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberShort", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets an unsigned 16-bit integer value from a table member
   * @function Kv3GetMemberUShort
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (uint16): Default value to return if member not found
   * @return uint16: uint16_t value or defaultValue
   */
  inline uint16_t Kv3GetMemberUShort(void* kv, plg::string name, uint16_t defaultValue) {
    using Kv3GetMemberUShortFn = uint16_t (*)(void*, plg::string, uint16_t);
    static Kv3GetMemberUShortFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberUShort", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a signed 32-bit integer value from a table member
   * @function Kv3GetMemberInt
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (int32): Default value to return if member not found
   * @return int32: int32_t value or defaultValue
   */
  inline int32_t Kv3GetMemberInt(void* kv, plg::string name, int32_t defaultValue) {
    using Kv3GetMemberIntFn = int32_t (*)(void*, plg::string, int32_t);
    static Kv3GetMemberIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberInt", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets an unsigned 32-bit integer value from a table member
   * @function Kv3GetMemberUInt
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (uint32): Default value to return if member not found
   * @return uint32: uint32_t value or defaultValue
   */
  inline uint32_t Kv3GetMemberUInt(void* kv, plg::string name, uint32_t defaultValue) {
    using Kv3GetMemberUIntFn = uint32_t (*)(void*, plg::string, uint32_t);
    static Kv3GetMemberUIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberUInt", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a signed 64-bit integer value from a table member
   * @function Kv3GetMemberInt64
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (int64): Default value to return if member not found
   * @return int64: int64_t value or defaultValue
   */
  inline int64_t Kv3GetMemberInt64(void* kv, plg::string name, int64_t defaultValue) {
    using Kv3GetMemberInt64Fn = int64_t (*)(void*, plg::string, int64_t);
    static Kv3GetMemberInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberInt64", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets an unsigned 64-bit integer value from a table member
   * @function Kv3GetMemberUInt64
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (uint64): Default value to return if member not found
   * @return uint64: uint64_t value or defaultValue
   */
  inline uint64_t Kv3GetMemberUInt64(void* kv, plg::string name, uint64_t defaultValue) {
    using Kv3GetMemberUInt64Fn = uint64_t (*)(void*, plg::string, uint64_t);
    static Kv3GetMemberUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberUInt64", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a float value from a table member
   * @function Kv3GetMemberFloat
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (float): Default value to return if member not found
   * @return float: Float value or defaultValue
   */
  inline float Kv3GetMemberFloat(void* kv, plg::string name, float defaultValue) {
    using Kv3GetMemberFloatFn = float (*)(void*, plg::string, float);
    static Kv3GetMemberFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberFloat", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a double value from a table member
   * @function Kv3GetMemberDouble
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (double): Default value to return if member not found
   * @return double: Double value or defaultValue
   */
  inline double Kv3GetMemberDouble(void* kv, plg::string name, double defaultValue) {
    using Kv3GetMemberDoubleFn = double (*)(void*, plg::string, double);
    static Kv3GetMemberDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberDouble", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a pointer value from a table member
   * @function Kv3GetMemberPointer
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (ptr64): Default value to return if member not found
   * @return ptr64: Pointer value as uintptr_t or defaultValue
   */
  inline void* Kv3GetMemberPointer(void* kv, plg::string name, void* defaultValue) {
    using Kv3GetMemberPointerFn = void* (*)(void*, plg::string, void*);
    static Kv3GetMemberPointerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberPointer", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a string token value from a table member
   * @function Kv3GetMemberStringToken
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (uint32): Default token value to return if member not found
   * @return uint32: String token hash code or defaultValue
   */
  inline uint32_t Kv3GetMemberStringToken(void* kv, plg::string name, uint32_t defaultValue) {
    using Kv3GetMemberStringTokenFn = uint32_t (*)(void*, plg::string, uint32_t);
    static Kv3GetMemberStringTokenFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberStringToken", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets an entity handle value from a table member
   * @function Kv3GetMemberEHandle
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (int32): Default entity handle value to return if member not found
   * @return int32: Entity handle as int32_t or defaultValue
   */
  inline int32_t Kv3GetMemberEHandle(void* kv, plg::string name, int32_t defaultValue) {
    using Kv3GetMemberEHandleFn = int32_t (*)(void*, plg::string, int32_t);
    static Kv3GetMemberEHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberEHandle", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a string value from a table member
   * @function Kv3GetMemberString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (string): Default string to return if member not found
   * @return string: String value or defaultValue
   */
  inline plg::string Kv3GetMemberString(void* kv, plg::string name, plg::string defaultValue) {
    using Kv3GetMemberStringFn = plg::string (*)(void*, plg::string, plg::string);
    static Kv3GetMemberStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberString", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a color value from a table member
   * @function Kv3GetMemberColor
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (int32): Default color value to return if member not found
   * @return int32: Color value as int32_t or defaultValue
   */
  inline int32_t Kv3GetMemberColor(void* kv, plg::string name, int32_t defaultValue) {
    using Kv3GetMemberColorFn = int32_t (*)(void*, plg::string, int32_t);
    static Kv3GetMemberColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberColor", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a 3D vector value from a table member
   * @function Kv3GetMemberVector
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (vec3): Default vector to return if member not found
   * @return vec3: 3D vector or defaultValue
   */
  inline plg::vec3 Kv3GetMemberVector(void* kv, plg::string name, plg::vec3 defaultValue) {
    using Kv3GetMemberVectorFn = plg::vec3 (*)(void*, plg::string, plg::vec3);
    static Kv3GetMemberVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberVector", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a 2D vector value from a table member
   * @function Kv3GetMemberVector2D
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (vec2): Default 2D vector to return if member not found
   * @return vec2: 2D vector or defaultValue
   */
  inline plg::vec2 Kv3GetMemberVector2D(void* kv, plg::string name, plg::vec2 defaultValue) {
    using Kv3GetMemberVector2DFn = plg::vec2 (*)(void*, plg::string, plg::vec2);
    static Kv3GetMemberVector2DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberVector2D", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a 4D vector value from a table member
   * @function Kv3GetMemberVector4D
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (vec4): Default 4D vector to return if member not found
   * @return vec4: 4D vector or defaultValue
   */
  inline plg::vec4 Kv3GetMemberVector4D(void* kv, plg::string name, plg::vec4 defaultValue) {
    using Kv3GetMemberVector4DFn = plg::vec4 (*)(void*, plg::string, plg::vec4);
    static Kv3GetMemberVector4DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberVector4D", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a quaternion value from a table member
   * @function Kv3GetMemberQuaternion
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (vec4): Default quaternion to return if member not found
   * @return vec4: Quaternion as vec4 or defaultValue
   */
  inline plg::vec4 Kv3GetMemberQuaternion(void* kv, plg::string name, plg::vec4 defaultValue) {
    using Kv3GetMemberQuaternionFn = plg::vec4 (*)(void*, plg::string, plg::vec4);
    static Kv3GetMemberQuaternionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberQuaternion", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets an angle (QAngle) value from a table member
   * @function Kv3GetMemberQAngle
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (vec3): Default angle to return if member not found
   * @return vec3: QAngle as vec3 or defaultValue
   */
  inline plg::vec3 Kv3GetMemberQAngle(void* kv, plg::string name, plg::vec3 defaultValue) {
    using Kv3GetMemberQAngleFn = plg::vec3 (*)(void*, plg::string, plg::vec3);
    static Kv3GetMemberQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberQAngle", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Gets a 3x4 matrix value from a table member
   * @function Kv3GetMemberMatrix3x4
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param defaultValue (mat4x4): Default matrix to return if member not found
   * @return mat4x4: 3x4 matrix as mat4x4 or defaultValue
   */
  inline plg::mat4x4 Kv3GetMemberMatrix3x4(void* kv, plg::string name, plg::mat4x4 defaultValue) {
    using Kv3GetMemberMatrix3x4Fn = plg::mat4x4 (*)(void*, plg::string, plg::mat4x4);
    static Kv3GetMemberMatrix3x4Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3GetMemberMatrix3x4", reinterpret_cast<void**>(&__func));
    return __func(kv, name, defaultValue);
  }

  /**
   * @brief Sets a table member to null
   * @function Kv3SetMemberToNull
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   */
  inline void Kv3SetMemberToNull(void* kv, plg::string name) {
    using Kv3SetMemberToNullFn = void (*)(void*, plg::string);
    static Kv3SetMemberToNullFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberToNull", reinterpret_cast<void**>(&__func));
    __func(kv, name);
  }

  /**
   * @brief Sets a table member to an empty array
   * @function Kv3SetMemberToEmptyArray
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   */
  inline void Kv3SetMemberToEmptyArray(void* kv, plg::string name) {
    using Kv3SetMemberToEmptyArrayFn = void (*)(void*, plg::string);
    static Kv3SetMemberToEmptyArrayFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberToEmptyArray", reinterpret_cast<void**>(&__func));
    __func(kv, name);
  }

  /**
   * @brief Sets a table member to an empty table
   * @function Kv3SetMemberToEmptyTable
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   */
  inline void Kv3SetMemberToEmptyTable(void* kv, plg::string name) {
    using Kv3SetMemberToEmptyTableFn = void (*)(void*, plg::string);
    static Kv3SetMemberToEmptyTableFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberToEmptyTable", reinterpret_cast<void**>(&__func));
    __func(kv, name);
  }

  /**
   * @brief Sets a table member to a binary blob (copies the data)
   * @function Kv3SetMemberToBinaryBlob
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param blob (uint8[]): Vector containing the binary blob data
   */
  inline void Kv3SetMemberToBinaryBlob(void* kv, plg::string name, plg::vector<uint8_t> blob) {
    using Kv3SetMemberToBinaryBlobFn = void (*)(void*, plg::string, plg::vector<uint8_t>);
    static Kv3SetMemberToBinaryBlobFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberToBinaryBlob", reinterpret_cast<void**>(&__func));
    __func(kv, name, blob);
  }

  /**
   * @brief Sets a table member to an external binary blob (does not copy)
   * @function Kv3SetMemberToBinaryBlobExternal
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param blob (uint8[]): Vector containing the external binary blob data
   * @param free_mem (bool): Whether to free the memory when the object is destroyed
   */
  inline void Kv3SetMemberToBinaryBlobExternal(void* kv, plg::string name, plg::vector<uint8_t> blob, bool free_mem) {
    using Kv3SetMemberToBinaryBlobExternalFn = void (*)(void*, plg::string, plg::vector<uint8_t>, bool);
    static Kv3SetMemberToBinaryBlobExternalFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberToBinaryBlobExternal", reinterpret_cast<void**>(&__func));
    __func(kv, name, blob, free_mem);
  }

  /**
   * @brief Sets a table member to a copy of another KeyValues3 value
   * @function Kv3SetMemberToCopyOfValue
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param other (ptr64): Pointer to the KeyValues3 object to copy
   */
  inline void Kv3SetMemberToCopyOfValue(void* kv, plg::string name, void* other) {
    using Kv3SetMemberToCopyOfValueFn = void (*)(void*, plg::string, void*);
    static Kv3SetMemberToCopyOfValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberToCopyOfValue", reinterpret_cast<void**>(&__func));
    __func(kv, name, other);
  }

  /**
   * @brief Sets a table member to a boolean value
   * @function Kv3SetMemberBool
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (bool): Boolean value to set
   */
  inline void Kv3SetMemberBool(void* kv, plg::string name, bool value) {
    using Kv3SetMemberBoolFn = void (*)(void*, plg::string, bool);
    static Kv3SetMemberBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberBool", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to a char value
   * @function Kv3SetMemberChar
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (char8): Char value to set
   */
  inline void Kv3SetMemberChar(void* kv, plg::string name, char value) {
    using Kv3SetMemberCharFn = void (*)(void*, plg::string, char);
    static Kv3SetMemberCharFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberChar", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to a 32-bit Unicode character value
   * @function Kv3SetMemberUChar32
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (uint32): 32-bit Unicode character value to set
   */
  inline void Kv3SetMemberUChar32(void* kv, plg::string name, uint32_t value) {
    using Kv3SetMemberUChar32Fn = void (*)(void*, plg::string, uint32_t);
    static Kv3SetMemberUChar32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberUChar32", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to a signed 8-bit integer value
   * @function Kv3SetMemberInt8
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (int8): int8_t value to set
   */
  inline void Kv3SetMemberInt8(void* kv, plg::string name, int8_t value) {
    using Kv3SetMemberInt8Fn = void (*)(void*, plg::string, int8_t);
    static Kv3SetMemberInt8Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberInt8", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to an unsigned 8-bit integer value
   * @function Kv3SetMemberUInt8
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (uint8): uint8_t value to set
   */
  inline void Kv3SetMemberUInt8(void* kv, plg::string name, uint8_t value) {
    using Kv3SetMemberUInt8Fn = void (*)(void*, plg::string, uint8_t);
    static Kv3SetMemberUInt8Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberUInt8", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to a signed 16-bit integer value
   * @function Kv3SetMemberShort
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (int16): int16_t value to set
   */
  inline void Kv3SetMemberShort(void* kv, plg::string name, int16_t value) {
    using Kv3SetMemberShortFn = void (*)(void*, plg::string, int16_t);
    static Kv3SetMemberShortFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberShort", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to an unsigned 16-bit integer value
   * @function Kv3SetMemberUShort
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (uint16): uint16_t value to set
   */
  inline void Kv3SetMemberUShort(void* kv, plg::string name, uint16_t value) {
    using Kv3SetMemberUShortFn = void (*)(void*, plg::string, uint16_t);
    static Kv3SetMemberUShortFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberUShort", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to a signed 32-bit integer value
   * @function Kv3SetMemberInt
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (int32): int32_t value to set
   */
  inline void Kv3SetMemberInt(void* kv, plg::string name, int32_t value) {
    using Kv3SetMemberIntFn = void (*)(void*, plg::string, int32_t);
    static Kv3SetMemberIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberInt", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to an unsigned 32-bit integer value
   * @function Kv3SetMemberUInt
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (uint32): uint32_t value to set
   */
  inline void Kv3SetMemberUInt(void* kv, plg::string name, uint32_t value) {
    using Kv3SetMemberUIntFn = void (*)(void*, plg::string, uint32_t);
    static Kv3SetMemberUIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberUInt", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to a signed 64-bit integer value
   * @function Kv3SetMemberInt64
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (int64): int64_t value to set
   */
  inline void Kv3SetMemberInt64(void* kv, plg::string name, int64_t value) {
    using Kv3SetMemberInt64Fn = void (*)(void*, plg::string, int64_t);
    static Kv3SetMemberInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberInt64", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to an unsigned 64-bit integer value
   * @function Kv3SetMemberUInt64
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (uint64): uint64_t value to set
   */
  inline void Kv3SetMemberUInt64(void* kv, plg::string name, uint64_t value) {
    using Kv3SetMemberUInt64Fn = void (*)(void*, plg::string, uint64_t);
    static Kv3SetMemberUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberUInt64", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to a float value
   * @function Kv3SetMemberFloat
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (float): Float value to set
   */
  inline void Kv3SetMemberFloat(void* kv, plg::string name, float value) {
    using Kv3SetMemberFloatFn = void (*)(void*, plg::string, float);
    static Kv3SetMemberFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberFloat", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to a double value
   * @function Kv3SetMemberDouble
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param value (double): Double value to set
   */
  inline void Kv3SetMemberDouble(void* kv, plg::string name, double value) {
    using Kv3SetMemberDoubleFn = void (*)(void*, plg::string, double);
    static Kv3SetMemberDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberDouble", reinterpret_cast<void**>(&__func));
    __func(kv, name, value);
  }

  /**
   * @brief Sets a table member to a pointer value
   * @function Kv3SetMemberPointer
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param ptr (ptr64): Pointer value as uintptr_t to set
   */
  inline void Kv3SetMemberPointer(void* kv, plg::string name, void* ptr) {
    using Kv3SetMemberPointerFn = void (*)(void*, plg::string, void*);
    static Kv3SetMemberPointerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberPointer", reinterpret_cast<void**>(&__func));
    __func(kv, name, ptr);
  }

  /**
   * @brief Sets a table member to a string token value
   * @function Kv3SetMemberStringToken
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param token (uint32): String token hash code to set
   */
  inline void Kv3SetMemberStringToken(void* kv, plg::string name, uint32_t token) {
    using Kv3SetMemberStringTokenFn = void (*)(void*, plg::string, uint32_t);
    static Kv3SetMemberStringTokenFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberStringToken", reinterpret_cast<void**>(&__func));
    __func(kv, name, token);
  }

  /**
   * @brief Sets a table member to an entity handle value
   * @function Kv3SetMemberEHandle
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param ehandle (int32): Entity handle value to set
   */
  inline void Kv3SetMemberEHandle(void* kv, plg::string name, int32_t ehandle) {
    using Kv3SetMemberEHandleFn = void (*)(void*, plg::string, int32_t);
    static Kv3SetMemberEHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberEHandle", reinterpret_cast<void**>(&__func));
    __func(kv, name, ehandle);
  }

  /**
   * @brief Sets a table member to a string value (copies the string)
   * @function Kv3SetMemberString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param str (string): String value to set
   * @param subtype (uint8): String subtype enumeration value
   */
  inline void Kv3SetMemberString(void* kv, plg::string name, plg::string str, uint8_t subtype) {
    using Kv3SetMemberStringFn = void (*)(void*, plg::string, plg::string, uint8_t);
    static Kv3SetMemberStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberString", reinterpret_cast<void**>(&__func));
    __func(kv, name, str, subtype);
  }

  /**
   * @brief Sets a table member to an external string value (does not copy)
   * @function Kv3SetMemberStringExternal
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param str (string): External string value to reference
   * @param subtype (uint8): String subtype enumeration value
   */
  inline void Kv3SetMemberStringExternal(void* kv, plg::string name, plg::string str, uint8_t subtype) {
    using Kv3SetMemberStringExternalFn = void (*)(void*, plg::string, plg::string, uint8_t);
    static Kv3SetMemberStringExternalFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberStringExternal", reinterpret_cast<void**>(&__func));
    __func(kv, name, str, subtype);
  }

  /**
   * @brief Sets a table member to a color value
   * @function Kv3SetMemberColor
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param color (int32): Color value as int32_t to set
   */
  inline void Kv3SetMemberColor(void* kv, plg::string name, int32_t color) {
    using Kv3SetMemberColorFn = void (*)(void*, plg::string, int32_t);
    static Kv3SetMemberColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberColor", reinterpret_cast<void**>(&__func));
    __func(kv, name, color);
  }

  /**
   * @brief Sets a table member to a 3D vector value
   * @function Kv3SetMemberVector
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param vec (vec3): 3D vector to set
   */
  inline void Kv3SetMemberVector(void* kv, plg::string name, plg::vec3 vec) {
    using Kv3SetMemberVectorFn = void (*)(void*, plg::string, plg::vec3);
    static Kv3SetMemberVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberVector", reinterpret_cast<void**>(&__func));
    __func(kv, name, vec);
  }

  /**
   * @brief Sets a table member to a 2D vector value
   * @function Kv3SetMemberVector2D
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param vec2d (vec2): 2D vector to set
   */
  inline void Kv3SetMemberVector2D(void* kv, plg::string name, plg::vec2 vec2d) {
    using Kv3SetMemberVector2DFn = void (*)(void*, plg::string, plg::vec2);
    static Kv3SetMemberVector2DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberVector2D", reinterpret_cast<void**>(&__func));
    __func(kv, name, vec2d);
  }

  /**
   * @brief Sets a table member to a 4D vector value
   * @function Kv3SetMemberVector4D
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param vec4d (vec4): 4D vector to set
   */
  inline void Kv3SetMemberVector4D(void* kv, plg::string name, plg::vec4 vec4d) {
    using Kv3SetMemberVector4DFn = void (*)(void*, plg::string, plg::vec4);
    static Kv3SetMemberVector4DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberVector4D", reinterpret_cast<void**>(&__func));
    __func(kv, name, vec4d);
  }

  /**
   * @brief Sets a table member to a quaternion value
   * @function Kv3SetMemberQuaternion
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param quat (vec4): Quaternion to set (as vec4)
   */
  inline void Kv3SetMemberQuaternion(void* kv, plg::string name, plg::vec4 quat) {
    using Kv3SetMemberQuaternionFn = void (*)(void*, plg::string, plg::vec4);
    static Kv3SetMemberQuaternionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberQuaternion", reinterpret_cast<void**>(&__func));
    __func(kv, name, quat);
  }

  /**
   * @brief Sets a table member to an angle (QAngle) value
   * @function Kv3SetMemberQAngle
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param ang (vec3): QAngle to set (as vec3)
   */
  inline void Kv3SetMemberQAngle(void* kv, plg::string name, plg::vec3 ang) {
    using Kv3SetMemberQAngleFn = void (*)(void*, plg::string, plg::vec3);
    static Kv3SetMemberQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberQAngle", reinterpret_cast<void**>(&__func));
    __func(kv, name, ang);
  }

  /**
   * @brief Sets a table member to a 3x4 matrix value
   * @function Kv3SetMemberMatrix3x4
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param name (string): Name of the member
   * @param matrix (mat4x4): 3x4 matrix to set (as mat4x4)
   */
  inline void Kv3SetMemberMatrix3x4(void* kv, plg::string name, plg::mat4x4 matrix) {
    using Kv3SetMemberMatrix3x4Fn = void (*)(void*, plg::string, plg::mat4x4);
    static Kv3SetMemberMatrix3x4Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SetMemberMatrix3x4", reinterpret_cast<void**>(&__func));
    __func(kv, name, matrix);
  }

  /**
   * @brief Prints debug information about the KeyValues3 object
   * @function Kv3DebugPrint
   * @param kv (ptr64): Pointer to the KeyValues3 object
   */
  inline void Kv3DebugPrint(void* kv) {
    using Kv3DebugPrintFn = void (*)(void*);
    static Kv3DebugPrintFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3DebugPrint", reinterpret_cast<void**>(&__func));
    __func(kv);
  }

  /**
   * @brief Loads KeyValues3 data from a buffer into a context
   * @function Kv3LoadFromBuffer
   * @param context (ptr64): Pointer to the KeyValues3 context
   * @param error (string&): Output string for error messages
   * @param input (uint8[]): Vector containing the input buffer data
   * @param kv_name (string): Name for the KeyValues3 object
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromBuffer(void* context, const plg::string& error, plg::vector<uint8_t> input, plg::string kv_name, uint32_t flags) {
    using Kv3LoadFromBufferFn = bool (*)(void*, const plg::string&, plg::vector<uint8_t>, plg::string, uint32_t);
    static Kv3LoadFromBufferFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromBuffer", reinterpret_cast<void**>(&__func));
    return __func(context, error, input, kv_name, flags);
  }

  /**
   * @brief Loads KeyValues3 data from a buffer
   * @function Kv3Load
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param input (uint8[]): Vector containing the input buffer data
   * @param kv_name (string): Name for the KeyValues3 object
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3Load(void* kv, const plg::string& error, plg::vector<uint8_t> input, plg::string kv_name, uint32_t flags) {
    using Kv3LoadFn = bool (*)(void*, const plg::string&, plg::vector<uint8_t>, plg::string, uint32_t);
    static Kv3LoadFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3Load", reinterpret_cast<void**>(&__func));
    return __func(kv, error, input, kv_name, flags);
  }

  /**
   * @brief Loads KeyValues3 data from a text string
   * @function Kv3LoadFromText
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param input (string): Text string containing KV3 data
   * @param kv_name (string): Name for the KeyValues3 object
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromText(void* kv, const plg::string& error, plg::string input, plg::string kv_name, uint32_t flags) {
    using Kv3LoadFromTextFn = bool (*)(void*, const plg::string&, plg::string, plg::string, uint32_t);
    static Kv3LoadFromTextFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromText", reinterpret_cast<void**>(&__func));
    return __func(kv, error, input, kv_name, flags);
  }

  /**
   * @brief Loads KeyValues3 data from a file into a context
   * @function Kv3LoadFromFileToContext
   * @param context (ptr64): Pointer to the KeyValues3 context
   * @param error (string&): Output string for error messages
   * @param filename (string): Name of the file to load
   * @param path (string): Path to the file
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromFileToContext(void* context, const plg::string& error, plg::string filename, plg::string path, uint32_t flags) {
    using Kv3LoadFromFileToContextFn = bool (*)(void*, const plg::string&, plg::string, plg::string, uint32_t);
    static Kv3LoadFromFileToContextFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromFileToContext", reinterpret_cast<void**>(&__func));
    return __func(context, error, filename, path, flags);
  }

  /**
   * @brief Loads KeyValues3 data from a file
   * @function Kv3LoadFromFile
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param filename (string): Name of the file to load
   * @param path (string): Path to the file
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromFile(void* kv, const plg::string& error, plg::string filename, plg::string path, uint32_t flags) {
    using Kv3LoadFromFileFn = bool (*)(void*, const plg::string&, plg::string, plg::string, uint32_t);
    static Kv3LoadFromFileFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromFile", reinterpret_cast<void**>(&__func));
    return __func(kv, error, filename, path, flags);
  }

  /**
   * @brief Loads KeyValues3 data from a JSON string
   * @function Kv3LoadFromJSON
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param input (string): JSON string
   * @param kv_name (string): Name for the KeyValues3 object
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromJSON(void* kv, const plg::string& error, plg::string input, plg::string kv_name, uint32_t flags) {
    using Kv3LoadFromJSONFn = bool (*)(void*, const plg::string&, plg::string, plg::string, uint32_t);
    static Kv3LoadFromJSONFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromJSON", reinterpret_cast<void**>(&__func));
    return __func(kv, error, input, kv_name, flags);
  }

  /**
   * @brief Loads KeyValues3 data from a JSON file
   * @function Kv3LoadFromJSONFile
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param path (string): Path to the file
   * @param filename (string): Name of the file to load
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromJSONFile(void* kv, const plg::string& error, plg::string path, plg::string filename, uint32_t flags) {
    using Kv3LoadFromJSONFileFn = bool (*)(void*, const plg::string&, plg::string, plg::string, uint32_t);
    static Kv3LoadFromJSONFileFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromJSONFile", reinterpret_cast<void**>(&__func));
    return __func(kv, error, path, filename, flags);
  }

  /**
   * @brief Loads KeyValues3 data from a KeyValues1 file
   * @function Kv3LoadFromKV1File
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param path (string): Path to the file
   * @param filename (string): Name of the file to load
   * @param esc_behavior (uint8): Escape sequence behavior for KV1 text
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromKV1File(void* kv, const plg::string& error, plg::string path, plg::string filename, uint8_t esc_behavior, uint32_t flags) {
    using Kv3LoadFromKV1FileFn = bool (*)(void*, const plg::string&, plg::string, plg::string, uint8_t, uint32_t);
    static Kv3LoadFromKV1FileFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromKV1File", reinterpret_cast<void**>(&__func));
    return __func(kv, error, path, filename, esc_behavior, flags);
  }

  /**
   * @brief Loads KeyValues3 data from a KeyValues1 text string
   * @function Kv3LoadFromKV1Text
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param input (string): KV1 text string
   * @param esc_behavior (uint8): Escape sequence behavior for KV1 text
   * @param kv_name (string): Name for the KeyValues3 object
   * @param unk (bool): Unknown boolean parameter
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromKV1Text(void* kv, const plg::string& error, plg::string input, uint8_t esc_behavior, plg::string kv_name, bool unk, uint32_t flags) {
    using Kv3LoadFromKV1TextFn = bool (*)(void*, const plg::string&, plg::string, uint8_t, plg::string, bool, uint32_t);
    static Kv3LoadFromKV1TextFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromKV1Text", reinterpret_cast<void**>(&__func));
    return __func(kv, error, input, esc_behavior, kv_name, unk, flags);
  }

  /**
   * @brief Loads KeyValues3 data from a KeyValues1 text string with translation
   * @function Kv3LoadFromKV1TextTranslated
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param input (string): KV1 text string
   * @param esc_behavior (uint8): Escape sequence behavior for KV1 text
   * @param translation (ptr64): Pointer to translation table
   * @param unk1 (int32): Unknown integer parameter
   * @param kv_name (string): Name for the KeyValues3 object
   * @param unk2 (bool): Unknown boolean parameter
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromKV1TextTranslated(void* kv, const plg::string& error, plg::string input, uint8_t esc_behavior, void* translation, int32_t unk1, plg::string kv_name, bool unk2, uint32_t flags) {
    using Kv3LoadFromKV1TextTranslatedFn = bool (*)(void*, const plg::string&, plg::string, uint8_t, void*, int32_t, plg::string, bool, uint32_t);
    static Kv3LoadFromKV1TextTranslatedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromKV1TextTranslated", reinterpret_cast<void**>(&__func));
    return __func(kv, error, input, esc_behavior, translation, unk1, kv_name, unk2, flags);
  }

  /**
   * @brief Loads data from a buffer that may be KV3 or KV1 format
   * @function Kv3LoadFromKV3OrKV1
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param input (uint8[]): Vector containing the input buffer data
   * @param kv_name (string): Name for the KeyValues3 object
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromKV3OrKV1(void* kv, const plg::string& error, plg::vector<uint8_t> input, plg::string kv_name, uint32_t flags) {
    using Kv3LoadFromKV3OrKV1Fn = bool (*)(void*, const plg::string&, plg::vector<uint8_t>, plg::string, uint32_t);
    static Kv3LoadFromKV3OrKV1Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromKV3OrKV1", reinterpret_cast<void**>(&__func));
    return __func(kv, error, input, kv_name, flags);
  }

  /**
   * @brief Loads KeyValues3 data from old schema text format
   * @function Kv3LoadFromOldSchemaText
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param input (uint8[]): Vector containing the input buffer data
   * @param kv_name (string): Name for the KeyValues3 object
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadFromOldSchemaText(void* kv, const plg::string& error, plg::vector<uint8_t> input, plg::string kv_name, uint32_t flags) {
    using Kv3LoadFromOldSchemaTextFn = bool (*)(void*, const plg::string&, plg::vector<uint8_t>, plg::string, uint32_t);
    static Kv3LoadFromOldSchemaTextFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadFromOldSchemaText", reinterpret_cast<void**>(&__func));
    return __func(kv, error, input, kv_name, flags);
  }

  /**
   * @brief Loads KeyValues3 text without a header
   * @function Kv3LoadTextNoHeader
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param input (string): Text string containing KV3 data
   * @param kv_name (string): Name for the KeyValues3 object
   * @param flags (uint32): Loading flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3LoadTextNoHeader(void* kv, const plg::string& error, plg::string input, plg::string kv_name, uint32_t flags) {
    using Kv3LoadTextNoHeaderFn = bool (*)(void*, const plg::string&, plg::string, plg::string, uint32_t);
    static Kv3LoadTextNoHeaderFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3LoadTextNoHeader", reinterpret_cast<void**>(&__func));
    return __func(kv, error, input, kv_name, flags);
  }

  /**
   * @brief Saves KeyValues3 data to a buffer
   * @function Kv3Save
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param output (uint8[]&): Vector to store the output buffer data
   * @param flags (uint32): Saving flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3Save(void* kv, const plg::string& error, plg::vector<uint8_t> output, uint32_t flags) {
    using Kv3SaveFn = bool (*)(void*, const plg::string&, plg::vector<uint8_t>, uint32_t);
    static Kv3SaveFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3Save", reinterpret_cast<void**>(&__func));
    return __func(kv, error, output, flags);
  }

  /**
   * @brief Saves KeyValues3 data as JSON to a buffer
   * @function Kv3SaveAsJSON
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param output (uint8[]&): Vector to store the output JSON data
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3SaveAsJSON(void* kv, const plg::string& error, plg::vector<uint8_t> output) {
    using Kv3SaveAsJSONFn = bool (*)(void*, const plg::string&, plg::vector<uint8_t>);
    static Kv3SaveAsJSONFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SaveAsJSON", reinterpret_cast<void**>(&__func));
    return __func(kv, error, output);
  }

  /**
   * @brief Saves KeyValues3 data as a JSON string
   * @function Kv3SaveAsJSONString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param output (string&): String to store the JSON output
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3SaveAsJSONString(void* kv, const plg::string& error, const plg::string& output) {
    using Kv3SaveAsJSONStringFn = bool (*)(void*, const plg::string&, const plg::string&);
    static Kv3SaveAsJSONStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SaveAsJSONString", reinterpret_cast<void**>(&__func));
    return __func(kv, error, output);
  }

  /**
   * @brief Saves KeyValues3 data as KeyValues1 text to a buffer
   * @function Kv3SaveAsKV1Text
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param output (uint8[]&): Vector to store the output KV1 text data
   * @param esc_behavior (uint8): Escape sequence behavior for KV1 text
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3SaveAsKV1Text(void* kv, const plg::string& error, plg::vector<uint8_t> output, uint8_t esc_behavior) {
    using Kv3SaveAsKV1TextFn = bool (*)(void*, const plg::string&, plg::vector<uint8_t>, uint8_t);
    static Kv3SaveAsKV1TextFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SaveAsKV1Text", reinterpret_cast<void**>(&__func));
    return __func(kv, error, output, esc_behavior);
  }

  /**
   * @brief Saves KeyValues3 data as KeyValues1 text with translation to a buffer
   * @function Kv3SaveAsKV1TextTranslated
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param output (uint8[]&): Vector to store the output KV1 text data
   * @param esc_behavior (uint8): Escape sequence behavior for KV1 text
   * @param translation (ptr64): Pointer to translation table
   * @param unk (int32): Unknown integer parameter
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3SaveAsKV1TextTranslated(void* kv, const plg::string& error, plg::vector<uint8_t> output, uint8_t esc_behavior, void* translation, int32_t unk) {
    using Kv3SaveAsKV1TextTranslatedFn = bool (*)(void*, const plg::string&, plg::vector<uint8_t>, uint8_t, void*, int32_t);
    static Kv3SaveAsKV1TextTranslatedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SaveAsKV1TextTranslated", reinterpret_cast<void**>(&__func));
    return __func(kv, error, output, esc_behavior, translation, unk);
  }

  /**
   * @brief Saves KeyValues3 text without a header to a buffer
   * @function Kv3SaveTextNoHeaderToBuffer
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param output (uint8[]&): Vector to store the output text data
   * @param flags (uint32): Saving flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3SaveTextNoHeaderToBuffer(void* kv, const plg::string& error, plg::vector<uint8_t> output, uint32_t flags) {
    using Kv3SaveTextNoHeaderToBufferFn = bool (*)(void*, const plg::string&, plg::vector<uint8_t>, uint32_t);
    static Kv3SaveTextNoHeaderToBufferFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SaveTextNoHeaderToBuffer", reinterpret_cast<void**>(&__func));
    return __func(kv, error, output, flags);
  }

  /**
   * @brief Saves KeyValues3 text without a header to a string
   * @function Kv3SaveTextNoHeader
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param output (string&): String to store the text output
   * @param flags (uint32): Saving flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3SaveTextNoHeader(void* kv, const plg::string& error, const plg::string& output, uint32_t flags) {
    using Kv3SaveTextNoHeaderFn = bool (*)(void*, const plg::string&, const plg::string&, uint32_t);
    static Kv3SaveTextNoHeaderFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SaveTextNoHeader", reinterpret_cast<void**>(&__func));
    return __func(kv, error, output, flags);
  }

  /**
   * @brief Saves KeyValues3 text to a string
   * @function Kv3SaveTextToString
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param output (string&): String to store the text output
   * @param flags (uint32): Saving flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3SaveTextToString(void* kv, const plg::string& error, const plg::string& output, uint32_t flags) {
    using Kv3SaveTextToStringFn = bool (*)(void*, const plg::string&, const plg::string&, uint32_t);
    static Kv3SaveTextToStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SaveTextToString", reinterpret_cast<void**>(&__func));
    return __func(kv, error, output, flags);
  }

  /**
   * @brief Saves KeyValues3 data to a file
   * @function Kv3SaveToFile
   * @param kv (ptr64): Pointer to the KeyValues3 object
   * @param error (string&): Output string for error messages
   * @param filename (string): Name of the file to save
   * @param path (string): Path to save the file
   * @param flags (uint32): Saving flags
   * @return bool: true if successful, false otherwise
   */
  inline bool Kv3SaveToFile(void* kv, const plg::string& error, plg::string filename, plg::string path, uint32_t flags) {
    using Kv3SaveToFileFn = bool (*)(void*, const plg::string&, plg::string, plg::string, uint32_t);
    static Kv3SaveToFileFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Kv3SaveToFile", reinterpret_cast<void**>(&__func));
    return __func(kv, error, filename, path, flags);
  }

  /**
   * @brief Triggers a breakpoint in the debugger.
   * @function DebugBreak
   */
  inline void DebugBreak() {
    using DebugBreakFn = void (*)();
    static DebugBreakFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugBreak", reinterpret_cast<void**>(&__func));
    __func();
  }

  /**
   * @brief Draws a debug overlay box.
   * @function DebugDrawBox
   * @param center (vec3): Center of the box in world space.
   * @param mins (vec3): Minimum bounds relative to the center.
   * @param maxs (vec3): Maximum bounds relative to the center.
   * @param r (int32): Red color value.
   * @param g (int32): Green color value.
   * @param b (int32): Blue color value.
   * @param a (int32): Alpha (transparency) value.
   * @param duration (float): Duration (in seconds) to display the box.
   */
  inline void DebugDrawBox(plg::vec3 center, plg::vec3 mins, plg::vec3 maxs, int32_t r, int32_t g, int32_t b, int32_t a, float duration) {
    using DebugDrawBoxFn = void (*)(plg::vec3, plg::vec3, plg::vec3, int32_t, int32_t, int32_t, int32_t, float);
    static DebugDrawBoxFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugDrawBox", reinterpret_cast<void**>(&__func));
    __func(center, mins, maxs, r, g, b, a, duration);
  }

  /**
   * @brief Draws a debug box oriented in the direction of a forward vector.
   * @function DebugDrawBoxDirection
   * @param center (vec3): Center of the box.
   * @param mins (vec3): Minimum bounds.
   * @param maxs (vec3): Maximum bounds.
   * @param forward (vec3): Forward direction vector.
   * @param color (vec3): RGB color vector.
   * @param alpha (float): Alpha transparency.
   * @param duration (float): Duration (in seconds) to display the box.
   */
  inline void DebugDrawBoxDirection(plg::vec3 center, plg::vec3 mins, plg::vec3 maxs, plg::vec3 forward, plg::vec3 color, float alpha, float duration) {
    using DebugDrawBoxDirectionFn = void (*)(plg::vec3, plg::vec3, plg::vec3, plg::vec3, plg::vec3, float, float);
    static DebugDrawBoxDirectionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugDrawBoxDirection", reinterpret_cast<void**>(&__func));
    __func(center, mins, maxs, forward, color, alpha, duration);
  }

  /**
   * @brief Draws a debug circle.
   * @function DebugDrawCircle
   * @param center (vec3): Center of the circle.
   * @param color (vec3): RGB color vector.
   * @param alpha (float): Alpha transparency.
   * @param radius (float): Circle radius.
   * @param zTest (bool): Whether to perform depth testing.
   * @param duration (float): Duration (in seconds) to display the circle.
   */
  inline void DebugDrawCircle(plg::vec3 center, plg::vec3 color, float alpha, float radius, bool zTest, float duration) {
    using DebugDrawCircleFn = void (*)(plg::vec3, plg::vec3, float, float, bool, float);
    static DebugDrawCircleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugDrawCircle", reinterpret_cast<void**>(&__func));
    __func(center, color, alpha, radius, zTest, duration);
  }

  /**
   * @brief Clears all debug overlays.
   * @function DebugDrawClear
   */
  inline void DebugDrawClear() {
    using DebugDrawClearFn = void (*)();
    static DebugDrawClearFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugDrawClear", reinterpret_cast<void**>(&__func));
    __func();
  }

  /**
   * @brief Draws a debug overlay line.
   * @function DebugDrawLine
   * @param origin (vec3): Start point of the line.
   * @param target (vec3): End point of the line.
   * @param r (int32): Red color value.
   * @param g (int32): Green color value.
   * @param b (int32): Blue color value.
   * @param zTest (bool): Whether to perform depth testing.
   * @param duration (float): Duration (in seconds) to display the line.
   */
  inline void DebugDrawLine(plg::vec3 origin, plg::vec3 target, int32_t r, int32_t g, int32_t b, bool zTest, float duration) {
    using DebugDrawLineFn = void (*)(plg::vec3, plg::vec3, int32_t, int32_t, int32_t, bool, float);
    static DebugDrawLineFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugDrawLine", reinterpret_cast<void**>(&__func));
    __func(origin, target, r, g, b, zTest, duration);
  }

  /**
   * @brief Draws a debug line using a color vector.
   * @function DebugDrawLine_vCol
   * @param start (vec3): Start point of the line.
   * @param end (vec3): End point of the line.
   * @param color (vec3): RGB color vector.
   * @param zTest (bool): Whether to perform depth testing.
   * @param duration (float): Duration (in seconds) to display the line.
   */
  inline void DebugDrawLine_vCol(plg::vec3 start, plg::vec3 end, plg::vec3 color, bool zTest, float duration) {
    using DebugDrawLine_vColFn = void (*)(plg::vec3, plg::vec3, plg::vec3, bool, float);
    static DebugDrawLine_vColFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugDrawLine_vCol", reinterpret_cast<void**>(&__func));
    __func(start, end, color, zTest, duration);
  }

  /**
   * @brief Draws text at a specified screen position with line offset.
   * @function DebugDrawScreenTextLine
   * @param x (float): X coordinate in screen space.
   * @param y (float): Y coordinate in screen space.
   * @param lineOffset (int32): Line offset value.
   * @param text (string): The text string to display.
   * @param r (int32): Red color value.
   * @param g (int32): Green color value.
   * @param b (int32): Blue color value.
   * @param a (int32): Alpha transparency value.
   * @param duration (float): Duration (in seconds) to display the text.
   */
  inline void DebugDrawScreenTextLine(float x, float y, int32_t lineOffset, plg::string text, int32_t r, int32_t g, int32_t b, int32_t a, float duration) {
    using DebugDrawScreenTextLineFn = void (*)(float, float, int32_t, plg::string, int32_t, int32_t, int32_t, int32_t, float);
    static DebugDrawScreenTextLineFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugDrawScreenTextLine", reinterpret_cast<void**>(&__func));
    __func(x, y, lineOffset, text, r, g, b, a, duration);
  }

  /**
   * @brief Draws a debug sphere.
   * @function DebugDrawSphere
   * @param center (vec3): Center of the sphere.
   * @param color (vec3): RGB color vector.
   * @param alpha (float): Alpha transparency.
   * @param radius (float): Radius of the sphere.
   * @param zTest (bool): Whether to perform depth testing.
   * @param duration (float): Duration (in seconds) to display the sphere.
   */
  inline void DebugDrawSphere(plg::vec3 center, plg::vec3 color, float alpha, float radius, bool zTest, float duration) {
    using DebugDrawSphereFn = void (*)(plg::vec3, plg::vec3, float, float, bool, float);
    static DebugDrawSphereFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugDrawSphere", reinterpret_cast<void**>(&__func));
    __func(center, color, alpha, radius, zTest, duration);
  }

  /**
   * @brief Draws text in 3D space.
   * @function DebugDrawText
   * @param origin (vec3): World-space position to draw the text at.
   * @param text (string): The text string to display.
   * @param viewCheck (bool): If true, only draws when visible to camera.
   * @param duration (float): Duration (in seconds) to display the text.
   */
  inline void DebugDrawText(plg::vec3 origin, plg::string text, bool viewCheck, float duration) {
    using DebugDrawTextFn = void (*)(plg::vec3, plg::string, bool, float);
    static DebugDrawTextFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugDrawText", reinterpret_cast<void**>(&__func));
    __func(origin, text, viewCheck, duration);
  }

  /**
   * @brief Draws styled debug text on screen.
   * @function DebugScreenTextPretty
   * @param x (float): X coordinate.
   * @param y (float): Y coordinate.
   * @param lineOffset (int32): Line offset value.
   * @param text (string): Text string.
   * @param r (int32): Red color value.
   * @param g (int32): Green color value.
   * @param b (int32): Blue color value.
   * @param a (int32): Alpha transparency.
   * @param duration (float): Duration (in seconds) to display the text.
   * @param font (string): Font name.
   * @param size (int32): Font size.
   * @param bold (bool): Whether text should be bold.
   */
  inline void DebugScreenTextPretty(float x, float y, int32_t lineOffset, plg::string text, int32_t r, int32_t g, int32_t b, int32_t a, float duration, plg::string font, int32_t size, bool bold) {
    using DebugScreenTextPrettyFn = void (*)(float, float, int32_t, plg::string, int32_t, int32_t, int32_t, int32_t, float, plg::string, int32_t, bool);
    static DebugScreenTextPrettyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugScreenTextPretty", reinterpret_cast<void**>(&__func));
    __func(x, y, lineOffset, text, r, g, b, a, duration, font, size, bold);
  }

  /**
   * @brief Performs an assertion and logs a message if the assertion fails.
   * @function DebugScriptAssert
   * @param assertion (bool): Boolean value to test.
   * @param message (string): Message to display if the assertion fails.
   */
  inline void DebugScriptAssert(bool assertion, plg::string message) {
    using DebugScriptAssertFn = void (*)(bool, plg::string);
    static DebugScriptAssertFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DebugScriptAssert", reinterpret_cast<void**>(&__func));
    __func(assertion, message);
  }

  /**
   * @brief Returns angular difference in degrees
   * @function AnglesDiff
   * @param angle1 (float): First angle in degrees
   * @param angle2 (float): Second angle in degrees
   * @return float: Angular difference in degrees
   */
  inline float AnglesDiff(float angle1, float angle2) {
    using AnglesDiffFn = float (*)(float, float);
    static AnglesDiffFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AnglesDiff", reinterpret_cast<void**>(&__func));
    return __func(angle1, angle2);
  }

  /**
   * @brief Converts QAngle to directional Vector
   * @function AnglesToVector
   * @param angles (vec3): The QAngle to convert
   * @return vec3: Directional vector
   */
  inline plg::vec3 AnglesToVector(plg::vec3 angles) {
    using AnglesToVectorFn = plg::vec3 (*)(plg::vec3);
    static AnglesToVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AnglesToVector", reinterpret_cast<void**>(&__func));
    return __func(angles);
  }

  /**
   * @brief Converts axis-angle representation to quaternion
   * @function AxisAngleToQuaternion
   * @param axis (vec3): Rotation axis (should be normalized)
   * @param angle (float): Rotation angle in radians
   * @return vec4: Resulting quaternion
   */
  inline plg::vec4 AxisAngleToQuaternion(plg::vec3 axis, float angle) {
    using AxisAngleToQuaternionFn = plg::vec4 (*)(plg::vec3, float);
    static AxisAngleToQuaternionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AxisAngleToQuaternion", reinterpret_cast<void**>(&__func));
    return __func(axis, angle);
  }

  /**
   * @brief Computes closest point on an entity's oriented bounding box (OBB)
   * @function CalcClosestPointOnEntityOBB
   * @param entityHandle (int32): Handle of the entity
   * @param position (vec3): Position to find closest point from
   * @return vec3: Closest point on the entity's OBB, or vec3_origin if entity is invalid
   */
  inline plg::vec3 CalcClosestPointOnEntityOBB(int32_t entityHandle, plg::vec3 position) {
    using CalcClosestPointOnEntityOBBFn = plg::vec3 (*)(int32_t, plg::vec3);
    static CalcClosestPointOnEntityOBBFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CalcClosestPointOnEntityOBB", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, position);
  }

  /**
   * @brief Computes distance between two entities' oriented bounding boxes (OBBs)
   * @function CalcDistanceBetweenEntityOBB
   * @param entityHandle1 (int32): Handle of the first entity
   * @param entityHandle2 (int32): Handle of the second entity
   * @return float: Distance between OBBs, or -1.0f if either entity is invalid
   */
  inline float CalcDistanceBetweenEntityOBB(int32_t entityHandle1, int32_t entityHandle2) {
    using CalcDistanceBetweenEntityOBBFn = float (*)(int32_t, int32_t);
    static CalcDistanceBetweenEntityOBBFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CalcDistanceBetweenEntityOBB", reinterpret_cast<void**>(&__func));
    return __func(entityHandle1, entityHandle2);
  }

  /**
   * @brief Computes shortest 2D distance from a point to a line segment
   * @function CalcDistanceToLineSegment2D
   * @param p (vec3): The point
   * @param vLineA (vec3): First endpoint of the line segment
   * @param vLineB (vec3): Second endpoint of the line segment
   * @return float: Shortest 2D distance
   */
  inline float CalcDistanceToLineSegment2D(plg::vec3 p, plg::vec3 vLineA, plg::vec3 vLineB) {
    using CalcDistanceToLineSegment2DFn = float (*)(plg::vec3, plg::vec3, plg::vec3);
    static CalcDistanceToLineSegment2DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CalcDistanceToLineSegment2D", reinterpret_cast<void**>(&__func));
    return __func(p, vLineA, vLineB);
  }

  /**
   * @brief Computes cross product of two vectors
   * @function CrossVectors
   * @param v1 (vec3): First vector
   * @param v2 (vec3): Second vector
   * @return vec3: Cross product vector (v1 × v2)
   */
  inline plg::vec3 CrossVectors(plg::vec3 v1, plg::vec3 v2) {
    using CrossVectorsFn = plg::vec3 (*)(plg::vec3, plg::vec3);
    static CrossVectorsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CrossVectors", reinterpret_cast<void**>(&__func));
    return __func(v1, v2);
  }

  /**
   * @brief Smooth exponential decay function
   * @function ExponentDecay
   * @param decayTo (float): Target value to decay towards
   * @param decayTime (float): Time constant for decay
   * @param dt (float): Delta time
   * @return float: Decay factor
   */
  inline float ExponentDecay(float decayTo, float decayTime, float dt) {
    using ExponentDecayFn = float (*)(float, float, float);
    static ExponentDecayFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ExponentDecay", reinterpret_cast<void**>(&__func));
    return __func(decayTo, decayTime, dt);
  }

  /**
   * @brief Linear interpolation between two vectors
   * @function LerpVectors
   * @param start (vec3): Starting vector
   * @param end (vec3): Ending vector
   * @param factor (float): Interpolation factor (0.0 to 1.0)
   * @return vec3: Interpolated vector
   */
  inline plg::vec3 LerpVectors(plg::vec3 start, plg::vec3 end, float factor) {
    using LerpVectorsFn = plg::vec3 (*)(plg::vec3, plg::vec3, float);
    static LerpVectorsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.LerpVectors", reinterpret_cast<void**>(&__func));
    return __func(start, end, factor);
  }

  /**
   * @brief Quaternion spherical linear interpolation for angles
   * @function QSlerp
   * @param fromAngle (vec3): Starting angle
   * @param toAngle (vec3): Ending angle
   * @param time (float): Interpolation time (0.0 to 1.0)
   * @return vec3: Interpolated angle
   */
  inline plg::vec3 QSlerp(plg::vec3 fromAngle, plg::vec3 toAngle, float time) {
    using QSlerpFn = plg::vec3 (*)(plg::vec3, plg::vec3, float);
    static QSlerpFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.QSlerp", reinterpret_cast<void**>(&__func));
    return __func(fromAngle, toAngle, time);
  }

  /**
   * @brief Rotate one QAngle by another
   * @function RotateOrientation
   * @param a1 (vec3): Base orientation
   * @param a2 (vec3): Rotation to apply
   * @return vec3: Rotated orientation
   */
  inline plg::vec3 RotateOrientation(plg::vec3 a1, plg::vec3 a2) {
    using RotateOrientationFn = plg::vec3 (*)(plg::vec3, plg::vec3);
    static RotateOrientationFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RotateOrientation", reinterpret_cast<void**>(&__func));
    return __func(a1, a2);
  }

  /**
   * @brief Rotate a vector around a point by specified angle
   * @function RotatePosition
   * @param rotationOrigin (vec3): Origin point of rotation
   * @param rotationAngle (vec3): Angle to rotate by
   * @param vectorToRotate (vec3): Vector to be rotated
   * @return vec3: Rotated vector
   */
  inline plg::vec3 RotatePosition(plg::vec3 rotationOrigin, plg::vec3 rotationAngle, plg::vec3 vectorToRotate) {
    using RotatePositionFn = plg::vec3 (*)(plg::vec3, plg::vec3, plg::vec3);
    static RotatePositionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RotatePosition", reinterpret_cast<void**>(&__func));
    return __func(rotationOrigin, rotationAngle, vectorToRotate);
  }

  /**
   * @brief Rotates quaternion by axis-angle representation
   * @function RotateQuaternionByAxisAngle
   * @param q (vec4): Quaternion to rotate
   * @param axis (vec3): Rotation axis
   * @param angle (float): Rotation angle in radians
   * @return vec4: Rotated quaternion
   */
  inline plg::vec4 RotateQuaternionByAxisAngle(plg::vec4 q, plg::vec3 axis, float angle) {
    using RotateQuaternionByAxisAngleFn = plg::vec4 (*)(plg::vec4, plg::vec3, float);
    static RotateQuaternionByAxisAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RotateQuaternionByAxisAngle", reinterpret_cast<void**>(&__func));
    return __func(q, axis, angle);
  }

  /**
   * @brief Finds angular delta between two QAngles
   * @function RotationDelta
   * @param src (vec3): Source angle
   * @param dest (vec3): Destination angle
   * @return vec3: Delta angle from src to dest
   */
  inline plg::vec3 RotationDelta(plg::vec3 src, plg::vec3 dest) {
    using RotationDeltaFn = plg::vec3 (*)(plg::vec3, plg::vec3);
    static RotationDeltaFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RotationDelta", reinterpret_cast<void**>(&__func));
    return __func(src, dest);
  }

  /**
   * @brief Converts delta QAngle to angular velocity vector
   * @function RotationDeltaAsAngularVelocity
   * @param a1 (vec3): First angle
   * @param a2 (vec3): Second angle
   * @return vec3: Angular velocity vector
   */
  inline plg::vec3 RotationDeltaAsAngularVelocity(plg::vec3 a1, plg::vec3 a2) {
    using RotationDeltaAsAngularVelocityFn = plg::vec3 (*)(plg::vec3, plg::vec3);
    static RotationDeltaAsAngularVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RotationDeltaAsAngularVelocity", reinterpret_cast<void**>(&__func));
    return __func(a1, a2);
  }

  /**
   * @brief Interpolates between two quaternions using spline
   * @function SplineQuaternions
   * @param q0 (vec4): Starting quaternion
   * @param q1 (vec4): Ending quaternion
   * @param t (float): Interpolation parameter (0.0 to 1.0)
   * @return vec4: Interpolated quaternion
   */
  inline plg::vec4 SplineQuaternions(plg::vec4 q0, plg::vec4 q1, float t) {
    using SplineQuaternionsFn = plg::vec4 (*)(plg::vec4, plg::vec4, float);
    static SplineQuaternionsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SplineQuaternions", reinterpret_cast<void**>(&__func));
    return __func(q0, q1, t);
  }

  /**
   * @brief Interpolates between two vectors using spline
   * @function SplineVectors
   * @param v0 (vec3): Starting vector
   * @param v1 (vec3): Ending vector
   * @param t (float): Interpolation parameter (0.0 to 1.0)
   * @return vec3: Interpolated vector
   */
  inline plg::vec3 SplineVectors(plg::vec3 v0, plg::vec3 v1, float t) {
    using SplineVectorsFn = plg::vec3 (*)(plg::vec3, plg::vec3, float);
    static SplineVectorsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SplineVectors", reinterpret_cast<void**>(&__func));
    return __func(v0, v1, t);
  }

  /**
   * @brief Converts directional vector to QAngle (no roll)
   * @function VectorToAngles
   * @param input (vec3): Direction vector
   * @return vec3: Angle representation with pitch and yaw (roll is 0)
   */
  inline plg::vec3 VectorToAngles(plg::vec3 input) {
    using VectorToAnglesFn = plg::vec3 (*)(plg::vec3);
    static VectorToAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.VectorToAngles", reinterpret_cast<void**>(&__func));
    return __func(input);
  }

  /**
   * @brief Returns random float between min and max
   * @function RandomFlt
   * @param min (float): Minimum value (inclusive)
   * @param max (float): Maximum value (inclusive)
   * @return float: Random float in range [min, max]
   */
  inline float RandomFlt(float min, float max) {
    using RandomFltFn = float (*)(float, float);
    static RandomFltFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RandomFlt", reinterpret_cast<void**>(&__func));
    return __func(min, max);
  }

  /**
   * @brief Returns random integer between min and max (inclusive)
   * @function RandomInt
   * @param min (int32): Minimum value (inclusive)
   * @param max (int32): Maximum value (inclusive)
   * @return int32: Random integer in range [min, max]
   */
  inline int32_t RandomInt(int32_t min, int32_t max) {
    using RandomIntFn = int32_t (*)(int32_t, int32_t);
    static RandomIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RandomInt", reinterpret_cast<void**>(&__func));
    return __func(min, max);
  }

  /**
   * @brief Performs a collideable trace using the VScript-compatible table call, exposing it through C++ exports.
   * @function TraceCollideable
   * @param start (vec3): Trace start position (world space)
   * @param end (vec3): Trace end position (world space)
   * @param entityHandle (int32): Entity handle of the collideable
   * @param outPos (vec3&): Output: position of impact
   * @param outFraction (double&): Output: fraction of trace completed
   * @param outHit (bool&): Output: whether a hit occurred
   * @param outStartSolid (bool&): Output: whether trace started inside solid
   * @param outNormal (vec3&): Output: surface normal at impact
   * @return bool: True if trace hit something, false otherwise
   */
  inline bool TraceCollideable(plg::vec3 start, plg::vec3 end, int32_t entityHandle, const plg::vec3& outPos, double outFraction, bool outHit, bool outStartSolid, const plg::vec3& outNormal) {
    using TraceCollideableFn = bool (*)(plg::vec3, plg::vec3, int32_t, const plg::vec3&, double, bool, bool, const plg::vec3&);
    static TraceCollideableFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TraceCollideable", reinterpret_cast<void**>(&__func));
    return __func(start, end, entityHandle, outPos, outFraction, outHit, outStartSolid, outNormal);
  }

  /**
   * @brief Performs a collideable trace using the VScript-compatible table call, exposing it through C++ exports.
   * @function TraceCollideable2
   * @param start (vec3): Trace start position (world space)
   * @param end (vec3): Trace end position (world space)
   * @param entityHandle (int32): Entity handle of the collideable
   * @param mins (ptr64): Bounding box minimums
   * @param maxs (ptr64): Bounding box maximums
   * @param outPos (vec3&): Output: position of impact
   * @param outFraction (double&): Output: fraction of trace completed
   * @param outHit (bool&): Output: whether a hit occurred
   * @param outStartSolid (bool&): Output: whether trace started inside solid
   * @param outNormal (vec3&): Output: surface normal at impact
   * @return bool: True if trace hit something, false otherwise
   */
  inline bool TraceCollideable2(plg::vec3 start, plg::vec3 end, int32_t entityHandle, void* mins, void* maxs, const plg::vec3& outPos, double outFraction, bool outHit, bool outStartSolid, const plg::vec3& outNormal) {
    using TraceCollideable2Fn = bool (*)(plg::vec3, plg::vec3, int32_t, void*, void*, const plg::vec3&, double, bool, bool, const plg::vec3&);
    static TraceCollideable2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TraceCollideable2", reinterpret_cast<void**>(&__func));
    return __func(start, end, entityHandle, mins, maxs, outPos, outFraction, outHit, outStartSolid, outNormal);
  }

  /**
   * @brief Performs a hull trace with specified dimensions and mask.
   * @function TraceHull
   * @param start (vec3): Trace start position
   * @param end (vec3): Trace end position
   * @param min (vec3): Local bounding box minimums
   * @param max (vec3): Local bounding box maximums
   * @param mask (int32): Trace mask
   * @param ignoreHandle (int32): Entity handle to ignore during trace
   * @param outPos (vec3&): Output: position of impact
   * @param outFraction (double&): Output: fraction of trace completed
   * @param outHit (bool&): Output: whether a hit occurred
   * @param outEntHit (int32&): Output: handle of entity hit
   * @param outStartSolid (bool&): Output: whether trace started inside solid
   * @return bool: True if trace hit something, false otherwise
   */
  inline bool TraceHull(plg::vec3 start, plg::vec3 end, plg::vec3 min, plg::vec3 max, int32_t mask, int32_t ignoreHandle, const plg::vec3& outPos, double outFraction, bool outHit, int32_t outEntHit, bool outStartSolid) {
    using TraceHullFn = bool (*)(plg::vec3, plg::vec3, plg::vec3, plg::vec3, int32_t, int32_t, const plg::vec3&, double, bool, int32_t, bool);
    static TraceHullFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TraceHull", reinterpret_cast<void**>(&__func));
    return __func(start, end, min, max, mask, ignoreHandle, outPos, outFraction, outHit, outEntHit, outStartSolid);
  }

  /**
   * @brief Performs a line trace between two points.
   * @function TraceLine
   * @param startPos (vec3): Trace start position
   * @param endPos (vec3): Trace end position
   * @param mask (int32): Trace mask
   * @param ignoreHandle (int32): Entity handle to ignore during trace
   * @param outPos (vec3&): Output: position of impact
   * @param outFraction (double&): Output: fraction of trace completed
   * @param outHit (bool&): Output: whether a hit occurred
   * @param outEntHit (int32&): Output: handle of entity hit
   * @param outStartSolid (bool&): Output: whether trace started inside solid
   * @return bool: True if trace hit something, false otherwise
   */
  inline bool TraceLine(plg::vec3 startPos, plg::vec3 endPos, int32_t mask, int32_t ignoreHandle, const plg::vec3& outPos, double outFraction, bool outHit, int32_t outEntHit, bool outStartSolid) {
    using TraceLineFn = bool (*)(plg::vec3, plg::vec3, int32_t, int32_t, const plg::vec3&, double, bool, int32_t, bool);
    static TraceLineFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TraceLine", reinterpret_cast<void**>(&__func));
    return __func(startPos, endPos, mask, ignoreHandle, outPos, outFraction, outHit, outEntHit, outStartSolid);
  }

  /**
   * @brief Applies an impulse to an entity at a specific world position.
   * @function AddBodyImpulseAtPosition
   * @param entityHandle (int32): The handle of the entity.
   * @param position (vec3): The world position where the impulse will be applied.
   * @param impulse (vec3): The impulse vector to apply.
   */
  inline void AddBodyImpulseAtPosition(int32_t entityHandle, plg::vec3 position, plg::vec3 impulse) {
    using AddBodyImpulseAtPositionFn = void (*)(int32_t, plg::vec3, plg::vec3);
    static AddBodyImpulseAtPositionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddBodyImpulseAtPosition", reinterpret_cast<void**>(&__func));
    __func(entityHandle, position, impulse);
  }

  /**
   * @brief Adds linear and angular velocity to the entity's physics object.
   * @function AddBodyVelocity
   * @param entityHandle (int32): The handle of the entity.
   * @param linearVelocity (vec3): The linear velocity vector to add.
   * @param angularVelocity (vec3): The angular velocity vector to add.
   */
  inline void AddBodyVelocity(int32_t entityHandle, plg::vec3 linearVelocity, plg::vec3 angularVelocity) {
    using AddBodyVelocityFn = void (*)(int32_t, plg::vec3, plg::vec3);
    static AddBodyVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddBodyVelocity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, linearVelocity, angularVelocity);
  }

  /**
   * @brief Detaches the entity from its parent.
   * @function DetachBodyFromParent
   * @param entityHandle (int32): The handle of the entity.
   */
  inline void DetachBodyFromParent(int32_t entityHandle) {
    using DetachBodyFromParentFn = void (*)(int32_t);
    static DetachBodyFromParentFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DetachBodyFromParent", reinterpret_cast<void**>(&__func));
    __func(entityHandle);
  }

  /**
   * @brief Retrieves the currently active sequence of the entity.
   * @function GetBodySequence
   * @param entityHandle (int32): The handle of the entity.
   * @return int32: The sequence ID of the active sequence, or -1 if invalid.
   */
  inline int32_t GetBodySequence(int32_t entityHandle) {
    using GetBodySequenceFn = int32_t (*)(int32_t);
    static GetBodySequenceFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetBodySequence", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Checks whether the entity is attached to a parent.
   * @function IsBodyAttachedToParent
   * @param entityHandle (int32): The handle of the entity.
   * @return bool: True if attached to a parent, false otherwise.
   */
  inline bool IsBodyAttachedToParent(int32_t entityHandle) {
    using IsBodyAttachedToParentFn = bool (*)(int32_t);
    static IsBodyAttachedToParentFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsBodyAttachedToParent", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Looks up a sequence ID by its name.
   * @function LookupBodySequence
   * @param entityHandle (int32): The handle of the entity.
   * @param name (string): The name of the sequence.
   * @return int32: The sequence ID, or -1 if not found.
   */
  inline int32_t LookupBodySequence(int32_t entityHandle, plg::string name) {
    using LookupBodySequenceFn = int32_t (*)(int32_t, plg::string);
    static LookupBodySequenceFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.LookupBodySequence", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, name);
  }

  /**
   * @brief Retrieves the duration of a specified sequence.
   * @function SetBodySequenceDuration
   * @param entityHandle (int32): The handle of the entity.
   * @param sequenceName (string): The name of the sequence.
   * @return float: The duration of the sequence in seconds, or 0 if invalid.
   */
  inline float SetBodySequenceDuration(int32_t entityHandle, plg::string sequenceName) {
    using SetBodySequenceDurationFn = float (*)(int32_t, plg::string);
    static SetBodySequenceDurationFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetBodySequenceDuration", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, sequenceName);
  }

  /**
   * @brief Sets the angular velocity of the entity.
   * @function SetBodyAngularVelocity
   * @param entityHandle (int32): The handle of the entity.
   * @param angVelocity (vec3): The new angular velocity vector.
   */
  inline void SetBodyAngularVelocity(int32_t entityHandle, plg::vec3 angVelocity) {
    using SetBodyAngularVelocityFn = void (*)(int32_t, plg::vec3);
    static SetBodyAngularVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetBodyAngularVelocity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, angVelocity);
  }

  /**
   * @brief Sets the material group of the entity.
   * @function SetBodyMaterialGroup
   * @param entityHandle (int32): The handle of the entity.
   * @param materialGroup (string): The material group token to assign.
   */
  inline void SetBodyMaterialGroup(int32_t entityHandle, plg::string materialGroup) {
    using SetBodyMaterialGroupFn = void (*)(int32_t, plg::string);
    static SetBodyMaterialGroupFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetBodyMaterialGroup", reinterpret_cast<void**>(&__func));
    __func(entityHandle, materialGroup);
  }

  /**
   * @brief Sets the linear velocity of the entity.
   * @function SetBodyVelocity
   * @param entityHandle (int32): The handle of the entity.
   * @param velocity (vec3): The new velocity vector.
   */
  inline void SetBodyVelocity(int32_t entityHandle, plg::vec3 velocity) {
    using SetBodyVelocityFn = void (*)(int32_t, plg::vec3);
    static SetBodyVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetBodyVelocity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, velocity);
  }

  /**
   * @brief Retrieves the player slot from a given entity pointer.
   * @function EntPointerToPlayerSlot
   * @param entity (ptr64): A pointer to the entity (CBaseEntity*).
   * @return int32: The player slot if valid, otherwise -1.
   */
  inline int32_t EntPointerToPlayerSlot(void* entity) {
    using EntPointerToPlayerSlotFn = int32_t (*)(void*);
    static EntPointerToPlayerSlotFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.EntPointerToPlayerSlot", reinterpret_cast<void**>(&__func));
    return __func(entity);
  }

  /**
   * @brief Returns a pointer to the entity instance by player slot index.
   * @function PlayerSlotToEntPointer
   * @param playerSlot (int32): Index of the player slot.
   * @return ptr64: Pointer to the entity instance, or nullptr if the slot is invalid.
   */
  inline void* PlayerSlotToEntPointer(int32_t playerSlot) {
    using PlayerSlotToEntPointerFn = void* (*)(int32_t);
    static PlayerSlotToEntPointerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PlayerSlotToEntPointer", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Returns the entity handle associated with a player slot index.
   * @function PlayerSlotToEntHandle
   * @param playerSlot (int32): Index of the player slot.
   * @return int32: The index of the entity, or -1 if the handle is invalid.
   */
  inline int32_t PlayerSlotToEntHandle(int32_t playerSlot) {
    using PlayerSlotToEntHandleFn = int32_t (*)(int32_t);
    static PlayerSlotToEntHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PlayerSlotToEntHandle", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the client object from a given player slot.
   * @function PlayerSlotToClientPtr
   * @param playerSlot (int32): The index of the player's slot (0-based).
   * @return ptr64: A pointer to the client object if found, otherwise nullptr.
   */
  inline void* PlayerSlotToClientPtr(int32_t playerSlot) {
    using PlayerSlotToClientPtrFn = void* (*)(int32_t);
    static PlayerSlotToClientPtrFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PlayerSlotToClientPtr", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the index of a given client object.
   * @function ClientPtrToPlayerSlot
   * @param client (ptr64): A pointer to the client object (CServerSideClient*).
   * @return int32: The player slot if found, otherwise -1.
   */
  inline int32_t ClientPtrToPlayerSlot(void* client) {
    using ClientPtrToPlayerSlotFn = int32_t (*)(void*);
    static ClientPtrToPlayerSlotFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ClientPtrToPlayerSlot", reinterpret_cast<void**>(&__func));
    return __func(client);
  }

  /**
   * @brief Returns the entity index for a given player slot.
   * @function PlayerSlotToClientIndex
   * @param playerSlot (int32): The index of the player's slot.
   * @return int32: The entity index if valid, otherwise 0.
   */
  inline int32_t PlayerSlotToClientIndex(int32_t playerSlot) {
    using PlayerSlotToClientIndexFn = int32_t (*)(int32_t);
    static PlayerSlotToClientIndexFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PlayerSlotToClientIndex", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the player slot from a given client index.
   * @function ClientIndexToPlayerSlot
   * @param clientIndex (int32): The index of the client.
   * @return int32: The player slot if valid, otherwise -1.
   */
  inline int32_t ClientIndexToPlayerSlot(int32_t clientIndex) {
    using ClientIndexToPlayerSlotFn = int32_t (*)(int32_t);
    static ClientIndexToPlayerSlotFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ClientIndexToPlayerSlot", reinterpret_cast<void**>(&__func));
    return __func(clientIndex);
  }

  /**
   * @brief Retrieves the player slot from a given player service.
   * @function PlayerServicesToPlayerSlot
   * @param service (ptr64): The service pointer. Like CCSPlayer_ItemServices, CCSPlayer_WeaponServices ect.
   * @return int32: The player slot if valid, otherwise -1.
   */
  inline int32_t PlayerServicesToPlayerSlot(void* service) {
    using PlayerServicesToPlayerSlotFn = int32_t (*)(void*);
    static PlayerServicesToPlayerSlotFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PlayerServicesToPlayerSlot", reinterpret_cast<void**>(&__func));
    return __func(service);
  }

  /**
   * @brief Retrieves a client's authentication string (SteamID).
   * @function GetClientAuthId
   * @param playerSlot (int32): The index of the player's slot whose authentication string is being retrieved.
   * @return string: The authentication string.
   */
  inline plg::string GetClientAuthId(int32_t playerSlot) {
    using GetClientAuthIdFn = plg::string (*)(int32_t);
    static GetClientAuthIdFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientAuthId", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Returns the client's Steam account ID, a unique number identifying a given Steam account.
   * @function GetClientAccountId
   * @param playerSlot (int32): The index of the player's slot.
   * @return uint32: uint32_t The client's steam account ID.
   */
  inline uint32_t GetClientAccountId(int32_t playerSlot) {
    using GetClientAccountIdFn = uint32_t (*)(int32_t);
    static GetClientAccountIdFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientAccountId", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Returns the client's SteamID64 â€” a unique 64-bit identifier of a Steam account.
   * @function GetClientSteamID64
   * @param playerSlot (int32): The index of the player's slot.
   * @return uint64: uint64_t The client's SteamID64.
   */
  inline uint64_t GetClientSteamID64(int32_t playerSlot) {
    using GetClientSteamID64Fn = uint64_t (*)(int32_t);
    static GetClientSteamID64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientSteamID64", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves a client's IP address.
   * @function GetClientIp
   * @param playerSlot (int32): The index of the player's slot.
   * @return string: The client's IP address.
   */
  inline plg::string GetClientIp(int32_t playerSlot) {
    using GetClientIpFn = plg::string (*)(int32_t);
    static GetClientIpFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientIp", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves a client's language.
   * @function GetClientLanguage
   * @param playerSlot (int32): The index of the player's slot.
   * @return string: The client's language.
   */
  inline plg::string GetClientLanguage(int32_t playerSlot) {
    using GetClientLanguageFn = plg::string (*)(int32_t);
    static GetClientLanguageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientLanguage", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves a client's operating system.
   * @function GetClientOS
   * @param playerSlot (int32): The index of the player's slot.
   * @return string: The client's operating system.
   */
  inline plg::string GetClientOS(int32_t playerSlot) {
    using GetClientOSFn = plg::string (*)(int32_t);
    static GetClientOSFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientOS", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Returns the client's name.
   * @function GetClientName
   * @param playerSlot (int32): The index of the player's slot.
   * @return string: The client's name.
   */
  inline plg::string GetClientName(int32_t playerSlot) {
    using GetClientNameFn = plg::string (*)(int32_t);
    static GetClientNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientName", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Returns the client's connection time in seconds.
   * @function GetClientTime
   * @param playerSlot (int32): The index of the player's slot.
   * @return float: float Connection time in seconds.
   */
  inline float GetClientTime(int32_t playerSlot) {
    using GetClientTimeFn = float (*)(int32_t);
    static GetClientTimeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientTime", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Returns the client's current latency (RTT).
   * @function GetClientLatency
   * @param playerSlot (int32): The index of the player's slot.
   * @return float: float Latency value.
   */
  inline float GetClientLatency(int32_t playerSlot) {
    using GetClientLatencyFn = float (*)(int32_t);
    static GetClientLatencyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientLatency", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Returns the client's access flags.
   * @function GetUserFlagBits
   * @param playerSlot (int32): The index of the player's slot.
   * @return uint64: uint64 Access flags as a bitmask.
   */
  inline uint64_t GetUserFlagBits(int32_t playerSlot) {
    using GetUserFlagBitsFn = uint64_t (*)(int32_t);
    static GetUserFlagBitsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetUserFlagBits", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the access flags on a client using a bitmask.
   * @function SetUserFlagBits
   * @param playerSlot (int32): The index of the player's slot.
   * @param flags (uint64): Bitmask representing the flags to be set.
   */
  inline void SetUserFlagBits(int32_t playerSlot, uint64_t flags) {
    using SetUserFlagBitsFn = void (*)(int32_t, uint64_t);
    static SetUserFlagBitsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetUserFlagBits", reinterpret_cast<void**>(&__func));
    __func(playerSlot, flags);
  }

  /**
   * @brief Adds access flags to a client.
   * @function AddUserFlags
   * @param playerSlot (int32): The index of the player's slot.
   * @param flags (uint64): Bitmask representing the flags to be added.
   */
  inline void AddUserFlags(int32_t playerSlot, uint64_t flags) {
    using AddUserFlagsFn = void (*)(int32_t, uint64_t);
    static AddUserFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddUserFlags", reinterpret_cast<void**>(&__func));
    __func(playerSlot, flags);
  }

  /**
   * @brief Removes access flags from a client.
   * @function RemoveUserFlags
   * @param playerSlot (int32): The index of the player's slot.
   * @param flags (uint64): Bitmask representing the flags to be removed.
   */
  inline void RemoveUserFlags(int32_t playerSlot, uint64_t flags) {
    using RemoveUserFlagsFn = void (*)(int32_t, uint64_t);
    static RemoveUserFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RemoveUserFlags", reinterpret_cast<void**>(&__func));
    __func(playerSlot, flags);
  }

  /**
   * @brief Checks if a certain player has been authenticated.
   * @function IsClientAuthorized
   * @param playerSlot (int32): The index of the player's slot.
   * @return bool: true if the player is authenticated, false otherwise.
   */
  inline bool IsClientAuthorized(int32_t playerSlot) {
    using IsClientAuthorizedFn = bool (*)(int32_t);
    static IsClientAuthorizedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsClientAuthorized", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Checks if a certain player is connected.
   * @function IsClientConnected
   * @param playerSlot (int32): The index of the player's slot.
   * @return bool: true if the player is connected, false otherwise.
   */
  inline bool IsClientConnected(int32_t playerSlot) {
    using IsClientConnectedFn = bool (*)(int32_t);
    static IsClientConnectedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsClientConnected", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Checks if a certain player has entered the game.
   * @function IsClientInGame
   * @param playerSlot (int32): The index of the player's slot.
   * @return bool: true if the player is in the game, false otherwise.
   */
  inline bool IsClientInGame(int32_t playerSlot) {
    using IsClientInGameFn = bool (*)(int32_t);
    static IsClientInGameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsClientInGame", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Checks if a certain player is the SourceTV bot.
   * @function IsClientSourceTV
   * @param playerSlot (int32): The index of the player's slot.
   * @return bool: true if the client is the SourceTV bot, false otherwise.
   */
  inline bool IsClientSourceTV(int32_t playerSlot) {
    using IsClientSourceTVFn = bool (*)(int32_t);
    static IsClientSourceTVFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsClientSourceTV", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Checks if the client is alive or dead.
   * @function IsClientAlive
   * @param playerSlot (int32): The index of the player's slot.
   * @return bool: true if the client is alive, false if dead.
   */
  inline bool IsClientAlive(int32_t playerSlot) {
    using IsClientAliveFn = bool (*)(int32_t);
    static IsClientAliveFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsClientAlive", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Checks if a certain player is a fake client.
   * @function IsFakeClient
   * @param playerSlot (int32): The index of the player's slot.
   * @return bool: true if the client is a fake client, false otherwise.
   */
  inline bool IsFakeClient(int32_t playerSlot) {
    using IsFakeClientFn = bool (*)(int32_t);
    static IsFakeClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsFakeClient", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the movement type of an client.
   * @function GetClientMoveType
   * @param playerSlot (int32): The index of the player's slot whose movement type is to be retrieved.
   * @return int32: The movement type of the entity, or 0 if the entity is invalid.
   */
  inline int32_t GetClientMoveType(int32_t playerSlot) {
    using GetClientMoveTypeFn = int32_t (*)(int32_t);
    static GetClientMoveTypeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientMoveType", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the movement type of an client.
   * @function SetClientMoveType
   * @param playerSlot (int32): The index of the player's slot whose movement type is to be set.
   * @param moveType (int32): The movement type of the entity, or 0 if the entity is invalid.
   */
  inline void SetClientMoveType(int32_t playerSlot, int32_t moveType) {
    using SetClientMoveTypeFn = void (*)(int32_t, int32_t);
    static SetClientMoveTypeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientMoveType", reinterpret_cast<void**>(&__func));
    __func(playerSlot, moveType);
  }

  /**
   * @brief Retrieves the gravity scale of an client.
   * @function GetClientGravity
   * @param playerSlot (int32): The index of the player's slot whose gravity scale is to be retrieved.
   * @return float: The gravity scale of the client, or 0.0f if the client is invalid.
   */
  inline float GetClientGravity(int32_t playerSlot) {
    using GetClientGravityFn = float (*)(int32_t);
    static GetClientGravityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientGravity", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the gravity scale of an client.
   * @function SetClientGravity
   * @param playerSlot (int32): The index of the player's slot whose gravity scale is to be set.
   * @param gravity (float): The new gravity scale to set for the client.
   */
  inline void SetClientGravity(int32_t playerSlot, float gravity) {
    using SetClientGravityFn = void (*)(int32_t, float);
    static SetClientGravityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientGravity", reinterpret_cast<void**>(&__func));
    __func(playerSlot, gravity);
  }

  /**
   * @brief Retrieves the flags of an client.
   * @function GetClientFlags
   * @param playerSlot (int32): The index of the player's slot whose flags are to be retrieved.
   * @return int32: The flags of the client, or 0 if the client is invalid.
   */
  inline int32_t GetClientFlags(int32_t playerSlot) {
    using GetClientFlagsFn = int32_t (*)(int32_t);
    static GetClientFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientFlags", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the flags of an client.
   * @function SetClientFlags
   * @param playerSlot (int32): The index of the player's slot whose flags are to be set.
   * @param flags (int32): The new flags to set for the client.
   */
  inline void SetClientFlags(int32_t playerSlot, int32_t flags) {
    using SetClientFlagsFn = void (*)(int32_t, int32_t);
    static SetClientFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientFlags", reinterpret_cast<void**>(&__func));
    __func(playerSlot, flags);
  }

  /**
   * @brief Retrieves the render color of an client.
   * @function GetClientRenderColor
   * @param playerSlot (int32): The index of the player's slot whose render color is to be retrieved.
   * @return int32: The raw color value of the client's render color, or 0 if the client is invalid.
   */
  inline int32_t GetClientRenderColor(int32_t playerSlot) {
    using GetClientRenderColorFn = int32_t (*)(int32_t);
    static GetClientRenderColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientRenderColor", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the render color of an client.
   * @function SetClientRenderColor
   * @param playerSlot (int32): The index of the player's slot whose render color is to be set.
   * @param color (int32): The new raw color value to set for the client's render color.
   */
  inline void SetClientRenderColor(int32_t playerSlot, int32_t color) {
    using SetClientRenderColorFn = void (*)(int32_t, int32_t);
    static SetClientRenderColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientRenderColor", reinterpret_cast<void**>(&__func));
    __func(playerSlot, color);
  }

  /**
   * @brief Retrieves the render mode of an client.
   * @function GetClientRenderMode
   * @param playerSlot (int32): The index of the player's slot whose render mode is to be retrieved.
   * @return uint8: The render mode of the client, or 0 if the client is invalid.
   */
  inline uint8_t GetClientRenderMode(int32_t playerSlot) {
    using GetClientRenderModeFn = uint8_t (*)(int32_t);
    static GetClientRenderModeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientRenderMode", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the render mode of an client.
   * @function SetClientRenderMode
   * @param playerSlot (int32): The index of the player's slot whose render mode is to be set.
   * @param renderMode (uint8): The new render mode to set for the client.
   */
  inline void SetClientRenderMode(int32_t playerSlot, uint8_t renderMode) {
    using SetClientRenderModeFn = void (*)(int32_t, uint8_t);
    static SetClientRenderModeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientRenderMode", reinterpret_cast<void**>(&__func));
    __func(playerSlot, renderMode);
  }

  /**
   * @brief Retrieves the mass of an client.
   * @function GetClientMass
   * @param playerSlot (int32): The index of the player's slot whose mass is to be retrieved.
   * @return int32: The mass of the client, or 0 if the client is invalid.
   */
  inline int32_t GetClientMass(int32_t playerSlot) {
    using GetClientMassFn = int32_t (*)(int32_t);
    static GetClientMassFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientMass", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the mass of an client.
   * @function SetClientMass
   * @param playerSlot (int32): The index of the player's slot whose mass is to be set.
   * @param mass (int32): The new mass value to set for the client.
   */
  inline void SetClientMass(int32_t playerSlot, int32_t mass) {
    using SetClientMassFn = void (*)(int32_t, int32_t);
    static SetClientMassFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientMass", reinterpret_cast<void**>(&__func));
    __func(playerSlot, mass);
  }

  /**
   * @brief Retrieves the friction of an client.
   * @function GetClientFriction
   * @param playerSlot (int32): The index of the player's slot whose friction is to be retrieved.
   * @return float: The friction of the client, or 0 if the client is invalid.
   */
  inline float GetClientFriction(int32_t playerSlot) {
    using GetClientFrictionFn = float (*)(int32_t);
    static GetClientFrictionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientFriction", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the friction of an client.
   * @function SetClientFriction
   * @param playerSlot (int32): The index of the player's slot whose friction is to be set.
   * @param friction (float): The new friction value to set for the client.
   */
  inline void SetClientFriction(int32_t playerSlot, float friction) {
    using SetClientFrictionFn = void (*)(int32_t, float);
    static SetClientFrictionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientFriction", reinterpret_cast<void**>(&__func));
    __func(playerSlot, friction);
  }

  /**
   * @brief Retrieves the health of an client.
   * @function GetClientHealth
   * @param playerSlot (int32): The index of the player's slot whose health is to be retrieved.
   * @return int32: The health of the client, or 0 if the client is invalid.
   */
  inline int32_t GetClientHealth(int32_t playerSlot) {
    using GetClientHealthFn = int32_t (*)(int32_t);
    static GetClientHealthFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientHealth", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the health of an client.
   * @function SetClientHealth
   * @param playerSlot (int32): The index of the player's slot whose health is to be set.
   * @param health (int32): The new health value to set for the client.
   */
  inline void SetClientHealth(int32_t playerSlot, int32_t health) {
    using SetClientHealthFn = void (*)(int32_t, int32_t);
    static SetClientHealthFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientHealth", reinterpret_cast<void**>(&__func));
    __func(playerSlot, health);
  }

  /**
   * @brief Retrieves the max health of an client.
   * @function GetClientMaxHealth
   * @param playerSlot (int32): The index of the player's slot whose max health is to be retrieved.
   * @return int32: The max health of the client, or 0 if the client is invalid.
   */
  inline int32_t GetClientMaxHealth(int32_t playerSlot) {
    using GetClientMaxHealthFn = int32_t (*)(int32_t);
    static GetClientMaxHealthFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientMaxHealth", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the max health of an client.
   * @function SetClientMaxHealth
   * @param playerSlot (int32): The index of the player's slot whose max health is to be set.
   * @param maxHealth (int32): The new max health value to set for the client.
   */
  inline void SetClientMaxHealth(int32_t playerSlot, int32_t maxHealth) {
    using SetClientMaxHealthFn = void (*)(int32_t, int32_t);
    static SetClientMaxHealthFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientMaxHealth", reinterpret_cast<void**>(&__func));
    __func(playerSlot, maxHealth);
  }

  /**
   * @brief Retrieves the team number of an client.
   * @function GetClientTeam
   * @param playerSlot (int32): The index of the player's slot whose team number is to be retrieved.
   * @return int32: The team number of the client, or 0 if the client is invalid.
   */
  inline int32_t GetClientTeam(int32_t playerSlot) {
    using GetClientTeamFn = int32_t (*)(int32_t);
    static GetClientTeamFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientTeam", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the team number of an client.
   * @function SetClientTeam
   * @param playerSlot (int32): The index of the player's slot whose team number is to be set.
   * @param team (int32): The new team number to set for the client.
   */
  inline void SetClientTeam(int32_t playerSlot, int32_t team) {
    using SetClientTeamFn = void (*)(int32_t, int32_t);
    static SetClientTeamFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientTeam", reinterpret_cast<void**>(&__func));
    __func(playerSlot, team);
  }

  /**
   * @brief Retrieves the absolute origin of an client.
   * @function GetClientAbsOrigin
   * @param playerSlot (int32): The index of the player's slot whose absolute origin is to be retrieved.
   * @return vec3: A vector where the absolute origin will be stored.
   */
  inline plg::vec3 GetClientAbsOrigin(int32_t playerSlot) {
    using GetClientAbsOriginFn = plg::vec3 (*)(int32_t);
    static GetClientAbsOriginFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientAbsOrigin", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the absolute origin of an client.
   * @function SetClientAbsOrigin
   * @param playerSlot (int32): The index of the player's slot whose absolute origin is to be set.
   * @param origin (vec3): The new absolute origin to set for the client.
   */
  inline void SetClientAbsOrigin(int32_t playerSlot, plg::vec3 origin) {
    using SetClientAbsOriginFn = void (*)(int32_t, plg::vec3);
    static SetClientAbsOriginFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientAbsOrigin", reinterpret_cast<void**>(&__func));
    __func(playerSlot, origin);
  }

  /**
   * @brief Retrieves the absolute scale of an client.
   * @function GetClientAbsScale
   * @param playerSlot (int32): The index of the player's slot whose absolute scale is to be retrieved.
   * @return float: A vector where the absolute scale will be stored.
   */
  inline float GetClientAbsScale(int32_t playerSlot) {
    using GetClientAbsScaleFn = float (*)(int32_t);
    static GetClientAbsScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientAbsScale", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the absolute scale of an client.
   * @function SetClientAbsScale
   * @param playerSlot (int32): The index of the player's slot whose absolute scale is to be set.
   * @param scale (float): The new absolute scale to set for the client.
   */
  inline void SetClientAbsScale(int32_t playerSlot, float scale) {
    using SetClientAbsScaleFn = void (*)(int32_t, float);
    static SetClientAbsScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientAbsScale", reinterpret_cast<void**>(&__func));
    __func(playerSlot, scale);
  }

  /**
   * @brief Retrieves the angular rotation of an client.
   * @function GetClientAbsAngles
   * @param playerSlot (int32): The index of the player's slot whose angular rotation is to be retrieved.
   * @return vec3: A QAngle where the angular rotation will be stored.
   */
  inline plg::vec3 GetClientAbsAngles(int32_t playerSlot) {
    using GetClientAbsAnglesFn = plg::vec3 (*)(int32_t);
    static GetClientAbsAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientAbsAngles", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the angular rotation of an client.
   * @function SetClientAbsAngles
   * @param playerSlot (int32): The index of the player's slot whose angular rotation is to be set.
   * @param angle (vec3): The new angular rotation to set for the client.
   */
  inline void SetClientAbsAngles(int32_t playerSlot, plg::vec3 angle) {
    using SetClientAbsAnglesFn = void (*)(int32_t, plg::vec3);
    static SetClientAbsAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientAbsAngles", reinterpret_cast<void**>(&__func));
    __func(playerSlot, angle);
  }

  /**
   * @brief Retrieves the local origin of an client.
   * @function GetClientLocalOrigin
   * @param playerSlot (int32): The index of the player's slot whose local origin is to be retrieved.
   * @return vec3: A vector where the local origin will be stored.
   */
  inline plg::vec3 GetClientLocalOrigin(int32_t playerSlot) {
    using GetClientLocalOriginFn = plg::vec3 (*)(int32_t);
    static GetClientLocalOriginFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientLocalOrigin", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the local origin of an client.
   * @function SetClientLocalOrigin
   * @param playerSlot (int32): The index of the player's slot whose local origin is to be set.
   * @param origin (vec3): The new local origin to set for the client.
   */
  inline void SetClientLocalOrigin(int32_t playerSlot, plg::vec3 origin) {
    using SetClientLocalOriginFn = void (*)(int32_t, plg::vec3);
    static SetClientLocalOriginFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientLocalOrigin", reinterpret_cast<void**>(&__func));
    __func(playerSlot, origin);
  }

  /**
   * @brief Retrieves the local scale of an client.
   * @function GetClientLocalScale
   * @param playerSlot (int32): The index of the player's slot whose local scale is to be retrieved.
   * @return float: A vector where the local scale will be stored.
   */
  inline float GetClientLocalScale(int32_t playerSlot) {
    using GetClientLocalScaleFn = float (*)(int32_t);
    static GetClientLocalScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientLocalScale", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the local scale of an client.
   * @function SetClientLocalScale
   * @param playerSlot (int32): The index of the player's slot whose local scale is to be set.
   * @param scale (float): The new local scale to set for the client.
   */
  inline void SetClientLocalScale(int32_t playerSlot, float scale) {
    using SetClientLocalScaleFn = void (*)(int32_t, float);
    static SetClientLocalScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientLocalScale", reinterpret_cast<void**>(&__func));
    __func(playerSlot, scale);
  }

  /**
   * @brief Retrieves the angular rotation of an client.
   * @function GetClientLocalAngles
   * @param playerSlot (int32): The index of the player's slot whose angular rotation is to be retrieved.
   * @return vec3: A QAngle where the angular rotation will be stored.
   */
  inline plg::vec3 GetClientLocalAngles(int32_t playerSlot) {
    using GetClientLocalAnglesFn = plg::vec3 (*)(int32_t);
    static GetClientLocalAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientLocalAngles", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the angular rotation of an client.
   * @function SetClientLocalAngles
   * @param playerSlot (int32): The index of the player's slot whose angular rotation is to be set.
   * @param angle (vec3): The new angular rotation to set for the client.
   */
  inline void SetClientLocalAngles(int32_t playerSlot, plg::vec3 angle) {
    using SetClientLocalAnglesFn = void (*)(int32_t, plg::vec3);
    static SetClientLocalAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientLocalAngles", reinterpret_cast<void**>(&__func));
    __func(playerSlot, angle);
  }

  /**
   * @brief Retrieves the absolute velocity of an client.
   * @function GetClientAbsVelocity
   * @param playerSlot (int32): The index of the player's slot whose absolute velocity is to be retrieved.
   * @return vec3: A vector where the absolute velocity will be stored.
   */
  inline plg::vec3 GetClientAbsVelocity(int32_t playerSlot) {
    using GetClientAbsVelocityFn = plg::vec3 (*)(int32_t);
    static GetClientAbsVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientAbsVelocity", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the absolute velocity of an client.
   * @function SetClientAbsVelocity
   * @param playerSlot (int32): The index of the player's slot whose absolute velocity is to be set.
   * @param velocity (vec3): The new absolute velocity to set for the client.
   */
  inline void SetClientAbsVelocity(int32_t playerSlot, plg::vec3 velocity) {
    using SetClientAbsVelocityFn = void (*)(int32_t, plg::vec3);
    static SetClientAbsVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientAbsVelocity", reinterpret_cast<void**>(&__func));
    __func(playerSlot, velocity);
  }

  /**
   * @brief Retrieves the base velocity of an client.
   * @function GetClientBaseVelocity
   * @param playerSlot (int32): The index of the player's slot whose base velocity is to be retrieved.
   * @return vec3: A vector where the base velocity will be stored.
   */
  inline plg::vec3 GetClientBaseVelocity(int32_t playerSlot) {
    using GetClientBaseVelocityFn = plg::vec3 (*)(int32_t);
    static GetClientBaseVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientBaseVelocity", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the local angular velocity of an client.
   * @function GetClientLocalAngVelocity
   * @param playerSlot (int32): The index of the player's slot whose local angular velocity is to be retrieved.
   * @return vec3: A vector where the local angular velocity will be stored.
   */
  inline plg::vec3 GetClientLocalAngVelocity(int32_t playerSlot) {
    using GetClientLocalAngVelocityFn = plg::vec3 (*)(int32_t);
    static GetClientLocalAngVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientLocalAngVelocity", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the angular velocity of an client.
   * @function GetClientAngVelocity
   * @param playerSlot (int32): The index of the player's slot whose angular velocity is to be retrieved.
   * @return vec3: A vector where the angular velocity will be stored.
   */
  inline plg::vec3 GetClientAngVelocity(int32_t playerSlot) {
    using GetClientAngVelocityFn = plg::vec3 (*)(int32_t);
    static GetClientAngVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientAngVelocity", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the angular velocity of an client.
   * @function SetClientAngVelocity
   * @param playerSlot (int32): The index of the player's slot whose angular velocity is to be set.
   * @param velocity (vec3): The new angular velocity to set for the client.
   */
  inline void SetClientAngVelocity(int32_t playerSlot, plg::vec3 velocity) {
    using SetClientAngVelocityFn = void (*)(int32_t, plg::vec3);
    static SetClientAngVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientAngVelocity", reinterpret_cast<void**>(&__func));
    __func(playerSlot, velocity);
  }

  /**
   * @brief Retrieves the local velocity of an client.
   * @function GetClientLocalVelocity
   * @param playerSlot (int32): The index of the player's slot whose local velocity is to be retrieved.
   * @return vec3: A vector where the local velocity will be stored.
   */
  inline plg::vec3 GetClientLocalVelocity(int32_t playerSlot) {
    using GetClientLocalVelocityFn = plg::vec3 (*)(int32_t);
    static GetClientLocalVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientLocalVelocity", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the angular rotation of an client.
   * @function GetClientAngRotation
   * @param playerSlot (int32): The index of the player's slot whose angular rotation is to be retrieved.
   * @return vec3: A vector where the angular rotation will be stored.
   */
  inline plg::vec3 GetClientAngRotation(int32_t playerSlot) {
    using GetClientAngRotationFn = plg::vec3 (*)(int32_t);
    static GetClientAngRotationFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientAngRotation", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the angular rotation of an client.
   * @function SetClientAngRotation
   * @param playerSlot (int32): The index of the player's slot whose angular rotation is to be set.
   * @param rotation (vec3): The new angular rotation to set for the client.
   */
  inline void SetClientAngRotation(int32_t playerSlot, plg::vec3 rotation) {
    using SetClientAngRotationFn = void (*)(int32_t, plg::vec3);
    static SetClientAngRotationFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientAngRotation", reinterpret_cast<void**>(&__func));
    __func(playerSlot, rotation);
  }

  /**
   * @brief Returns the input Vector transformed from client to world space.
   * @function TransformPointClientToWorld
   * @param playerSlot (int32): The index of the player's slot
   * @param point (vec3): Point in client local space to transform
   * @return vec3: The point transformed to world space coordinates
   */
  inline plg::vec3 TransformPointClientToWorld(int32_t playerSlot, plg::vec3 point) {
    using TransformPointClientToWorldFn = plg::vec3 (*)(int32_t, plg::vec3);
    static TransformPointClientToWorldFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TransformPointClientToWorld", reinterpret_cast<void**>(&__func));
    return __func(playerSlot, point);
  }

  /**
   * @brief Returns the input Vector transformed from world to client space.
   * @function TransformPointWorldToClient
   * @param playerSlot (int32): The index of the player's slot
   * @param point (vec3): Point in world space to transform
   * @return vec3: The point transformed to client local space coordinates
   */
  inline plg::vec3 TransformPointWorldToClient(int32_t playerSlot, plg::vec3 point) {
    using TransformPointWorldToClientFn = plg::vec3 (*)(int32_t, plg::vec3);
    static TransformPointWorldToClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TransformPointWorldToClient", reinterpret_cast<void**>(&__func));
    return __func(playerSlot, point);
  }

  /**
   * @brief Get vector to eye position - absolute coords.
   * @function GetClientEyePosition
   * @param playerSlot (int32): The index of the player's slot
   * @return vec3: Eye position in absolute/world coordinates
   */
  inline plg::vec3 GetClientEyePosition(int32_t playerSlot) {
    using GetClientEyePositionFn = plg::vec3 (*)(int32_t);
    static GetClientEyePositionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientEyePosition", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Get the qangles that this client is looking at.
   * @function GetClientEyeAngles
   * @param playerSlot (int32): The index of the player's slot
   * @return vec3: Eye angles as a vector (pitch, yaw, roll)
   */
  inline plg::vec3 GetClientEyeAngles(int32_t playerSlot) {
    using GetClientEyeAnglesFn = plg::vec3 (*)(int32_t);
    static GetClientEyeAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientEyeAngles", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the forward velocity of an client.
   * @function SetClientForwardVector
   * @param playerSlot (int32): The index of the player's slot whose forward velocity is to be set.
   * @param forward (vec3)
   */
  inline void SetClientForwardVector(int32_t playerSlot, plg::vec3 forward) {
    using SetClientForwardVectorFn = void (*)(int32_t, plg::vec3);
    static SetClientForwardVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientForwardVector", reinterpret_cast<void**>(&__func));
    __func(playerSlot, forward);
  }

  /**
   * @brief Get the forward vector of the client.
   * @function GetClientForwardVector
   * @param playerSlot (int32): The index of the player's slot to query
   * @return vec3: Forward-facing direction vector of the client
   */
  inline plg::vec3 GetClientForwardVector(int32_t playerSlot) {
    using GetClientForwardVectorFn = plg::vec3 (*)(int32_t);
    static GetClientForwardVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientForwardVector", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Get the left vector of the client.
   * @function GetClientLeftVector
   * @param playerSlot (int32): The index of the player's slot to query
   * @return vec3: Left-facing direction vector of the client (aligned with the y axis)
   */
  inline plg::vec3 GetClientLeftVector(int32_t playerSlot) {
    using GetClientLeftVectorFn = plg::vec3 (*)(int32_t);
    static GetClientLeftVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientLeftVector", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Get the right vector of the client.
   * @function GetClientRightVector
   * @param playerSlot (int32): The index of the player's slot to query
   * @return vec3: Right-facing direction vector of the client
   */
  inline plg::vec3 GetClientRightVector(int32_t playerSlot) {
    using GetClientRightVectorFn = plg::vec3 (*)(int32_t);
    static GetClientRightVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientRightVector", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Get the up vector of the client.
   * @function GetClientUpVector
   * @param playerSlot (int32): The index of the player's slot to query
   * @return vec3: Up-facing direction vector of the client
   */
  inline plg::vec3 GetClientUpVector(int32_t playerSlot) {
    using GetClientUpVectorFn = plg::vec3 (*)(int32_t);
    static GetClientUpVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientUpVector", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Get the client-to-world transformation matrix.
   * @function GetClientTransform
   * @param playerSlot (int32): The index of the player's slot to query
   * @return mat4x4: 4x4 transformation matrix representing client's position, rotation, and scale in world space
   */
  inline plg::mat4x4 GetClientTransform(int32_t playerSlot) {
    using GetClientTransformFn = plg::mat4x4 (*)(int32_t);
    static GetClientTransformFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientTransform", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the model name of an client.
   * @function GetClientModel
   * @param playerSlot (int32): The index of the player's slot whose model name is to be retrieved.
   * @return string: A string where the model name will be stored.
   */
  inline plg::string GetClientModel(int32_t playerSlot) {
    using GetClientModelFn = plg::string (*)(int32_t);
    static GetClientModelFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientModel", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the model name of an client.
   * @function SetClientModel
   * @param playerSlot (int32): The index of the player's slot whose model name is to be set.
   * @param model (string): The new model name to set for the client.
   */
  inline void SetClientModel(int32_t playerSlot, plg::string model) {
    using SetClientModelFn = void (*)(int32_t, plg::string);
    static SetClientModelFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientModel", reinterpret_cast<void**>(&__func));
    __func(playerSlot, model);
  }

  /**
   * @brief Retrieves the water level of an client.
   * @function GetClientWaterLevel
   * @param playerSlot (int32): The index of the player's slot whose water level is to be retrieved.
   * @return float: The water level of the client, or 0.0f if the client is invalid.
   */
  inline float GetClientWaterLevel(int32_t playerSlot) {
    using GetClientWaterLevelFn = float (*)(int32_t);
    static GetClientWaterLevelFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientWaterLevel", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the ground client of an client.
   * @function GetClientGroundEntity
   * @param playerSlot (int32): The index of the player's slot whose ground client is to be retrieved.
   * @return int32: The handle of the ground client, or INVALID_EHANDLE_INDEX if the client is invalid.
   */
  inline int32_t GetClientGroundEntity(int32_t playerSlot) {
    using GetClientGroundEntityFn = int32_t (*)(int32_t);
    static GetClientGroundEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientGroundEntity", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves the effects of an client.
   * @function GetClientEffects
   * @param playerSlot (int32): The index of the player's slot whose effects are to be retrieved.
   * @return int32: The effect flags of the client, or 0 if the client is invalid.
   */
  inline int32_t GetClientEffects(int32_t playerSlot) {
    using GetClientEffectsFn = int32_t (*)(int32_t);
    static GetClientEffectsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientEffects", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Adds the render effect flag to an client.
   * @function AddClientEffects
   * @param playerSlot (int32): The index of the player's slot to modify
   * @param effects (int32): Render effect flags to add
   */
  inline void AddClientEffects(int32_t playerSlot, int32_t effects) {
    using AddClientEffectsFn = void (*)(int32_t, int32_t);
    static AddClientEffectsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddClientEffects", reinterpret_cast<void**>(&__func));
    __func(playerSlot, effects);
  }

  /**
   * @brief Removes the render effect flag from an client.
   * @function RemoveClientEffects
   * @param playerSlot (int32): The index of the player's slot to modify
   * @param effects (int32): Render effect flags to remove
   */
  inline void RemoveClientEffects(int32_t playerSlot, int32_t effects) {
    using RemoveClientEffectsFn = void (*)(int32_t, int32_t);
    static RemoveClientEffectsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RemoveClientEffects", reinterpret_cast<void**>(&__func));
    __func(playerSlot, effects);
  }

  /**
   * @brief Get a vector containing max bounds, centered on object.
   * @function GetClientBoundingMaxs
   * @param playerSlot (int32): The index of the player's slot to query
   * @return vec3: Vector containing the maximum bounds of the client's bounding box
   */
  inline plg::vec3 GetClientBoundingMaxs(int32_t playerSlot) {
    using GetClientBoundingMaxsFn = plg::vec3 (*)(int32_t);
    static GetClientBoundingMaxsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientBoundingMaxs", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Get a vector containing min bounds, centered on object.
   * @function GetClientBoundingMins
   * @param playerSlot (int32): The index of the player's slot to query
   * @return vec3: Vector containing the minimum bounds of the client's bounding box
   */
  inline plg::vec3 GetClientBoundingMins(int32_t playerSlot) {
    using GetClientBoundingMinsFn = plg::vec3 (*)(int32_t);
    static GetClientBoundingMinsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientBoundingMins", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Get vector to center of object - absolute coords.
   * @function GetClientCenter
   * @param playerSlot (int32): The index of the player's slot to query
   * @return vec3: Vector pointing to the center of the client in absolute/world coordinates
   */
  inline plg::vec3 GetClientCenter(int32_t playerSlot) {
    using GetClientCenterFn = plg::vec3 (*)(int32_t);
    static GetClientCenterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientCenter", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Teleports an client to a specified location and orientation.
   * @function TeleportClient
   * @param playerSlot (int32): The index of the player's slot to teleport.
   * @param origin (ptr64): A pointer to a Vector representing the new absolute position. Can be nullptr.
   * @param angles (ptr64): A pointer to a QAngle representing the new orientation. Can be nullptr.
   * @param velocity (ptr64): A pointer to a Vector representing the new velocity. Can be nullptr.
   */
  inline void TeleportClient(int32_t playerSlot, void* origin, void* angles, void* velocity) {
    using TeleportClientFn = void (*)(int32_t, void*, void*, void*);
    static TeleportClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TeleportClient", reinterpret_cast<void**>(&__func));
    __func(playerSlot, origin, angles, velocity);
  }

  /**
   * @brief Apply an absolute velocity impulse to an client.
   * @function ApplyAbsVelocityImpulseToClient
   * @param playerSlot (int32): The index of the player's slot to apply impulse to
   * @param vecImpulse (vec3): Velocity impulse vector to apply
   */
  inline void ApplyAbsVelocityImpulseToClient(int32_t playerSlot, plg::vec3 vecImpulse) {
    using ApplyAbsVelocityImpulseToClientFn = void (*)(int32_t, plg::vec3);
    static ApplyAbsVelocityImpulseToClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ApplyAbsVelocityImpulseToClient", reinterpret_cast<void**>(&__func));
    __func(playerSlot, vecImpulse);
  }

  /**
   * @brief Apply a local angular velocity impulse to an client.
   * @function ApplyLocalAngularVelocityImpulseToClient
   * @param playerSlot (int32): The index of the player's slot to apply impulse to
   * @param angImpulse (vec3): Angular velocity impulse vector to apply
   */
  inline void ApplyLocalAngularVelocityImpulseToClient(int32_t playerSlot, plg::vec3 angImpulse) {
    using ApplyLocalAngularVelocityImpulseToClientFn = void (*)(int32_t, plg::vec3);
    static ApplyLocalAngularVelocityImpulseToClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ApplyLocalAngularVelocityImpulseToClient", reinterpret_cast<void**>(&__func));
    __func(playerSlot, angImpulse);
  }

  /**
   * @brief Invokes a named input method on a specified client.
   * @function AcceptClientInput
   * @param playerSlot (int32): The handle of the target client that will receive the input.
   * @param inputName (string): The name of the input action to invoke.
   * @param activatorHandle (int32): The index of the player's slot that initiated the sequence of actions.
   * @param callerHandle (int32): The index of the player's slot sending this event. Use -1 to specify
   * @param value (any): The value associated with the input action.
   * @param type (int32): The type or classification of the value.
   * @param outputId (int32): An identifier for tracking the output of this operation.
   */
  inline void AcceptClientInput(int32_t playerSlot, plg::string inputName, int32_t activatorHandle, int32_t callerHandle, plg::any value, int32_t type, int32_t outputId) {
    using AcceptClientInputFn = void (*)(int32_t, plg::string, int32_t, int32_t, plg::any, int32_t, int32_t);
    static AcceptClientInputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AcceptClientInput", reinterpret_cast<void**>(&__func));
    __func(playerSlot, inputName, activatorHandle, callerHandle, value, type, outputId);
  }

  /**
   * @brief Connects a script function to an player output.
   * @function ConnectClientOutput
   * @param playerSlot (int32): The handle of the player.
   * @param output (string): The name of the output to connect to.
   * @param functionName (string): The name of the script function to call.
   */
  inline void ConnectClientOutput(int32_t playerSlot, plg::string output, plg::string functionName) {
    using ConnectClientOutputFn = void (*)(int32_t, plg::string, plg::string);
    static ConnectClientOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ConnectClientOutput", reinterpret_cast<void**>(&__func));
    __func(playerSlot, output, functionName);
  }

  /**
   * @brief Disconnects a script function from an player output.
   * @function DisconnectClientOutput
   * @param playerSlot (int32): The handle of the player.
   * @param output (string): The name of the output.
   * @param functionName (string): The name of the script function to disconnect.
   */
  inline void DisconnectClientOutput(int32_t playerSlot, plg::string output, plg::string functionName) {
    using DisconnectClientOutputFn = void (*)(int32_t, plg::string, plg::string);
    static DisconnectClientOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DisconnectClientOutput", reinterpret_cast<void**>(&__func));
    __func(playerSlot, output, functionName);
  }

  /**
   * @brief Disconnects a script function from an I/O event on a different player.
   * @function DisconnectClientRedirectedOutput
   * @param playerSlot (int32): The handle of the calling player.
   * @param output (string): The name of the output.
   * @param functionName (string): The function name to disconnect.
   * @param targetHandle (int32): The handle of the entity whose output is being disconnected.
   */
  inline void DisconnectClientRedirectedOutput(int32_t playerSlot, plg::string output, plg::string functionName, int32_t targetHandle) {
    using DisconnectClientRedirectedOutputFn = void (*)(int32_t, plg::string, plg::string, int32_t);
    static DisconnectClientRedirectedOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DisconnectClientRedirectedOutput", reinterpret_cast<void**>(&__func));
    __func(playerSlot, output, functionName, targetHandle);
  }

  /**
   * @brief Fires an player output.
   * @function FireClientOutput
   * @param playerSlot (int32): The handle of the player firing the output.
   * @param outputName (string): The name of the output to fire.
   * @param activatorHandle (int32): The entity activating the output.
   * @param callerHandle (int32): The entity that called the output.
   * @param value (any): The value associated with the input action.
   * @param type (int32): The type or classification of the value.
   * @param delay (float): Delay in seconds before firing the output.
   */
  inline void FireClientOutput(int32_t playerSlot, plg::string outputName, int32_t activatorHandle, int32_t callerHandle, plg::any value, int32_t type, float delay) {
    using FireClientOutputFn = void (*)(int32_t, plg::string, int32_t, int32_t, plg::any, int32_t, float);
    static FireClientOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FireClientOutput", reinterpret_cast<void**>(&__func));
    __func(playerSlot, outputName, activatorHandle, callerHandle, value, type, delay);
  }

  /**
   * @brief Redirects an player output to call a function on another player.
   * @function RedirectClientOutput
   * @param playerSlot (int32): The handle of the player whose output is being redirected.
   * @param output (string): The name of the output to redirect.
   * @param functionName (string): The function name to call on the target player.
   * @param targetHandle (int32): The handle of the entity that will receive the output call.
   */
  inline void RedirectClientOutput(int32_t playerSlot, plg::string output, plg::string functionName, int32_t targetHandle) {
    using RedirectClientOutputFn = void (*)(int32_t, plg::string, plg::string, int32_t);
    static RedirectClientOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RedirectClientOutput", reinterpret_cast<void**>(&__func));
    __func(playerSlot, output, functionName, targetHandle);
  }

  /**
   * @brief Makes an client follow another client with optional bone merging.
   * @function FollowClient
   * @param playerSlot (int32): The index of the player's slot that will follow
   * @param attachmentHandle (int32): The index of the player's slot to follow
   * @param boneMerge (bool): If true, bones will be merged between entities
   */
  inline void FollowClient(int32_t playerSlot, int32_t attachmentHandle, bool boneMerge) {
    using FollowClientFn = void (*)(int32_t, int32_t, bool);
    static FollowClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FollowClient", reinterpret_cast<void**>(&__func));
    __func(playerSlot, attachmentHandle, boneMerge);
  }

  /**
   * @brief Makes an client follow another client and merge with a specific bone or attachment.
   * @function FollowClientMerge
   * @param playerSlot (int32): The index of the player's slot that will follow
   * @param attachmentHandle (int32): The index of the player's slot to follow
   * @param boneOrAttachName (string): Name of the bone or attachment point to merge with
   */
  inline void FollowClientMerge(int32_t playerSlot, int32_t attachmentHandle, plg::string boneOrAttachName) {
    using FollowClientMergeFn = void (*)(int32_t, int32_t, plg::string);
    static FollowClientMergeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FollowClientMerge", reinterpret_cast<void**>(&__func));
    __func(playerSlot, attachmentHandle, boneOrAttachName);
  }

  /**
   * @brief Apply damage to an client.
   * @function TakeClientDamage
   * @param playerSlot (int32): The index of the player's slot receiving damage
   * @param inflictorSlot (int32): The index of the player's slot inflicting damage (e.g., projectile)
   * @param attackerSlot (int32): The index of the attacking client
   * @param force (vec3): Direction and magnitude of force to apply
   * @param hitPos (vec3): Position where the damage hit occurred
   * @param damage (float): Amount of damage to apply
   * @param damageTypes (int32): Bitfield of damage type flags
   * @return int32: Amount of damage actually applied to the client
   */
  inline int32_t TakeClientDamage(int32_t playerSlot, int32_t inflictorSlot, int32_t attackerSlot, plg::vec3 force, plg::vec3 hitPos, float damage, int32_t damageTypes) {
    using TakeClientDamageFn = int32_t (*)(int32_t, int32_t, int32_t, plg::vec3, plg::vec3, float, int32_t);
    static TakeClientDamageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TakeClientDamage", reinterpret_cast<void**>(&__func));
    return __func(playerSlot, inflictorSlot, attackerSlot, force, hitPos, damage, damageTypes);
  }

  /**
   * @brief Retrieves the pawn entity pointer associated with a client.
   * @function GetClientPawn
   * @param playerSlot (int32): The index of the player's slot.
   * @return ptr64: A pointer to the client's pawn entity, or nullptr if the client or controller is invalid.
   */
  inline void* GetClientPawn(int32_t playerSlot) {
    using GetClientPawnFn = void* (*)(int32_t);
    static GetClientPawnFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientPawn", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Processes the target string to determine if one user can target another.
   * @function ProcessTargetString
   * @param caller (int32): The index of the player's slot making the target request.
   * @param target (string): The target string specifying the player or players to be targeted.
   * @return int32[]: A vector where the result of the targeting operation will be stored.
   */
  inline plg::vector<int32_t> ProcessTargetString(int32_t caller, plg::string target) {
    using ProcessTargetStringFn = plg::vector<int32_t> (*)(int32_t, plg::string);
    static ProcessTargetStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ProcessTargetString", reinterpret_cast<void**>(&__func));
    return __func(caller, target);
  }

  /**
   * @brief Switches the player's team.
   * @function SwitchClientTeam
   * @param playerSlot (int32): The index of the player's slot.
   * @param team (int32): The team index to switch the client to.
   */
  inline void SwitchClientTeam(int32_t playerSlot, int32_t team) {
    using SwitchClientTeamFn = void (*)(int32_t, int32_t);
    static SwitchClientTeamFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SwitchClientTeam", reinterpret_cast<void**>(&__func));
    __func(playerSlot, team);
  }

  /**
   * @brief Respawns a player.
   * @function RespawnClient
   * @param playerSlot (int32): The index of the player's slot to respawn.
   */
  inline void RespawnClient(int32_t playerSlot) {
    using RespawnClientFn = void (*)(int32_t);
    static RespawnClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RespawnClient", reinterpret_cast<void**>(&__func));
    __func(playerSlot);
  }

  /**
   * @brief Forces a player to commit suicide.
   * @function ForcePlayerSuicide
   * @param playerSlot (int32): The index of the player's slot.
   * @param explode (bool): If true, the client will explode upon death.
   * @param force (bool): If true, the suicide will be forced.
   */
  inline void ForcePlayerSuicide(int32_t playerSlot, bool explode, bool force) {
    using ForcePlayerSuicideFn = void (*)(int32_t, bool, bool);
    static ForcePlayerSuicideFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ForcePlayerSuicide", reinterpret_cast<void**>(&__func));
    __func(playerSlot, explode, force);
  }

  /**
   * @brief Disconnects a client from the server as soon as the next frame starts.
   * @function KickClient
   * @param playerSlot (int32): The index of the player's slot to be kicked.
   */
  inline void KickClient(int32_t playerSlot) {
    using KickClientFn = void (*)(int32_t);
    static KickClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.KickClient", reinterpret_cast<void**>(&__func));
    __func(playerSlot);
  }

  /**
   * @brief Bans a client for a specified duration.
   * @function BanClient
   * @param playerSlot (int32): The index of the player's slot to be banned.
   * @param duration (float): Duration of the ban in seconds.
   * @param kick (bool): If true, the client will be kicked immediately after being banned.
   */
  inline void BanClient(int32_t playerSlot, float duration, bool kick) {
    using BanClientFn = void (*)(int32_t, float, bool);
    static BanClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.BanClient", reinterpret_cast<void**>(&__func));
    __func(playerSlot, duration, kick);
  }

  /**
   * @brief Bans an identity (either an IP address or a Steam authentication string).
   * @function BanIdentity
   * @param steamId (uint64): The Steam ID to ban.
   * @param duration (float): Duration of the ban in seconds.
   * @param kick (bool): If true, the client will be kicked immediately after being banned.
   */
  inline void BanIdentity(uint64_t steamId, float duration, bool kick) {
    using BanIdentityFn = void (*)(uint64_t, float, bool);
    static BanIdentityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.BanIdentity", reinterpret_cast<void**>(&__func));
    __func(steamId, duration, kick);
  }

  /**
   * @brief Retrieves the handle of the client's currently active weapon.
   * @function GetClientActiveWeapon
   * @param playerSlot (int32): The index of the player's slot.
   * @return int32: The entity handle of the active weapon, or INVALID_EHANDLE_INDEX if the client is invalid or has no active weapon.
   */
  inline int32_t GetClientActiveWeapon(int32_t playerSlot) {
    using GetClientActiveWeaponFn = int32_t (*)(int32_t);
    static GetClientActiveWeaponFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientActiveWeapon", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Retrieves a list of weapon handles owned by the client.
   * @function GetClientWeapons
   * @param playerSlot (int32): The index of the player's slot.
   * @return int32[]: A vector of entity handles for the client's weapons, or an empty vector if the client is invalid or has no weapons.
   */
  inline plg::vector<int32_t> GetClientWeapons(int32_t playerSlot) {
    using GetClientWeaponsFn = plg::vector<int32_t> (*)(int32_t);
    static GetClientWeaponsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientWeapons", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Removes all weapons from a client, with an option to remove the suit as well.
   * @function RemoveWeapons
   * @param playerSlot (int32): The index of the player's slot.
   * @param removeSuit (bool): A boolean indicating whether to also remove the client's suit.
   */
  inline void RemoveWeapons(int32_t playerSlot, bool removeSuit) {
    using RemoveWeaponsFn = void (*)(int32_t, bool);
    static RemoveWeaponsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RemoveWeapons", reinterpret_cast<void**>(&__func));
    __func(playerSlot, removeSuit);
  }

  /**
   * @brief Forces a player to drop their weapon.
   * @function DropWeapon
   * @param playerSlot (int32): The index of the player's slot.
   * @param weaponHandle (int32): The handle of weapon to drop.
   * @param target (vec3): Target direction.
   * @param velocity (vec3): Velocity to toss weapon or zero to just drop weapon.
   */
  inline void DropWeapon(int32_t playerSlot, int32_t weaponHandle, plg::vec3 target, plg::vec3 velocity) {
    using DropWeaponFn = void (*)(int32_t, int32_t, plg::vec3, plg::vec3);
    static DropWeaponFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DropWeapon", reinterpret_cast<void**>(&__func));
    __func(playerSlot, weaponHandle, target, velocity);
  }

  /**
   * @brief Selects a player's weapon.
   * @function SelectWeapon
   * @param playerSlot (int32): The index of the player's slot.
   * @param weaponHandle (int32): The handle of weapon to bump.
   */
  inline void SelectWeapon(int32_t playerSlot, int32_t weaponHandle) {
    using SelectWeaponFn = void (*)(int32_t, int32_t);
    static SelectWeaponFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SelectWeapon", reinterpret_cast<void**>(&__func));
    __func(playerSlot, weaponHandle);
  }

  /**
   * @brief Switches a player's weapon.
   * @function SwitchWeapon
   * @param playerSlot (int32): The index of the player's slot.
   * @param weaponHandle (int32): The handle of weapon to switch.
   */
  inline void SwitchWeapon(int32_t playerSlot, int32_t weaponHandle) {
    using SwitchWeaponFn = void (*)(int32_t, int32_t);
    static SwitchWeaponFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SwitchWeapon", reinterpret_cast<void**>(&__func));
    __func(playerSlot, weaponHandle);
  }

  /**
   * @brief Removes a player's weapon.
   * @function RemoveWeapon
   * @param playerSlot (int32): The index of the player's slot.
   * @param weaponHandle (int32): The handle of weapon to remove.
   */
  inline void RemoveWeapon(int32_t playerSlot, int32_t weaponHandle) {
    using RemoveWeaponFn = void (*)(int32_t, int32_t);
    static RemoveWeaponFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RemoveWeapon", reinterpret_cast<void**>(&__func));
    __func(playerSlot, weaponHandle);
  }

  /**
   * @brief Gives a named item (e.g., weapon) to a client.
   * @function GiveNamedItem
   * @param playerSlot (int32): The index of the player's slot.
   * @param itemName (string): The name of the item to give.
   * @return int32: The entity handle of the created item, or INVALID_EHANDLE_INDEX if the client or item is invalid.
   */
  inline int32_t GiveNamedItem(int32_t playerSlot, plg::string itemName) {
    using GiveNamedItemFn = int32_t (*)(int32_t, plg::string);
    static GiveNamedItemFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GiveNamedItem", reinterpret_cast<void**>(&__func));
    return __func(playerSlot, itemName);
  }

  /**
   * @brief Retrieves the state of a specific button for a client.
   * @function GetClientButtons
   * @param playerSlot (int32): The index of the player's slot.
   * @param buttonIndex (int32): The index of the button (0-2).
   * @return uint64: uint64_t The state of the specified button, or 0 if the client or button index is invalid.
   */
  inline uint64_t GetClientButtons(int32_t playerSlot, int32_t buttonIndex) {
    using GetClientButtonsFn = uint64_t (*)(int32_t, int32_t);
    static GetClientButtonsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientButtons", reinterpret_cast<void**>(&__func));
    return __func(playerSlot, buttonIndex);
  }

  /**
   * @brief Returns the client's armor value.
   * @function GetClientArmor
   * @param playerSlot (int32): The index of the player's slot.
   * @return int32: The armor value of the client.
   */
  inline int32_t GetClientArmor(int32_t playerSlot) {
    using GetClientArmorFn = int32_t (*)(int32_t);
    static GetClientArmorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientArmor", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the client's armor value.
   * @function SetClientArmor
   * @param playerSlot (int32): The index of the player's slot.
   * @param armor (int32): The armor value to set.
   */
  inline void SetClientArmor(int32_t playerSlot, int32_t armor) {
    using SetClientArmorFn = void (*)(int32_t, int32_t);
    static SetClientArmorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientArmor", reinterpret_cast<void**>(&__func));
    __func(playerSlot, armor);
  }

  /**
   * @brief Returns the client's speed value.
   * @function GetClientSpeed
   * @param playerSlot (int32): The index of the player's slot.
   * @return float: The speed value of the client.
   */
  inline float GetClientSpeed(int32_t playerSlot) {
    using GetClientSpeedFn = float (*)(int32_t);
    static GetClientSpeedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientSpeed", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the client's speed value.
   * @function SetClientSpeed
   * @param playerSlot (int32): The index of the player's slot.
   * @param speed (float): The speed value to set.
   */
  inline void SetClientSpeed(int32_t playerSlot, float speed) {
    using SetClientSpeedFn = void (*)(int32_t, float);
    static SetClientSpeedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientSpeed", reinterpret_cast<void**>(&__func));
    __func(playerSlot, speed);
  }

  /**
   * @brief Retrieves the amount of money a client has.
   * @function GetClientMoney
   * @param playerSlot (int32): The index of the player's slot.
   * @return int32: The amount of money the client has, or 0 if the player slot is invalid.
   */
  inline int32_t GetClientMoney(int32_t playerSlot) {
    using GetClientMoneyFn = int32_t (*)(int32_t);
    static GetClientMoneyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientMoney", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the amount of money for a client.
   * @function SetClientMoney
   * @param playerSlot (int32): The index of the player's slot.
   * @param money (int32): The amount of money to set.
   */
  inline void SetClientMoney(int32_t playerSlot, int32_t money) {
    using SetClientMoneyFn = void (*)(int32_t, int32_t);
    static SetClientMoneyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientMoney", reinterpret_cast<void**>(&__func));
    __func(playerSlot, money);
  }

  /**
   * @brief Retrieves the number of kills for a client.
   * @function GetClientKills
   * @param playerSlot (int32): The index of the player's slot.
   * @return int32: The number of kills the client has, or 0 if the player slot is invalid.
   */
  inline int32_t GetClientKills(int32_t playerSlot) {
    using GetClientKillsFn = int32_t (*)(int32_t);
    static GetClientKillsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientKills", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the number of kills for a client.
   * @function SetClientKills
   * @param playerSlot (int32): The index of the player's slot.
   * @param kills (int32): The number of kills to set.
   */
  inline void SetClientKills(int32_t playerSlot, int32_t kills) {
    using SetClientKillsFn = void (*)(int32_t, int32_t);
    static SetClientKillsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientKills", reinterpret_cast<void**>(&__func));
    __func(playerSlot, kills);
  }

  /**
   * @brief Retrieves the number of deaths for a client.
   * @function GetClientDeaths
   * @param playerSlot (int32): The index of the player's slot.
   * @return int32: The number of deaths the client has, or 0 if the player slot is invalid.
   */
  inline int32_t GetClientDeaths(int32_t playerSlot) {
    using GetClientDeathsFn = int32_t (*)(int32_t);
    static GetClientDeathsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientDeaths", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the number of deaths for a client.
   * @function SetClientDeaths
   * @param playerSlot (int32): The index of the player's slot.
   * @param deaths (int32): The number of deaths to set.
   */
  inline void SetClientDeaths(int32_t playerSlot, int32_t deaths) {
    using SetClientDeathsFn = void (*)(int32_t, int32_t);
    static SetClientDeathsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientDeaths", reinterpret_cast<void**>(&__func));
    __func(playerSlot, deaths);
  }

  /**
   * @brief Retrieves the number of assists for a client.
   * @function GetClientAssists
   * @param playerSlot (int32): The index of the player's slot.
   * @return int32: The number of assists the client has, or 0 if the player slot is invalid.
   */
  inline int32_t GetClientAssists(int32_t playerSlot) {
    using GetClientAssistsFn = int32_t (*)(int32_t);
    static GetClientAssistsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientAssists", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the number of assists for a client.
   * @function SetClientAssists
   * @param playerSlot (int32): The index of the player's slot.
   * @param assists (int32): The number of assists to set.
   */
  inline void SetClientAssists(int32_t playerSlot, int32_t assists) {
    using SetClientAssistsFn = void (*)(int32_t, int32_t);
    static SetClientAssistsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientAssists", reinterpret_cast<void**>(&__func));
    __func(playerSlot, assists);
  }

  /**
   * @brief Retrieves the total damage dealt by a client.
   * @function GetClientDamage
   * @param playerSlot (int32): The index of the player's slot.
   * @return int32: The total damage dealt by the client, or 0 if the player slot is invalid.
   */
  inline int32_t GetClientDamage(int32_t playerSlot) {
    using GetClientDamageFn = int32_t (*)(int32_t);
    static GetClientDamageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientDamage", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Sets the total damage dealt by a client.
   * @function SetClientDamage
   * @param playerSlot (int32): The index of the player's slot.
   * @param damage (int32): The amount of damage to set.
   */
  inline void SetClientDamage(int32_t playerSlot, int32_t damage) {
    using SetClientDamageFn = void (*)(int32_t, int32_t);
    static SetClientDamageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetClientDamage", reinterpret_cast<void**>(&__func));
    __func(playerSlot, damage);
  }

  /**
   * @brief Creates a console command as an administrative command.
   * @function AddAdminCommand
   * @param name (string): The name of the console command.
   * @param adminFlags (int64): The admin flags that indicate which admin level can use this command.
   * @param description (string): A brief description of what the command does.
   * @param flags (int64): Command flags that define the behavior of the command.
   * @param callback (function): A callback function that is invoked when the command is executed.
   * @param type (uint8): Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @return bool: true if the command was successfully created; otherwise, false.
   */
  inline bool AddAdminCommand(plg::string name, int64_t adminFlags, plg::string description, int64_t flags, function callback, uint8_t type) {
    using AddAdminCommandFn = bool (*)(plg::string, int64_t, plg::string, int64_t, function, uint8_t);
    static AddAdminCommandFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddAdminCommand", reinterpret_cast<void**>(&__func));
    return __func(name, adminFlags, description, flags, callback, type);
  }

  /**
   * @brief Creates a console command or hooks an already existing one.
   * @function AddConsoleCommand
   * @param name (string): The name of the console command.
   * @param description (string): A brief description of what the command does.
   * @param flags (int64): Command flags that define the behavior of the command.
   * @param callback (function): A callback function that is invoked when the command is executed.
   * @param type (uint8): Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @return bool: true if the command was successfully created; otherwise, false.
   */
  inline bool AddConsoleCommand(plg::string name, plg::string description, int64_t flags, function callback, uint8_t type) {
    using AddConsoleCommandFn = bool (*)(plg::string, plg::string, int64_t, function, uint8_t);
    static AddConsoleCommandFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddConsoleCommand", reinterpret_cast<void**>(&__func));
    return __func(name, description, flags, callback, type);
  }

  /**
   * @brief Removes a console command from the system.
   * @function RemoveCommand
   * @param name (string): The name of the command to be removed.
   * @param callback (function): The callback function associated with the command to be removed.
   * @return bool: true if the command was successfully removed; otherwise, false.
   */
  inline bool RemoveCommand(plg::string name, function callback) {
    using RemoveCommandFn = bool (*)(plg::string, function);
    static RemoveCommandFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RemoveCommand", reinterpret_cast<void**>(&__func));
    return __func(name, callback);
  }

  /**
   * @brief Adds a callback that will fire when a command is sent to the server.
   * @function AddCommandListener
   * @param name (string): The name of the command.
   * @param callback (function): The callback function that will be invoked when the command is executed.
   * @param type (uint8): Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @return bool: Returns true if the callback was successfully added, false otherwise.
   */
  inline bool AddCommandListener(plg::string name, function callback, uint8_t type) {
    using AddCommandListenerFn = bool (*)(plg::string, function, uint8_t);
    static AddCommandListenerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddCommandListener", reinterpret_cast<void**>(&__func));
    return __func(name, callback, type);
  }

  /**
   * @brief Removes a callback that fires when a command is sent to the server.
   * @function RemoveCommandListener
   * @param name (string): The name of the command.
   * @param callback (function): The callback function to be removed.
   * @param type (uint8): Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @return bool: Returns true if the callback was successfully removed, false otherwise.
   */
  inline bool RemoveCommandListener(plg::string name, function callback, uint8_t type) {
    using RemoveCommandListenerFn = bool (*)(plg::string, function, uint8_t);
    static RemoveCommandListenerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RemoveCommandListener", reinterpret_cast<void**>(&__func));
    return __func(name, callback, type);
  }

  /**
   * @brief Executes a server command as if it were run on the server console or through RCON.
   * @function ServerCommand
   * @param command (string): The command to execute on the server.
   */
  inline void ServerCommand(plg::string command) {
    using ServerCommandFn = void (*)(plg::string);
    static ServerCommandFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ServerCommand", reinterpret_cast<void**>(&__func));
    __func(command);
  }

  /**
   * @brief Executes a server command as if it were on the server console (or RCON) and stores the printed text into buffer.
   * @function ServerCommandEx
   * @param command (string): The command to execute on the server.
   * @return string: String to store command result into.
   */
  inline plg::string ServerCommandEx(plg::string command) {
    using ServerCommandExFn = plg::string (*)(plg::string);
    static ServerCommandExFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ServerCommandEx", reinterpret_cast<void**>(&__func));
    return __func(command);
  }

  /**
   * @brief Executes a client command.
   * @function ClientCommand
   * @param playerSlot (int32): The index of the client executing the command.
   * @param command (string): The command to execute on the client.
   */
  inline void ClientCommand(int32_t playerSlot, plg::string command) {
    using ClientCommandFn = void (*)(int32_t, plg::string);
    static ClientCommandFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ClientCommand", reinterpret_cast<void**>(&__func));
    __func(playerSlot, command);
  }

  /**
   * @brief Executes a client command on the server without network communication.
   * @function FakeClientCommand
   * @param playerSlot (int32): The index of the client.
   * @param command (string): The command to be executed by the client.
   */
  inline void FakeClientCommand(int32_t playerSlot, plg::string command) {
    using FakeClientCommandFn = void (*)(int32_t, plg::string);
    static FakeClientCommandFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FakeClientCommand", reinterpret_cast<void**>(&__func));
    __func(playerSlot, command);
  }

  /**
   * @brief Sends a message to the server console.
   * @function PrintToServer
   * @param msg (string): The message to be sent to the server console.
   */
  inline void PrintToServer(plg::string msg) {
    using PrintToServerFn = void (*)(plg::string);
    static PrintToServerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintToServer", reinterpret_cast<void**>(&__func));
    __func(msg);
  }

  /**
   * @brief Sends a message to a client's console.
   * @function PrintToConsole
   * @param playerSlot (int32): The index of the player's slot to whom the message will be sent.
   * @param message (string): The message to be sent to the client's console.
   */
  inline void PrintToConsole(int32_t playerSlot, plg::string message) {
    using PrintToConsoleFn = void (*)(int32_t, plg::string);
    static PrintToConsoleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintToConsole", reinterpret_cast<void**>(&__func));
    __func(playerSlot, message);
  }

  /**
   * @brief Prints a message to a specific client in the chat area.
   * @function PrintToChat
   * @param playerSlot (int32): The index of the player's slot to whom the message will be sent.
   * @param message (string): The message to be printed in the chat area.
   */
  inline void PrintToChat(int32_t playerSlot, plg::string message) {
    using PrintToChatFn = void (*)(int32_t, plg::string);
    static PrintToChatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintToChat", reinterpret_cast<void**>(&__func));
    __func(playerSlot, message);
  }

  /**
   * @brief Prints a message to a specific client in the center of the screen.
   * @function PrintCenterText
   * @param playerSlot (int32): The index of the player's slot to whom the message will be sent.
   * @param message (string): The message to be printed in the center of the screen.
   */
  inline void PrintCenterText(int32_t playerSlot, plg::string message) {
    using PrintCenterTextFn = void (*)(int32_t, plg::string);
    static PrintCenterTextFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintCenterText", reinterpret_cast<void**>(&__func));
    __func(playerSlot, message);
  }

  /**
   * @brief Prints a message to a specific client with an alert box.
   * @function PrintAlertText
   * @param playerSlot (int32): The index of the player's slot to whom the message will be sent.
   * @param message (string): The message to be printed in the alert box.
   */
  inline void PrintAlertText(int32_t playerSlot, plg::string message) {
    using PrintAlertTextFn = void (*)(int32_t, plg::string);
    static PrintAlertTextFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintAlertText", reinterpret_cast<void**>(&__func));
    __func(playerSlot, message);
  }

  /**
   * @brief Prints a html message to a specific client in the center of the screen.
   * @function PrintCentreHtml
   * @param playerSlot (int32): The index of the player's slot to whom the message will be sent.
   * @param message (string): The HTML-formatted message to be printed.
   * @param duration (int32): The duration of the message in seconds.
   */
  inline void PrintCentreHtml(int32_t playerSlot, plg::string message, int32_t duration) {
    using PrintCentreHtmlFn = void (*)(int32_t, plg::string, int32_t);
    static PrintCentreHtmlFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintCentreHtml", reinterpret_cast<void**>(&__func));
    __func(playerSlot, message, duration);
  }

  /**
   * @brief Sends a message to every client's console.
   * @function PrintToConsoleAll
   * @param message (string): The message to be sent to all clients' consoles.
   */
  inline void PrintToConsoleAll(plg::string message) {
    using PrintToConsoleAllFn = void (*)(plg::string);
    static PrintToConsoleAllFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintToConsoleAll", reinterpret_cast<void**>(&__func));
    __func(message);
  }

  /**
   * @brief Prints a message to all clients in the chat area.
   * @function PrintToChatAll
   * @param message (string): The message to be printed in the chat area for all clients.
   */
  inline void PrintToChatAll(plg::string message) {
    using PrintToChatAllFn = void (*)(plg::string);
    static PrintToChatAllFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintToChatAll", reinterpret_cast<void**>(&__func));
    __func(message);
  }

  /**
   * @brief Prints a message to all clients in the center of the screen.
   * @function PrintCenterTextAll
   * @param message (string): The message to be printed in the center of the screen for all clients.
   */
  inline void PrintCenterTextAll(plg::string message) {
    using PrintCenterTextAllFn = void (*)(plg::string);
    static PrintCenterTextAllFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintCenterTextAll", reinterpret_cast<void**>(&__func));
    __func(message);
  }

  /**
   * @brief Prints a message to all clients with an alert box.
   * @function PrintAlertTextAll
   * @param message (string): The message to be printed in an alert box for all clients.
   */
  inline void PrintAlertTextAll(plg::string message) {
    using PrintAlertTextAllFn = void (*)(plg::string);
    static PrintAlertTextAllFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintAlertTextAll", reinterpret_cast<void**>(&__func));
    __func(message);
  }

  /**
   * @brief Prints a html message to all clients in the center of the screen.
   * @function PrintCentreHtmlAll
   * @param message (string): The HTML-formatted message to be printed in the center of the screen for all clients.
   * @param duration (int32): The duration of the message in seconds.
   */
  inline void PrintCentreHtmlAll(plg::string message, int32_t duration) {
    using PrintCentreHtmlAllFn = void (*)(plg::string, int32_t);
    static PrintCentreHtmlAllFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintCentreHtmlAll", reinterpret_cast<void**>(&__func));
    __func(message, duration);
  }

  /**
   * @brief Prints a colored message to a specific client in the chat area.
   * @function PrintToChatColored
   * @param playerSlot (int32): The index of the player's slot to whom the message will be sent.
   * @param message (string): The message to be printed in the chat area with color.
   */
  inline void PrintToChatColored(int32_t playerSlot, plg::string message) {
    using PrintToChatColoredFn = void (*)(int32_t, plg::string);
    static PrintToChatColoredFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintToChatColored", reinterpret_cast<void**>(&__func));
    __func(playerSlot, message);
  }

  /**
   * @brief Prints a colored message to all clients in the chat area.
   * @function PrintToChatColoredAll
   * @param message (string): The colored message to be printed in the chat area for all clients.
   */
  inline void PrintToChatColoredAll(plg::string message) {
    using PrintToChatColoredAllFn = void (*)(plg::string);
    static PrintToChatColoredAllFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PrintToChatColoredAll", reinterpret_cast<void**>(&__func));
    __func(message);
  }

  /**
   * @brief Sends a reply message to a player or to the server console depending on the command context.
   * @function ReplyToCommand
   * @param context (int32): The context from which the command was called (e.g., Console or Chat).
   * @param playerSlot (int32): The slot/index of the player receiving the message.
   * @param message (string): The message string to be sent as a reply.
   */
  inline void ReplyToCommand(int32_t context, int32_t playerSlot, plg::string message) {
    using ReplyToCommandFn = void (*)(int32_t, int32_t, plg::string);
    static ReplyToCommandFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ReplyToCommand", reinterpret_cast<void**>(&__func));
    __func(context, playerSlot, message);
  }

  /**
   * @brief Creates a new console variable.
   * @function CreateConVar
   * @param name (string): The name of the console variable.
   * @param defaultValue (any): The default value of the console variable.
   * @param description (string): A description of the console variable's purpose.
   * @param flags (int64): Additional flags for the console variable.
   * @return uint64: A handle to the created console variable.
   */
  inline uint64_t CreateConVar(plg::string name, plg::any defaultValue, plg::string description, int64_t flags) {
    using CreateConVarFn = uint64_t (*)(plg::string, plg::any, plg::string, int64_t);
    static CreateConVarFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVar", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags);
  }

  /**
   * @brief Creates a new boolean console variable.
   * @function CreateConVarBool
   * @param name (string): The name of the console variable.
   * @param defaultValue (bool): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (bool): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (bool): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarBool(plg::string name, bool defaultValue, plg::string description, int64_t flags, bool hasMin, bool min, bool hasMax, bool max) {
    using CreateConVarBoolFn = uint64_t (*)(plg::string, bool, plg::string, int64_t, bool, bool, bool, bool);
    static CreateConVarBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarBool", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new 16-bit signed integer console variable.
   * @function CreateConVarInt16
   * @param name (string): The name of the console variable.
   * @param defaultValue (int16): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (int16): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (int16): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarInt16(plg::string name, int16_t defaultValue, plg::string description, int64_t flags, bool hasMin, int16_t min, bool hasMax, int16_t max) {
    using CreateConVarInt16Fn = uint64_t (*)(plg::string, int16_t, plg::string, int64_t, bool, int16_t, bool, int16_t);
    static CreateConVarInt16Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarInt16", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new 16-bit unsigned integer console variable.
   * @function CreateConVarUInt16
   * @param name (string): The name of the console variable.
   * @param defaultValue (uint16): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (uint16): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (uint16): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarUInt16(plg::string name, uint16_t defaultValue, plg::string description, int64_t flags, bool hasMin, uint16_t min, bool hasMax, uint16_t max) {
    using CreateConVarUInt16Fn = uint64_t (*)(plg::string, uint16_t, plg::string, int64_t, bool, uint16_t, bool, uint16_t);
    static CreateConVarUInt16Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarUInt16", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new 32-bit signed integer console variable.
   * @function CreateConVarInt32
   * @param name (string): The name of the console variable.
   * @param defaultValue (int32): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (int32): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (int32): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarInt32(plg::string name, int32_t defaultValue, plg::string description, int64_t flags, bool hasMin, int32_t min, bool hasMax, int32_t max) {
    using CreateConVarInt32Fn = uint64_t (*)(plg::string, int32_t, plg::string, int64_t, bool, int32_t, bool, int32_t);
    static CreateConVarInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarInt32", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new 32-bit unsigned integer console variable.
   * @function CreateConVarUInt32
   * @param name (string): The name of the console variable.
   * @param defaultValue (uint32): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (uint32): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (uint32): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarUInt32(plg::string name, uint32_t defaultValue, plg::string description, int64_t flags, bool hasMin, uint32_t min, bool hasMax, uint32_t max) {
    using CreateConVarUInt32Fn = uint64_t (*)(plg::string, uint32_t, plg::string, int64_t, bool, uint32_t, bool, uint32_t);
    static CreateConVarUInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarUInt32", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new 64-bit signed integer console variable.
   * @function CreateConVarInt64
   * @param name (string): The name of the console variable.
   * @param defaultValue (int64): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (int64): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (int64): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarInt64(plg::string name, int64_t defaultValue, plg::string description, int64_t flags, bool hasMin, int64_t min, bool hasMax, int64_t max) {
    using CreateConVarInt64Fn = uint64_t (*)(plg::string, int64_t, plg::string, int64_t, bool, int64_t, bool, int64_t);
    static CreateConVarInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarInt64", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new 64-bit unsigned integer console variable.
   * @function CreateConVarUInt64
   * @param name (string): The name of the console variable.
   * @param defaultValue (uint64): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (uint64): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (uint64): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarUInt64(plg::string name, uint64_t defaultValue, plg::string description, int64_t flags, bool hasMin, uint64_t min, bool hasMax, uint64_t max) {
    using CreateConVarUInt64Fn = uint64_t (*)(plg::string, uint64_t, plg::string, int64_t, bool, uint64_t, bool, uint64_t);
    static CreateConVarUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarUInt64", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new floating-point console variable.
   * @function CreateConVarFloat
   * @param name (string): The name of the console variable.
   * @param defaultValue (float): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (float): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (float): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarFloat(plg::string name, float defaultValue, plg::string description, int64_t flags, bool hasMin, float min, bool hasMax, float max) {
    using CreateConVarFloatFn = uint64_t (*)(plg::string, float, plg::string, int64_t, bool, float, bool, float);
    static CreateConVarFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarFloat", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new double-precision console variable.
   * @function CreateConVarDouble
   * @param name (string): The name of the console variable.
   * @param defaultValue (double): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (double): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (double): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarDouble(plg::string name, double defaultValue, plg::string description, int64_t flags, bool hasMin, double min, bool hasMax, double max) {
    using CreateConVarDoubleFn = uint64_t (*)(plg::string, double, plg::string, int64_t, bool, double, bool, double);
    static CreateConVarDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarDouble", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new color console variable.
   * @function CreateConVarColor
   * @param name (string): The name of the console variable.
   * @param defaultValue (int32): The default color value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (int32): The minimum color value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (int32): The maximum color value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarColor(plg::string name, int32_t defaultValue, plg::string description, int64_t flags, bool hasMin, int32_t min, bool hasMax, int32_t max) {
    using CreateConVarColorFn = uint64_t (*)(plg::string, int32_t, plg::string, int64_t, bool, int32_t, bool, int32_t);
    static CreateConVarColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarColor", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new 2D vector console variable.
   * @function CreateConVarVector2
   * @param name (string): The name of the console variable.
   * @param defaultValue (vec2): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (vec2): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (vec2): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarVector2(plg::string name, plg::vec2 defaultValue, plg::string description, int64_t flags, bool hasMin, plg::vec2 min, bool hasMax, plg::vec2 max) {
    using CreateConVarVector2Fn = uint64_t (*)(plg::string, plg::vec2, plg::string, int64_t, bool, plg::vec2, bool, plg::vec2);
    static CreateConVarVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarVector2", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new 3D vector console variable.
   * @function CreateConVarVector3
   * @param name (string): The name of the console variable.
   * @param defaultValue (vec3): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (vec3): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (vec3): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarVector3(plg::string name, plg::vec3 defaultValue, plg::string description, int64_t flags, bool hasMin, plg::vec3 min, bool hasMax, plg::vec3 max) {
    using CreateConVarVector3Fn = uint64_t (*)(plg::string, plg::vec3, plg::string, int64_t, bool, plg::vec3, bool, plg::vec3);
    static CreateConVarVector3Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarVector3", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new 4D vector console variable.
   * @function CreateConVarVector4
   * @param name (string): The name of the console variable.
   * @param defaultValue (vec4): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (vec4): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (vec4): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarVector4(plg::string name, plg::vec4 defaultValue, plg::string description, int64_t flags, bool hasMin, plg::vec4 min, bool hasMax, plg::vec4 max) {
    using CreateConVarVector4Fn = uint64_t (*)(plg::string, plg::vec4, plg::string, int64_t, bool, plg::vec4, bool, plg::vec4);
    static CreateConVarVector4Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarVector4", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new quaternion angle console variable.
   * @function CreateConVarQAngle
   * @param name (string): The name of the console variable.
   * @param defaultValue (vec3): The default value for the console variable.
   * @param description (string): A brief description of the console variable.
   * @param flags (int64): Flags that define the behavior of the console variable.
   * @param hasMin (bool): Indicates if a minimum value is provided.
   * @param min (vec3): The minimum value if hasMin is true.
   * @param hasMax (bool): Indicates if a maximum value is provided.
   * @param max (vec3): The maximum value if hasMax is true.
   * @return uint64: A handle to the created console variable data.
   */
  inline uint64_t CreateConVarQAngle(plg::string name, plg::vec3 defaultValue, plg::string description, int64_t flags, bool hasMin, plg::vec3 min, bool hasMax, plg::vec3 max) {
    using CreateConVarQAngleFn = uint64_t (*)(plg::string, plg::vec3, plg::string, int64_t, bool, plg::vec3, bool, plg::vec3);
    static CreateConVarQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarQAngle", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags, hasMin, min, hasMax, max);
  }

  /**
   * @brief Creates a new string console variable.
   * @function CreateConVarString
   * @param name (string): The name of the console variable.
   * @param defaultValue (string): The default value of the console variable.
   * @param description (string): A description of the console variable's purpose.
   * @param flags (int64): Additional flags for the console variable.
   * @return uint64: A handle to the created console variable.
   */
  inline uint64_t CreateConVarString(plg::string name, plg::string defaultValue, plg::string description, int64_t flags) {
    using CreateConVarStringFn = uint64_t (*)(plg::string, plg::string, plg::string, int64_t);
    static CreateConVarStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateConVarString", reinterpret_cast<void**>(&__func));
    return __func(name, defaultValue, description, flags);
  }

  /**
   * @brief Searches for a console variable.
   * @function FindConVar
   * @param name (string): The name of the console variable to search for.
   * @return uint64: A handle to the console variable data if found; otherwise, nullptr.
   */
  inline uint64_t FindConVar(plg::string name) {
    using FindConVarFn = uint64_t (*)(plg::string);
    static FindConVarFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindConVar", reinterpret_cast<void**>(&__func));
    return __func(name);
  }

  /**
   * @brief Searches for a console variable of a specific type.
   * @function FindConVar2
   * @param name (string): The name of the console variable to search for.
   * @param type (int16): The type of the console variable to search for.
   * @return uint64: A handle to the console variable data if found; otherwise, nullptr.
   */
  inline uint64_t FindConVar2(plg::string name, int16_t type) {
    using FindConVar2Fn = uint64_t (*)(plg::string, int16_t);
    static FindConVar2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindConVar2", reinterpret_cast<void**>(&__func));
    return __func(name, type);
  }

  /**
   * @brief Creates a hook for when a console variable's value is changed.
   * @function HookConVarChange
   * @param conVarHandle (uint64): TThe handle to the console variable data.
   * @param callback (function): The callback function to be executed when the variable's value changes.
   */
  inline void HookConVarChange(uint64_t conVarHandle, function callback) {
    using HookConVarChangeFn = void (*)(uint64_t, function);
    static HookConVarChangeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.HookConVarChange", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, callback);
  }

  /**
   * @brief Removes a hook for when a console variable's value is changed.
   * @function UnhookConVarChange
   * @param uint64 (string): The handle to the console variable data.
   * @param callback (function): The callback function to be removed.
   */
  inline void UnhookConVarChange(plg::string uint64, function callback) {
    using UnhookConVarChangeFn = void (*)(plg::string, function);
    static UnhookConVarChangeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UnhookConVarChange", reinterpret_cast<void**>(&__func));
    __func(uint64, callback);
  }

  /**
   * @brief Checks if a specific flag is set for a console variable.
   * @function IsConVarFlagSet
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param flag (int64): The flag to check against the console variable.
   * @return bool: True if the flag is set; otherwise, false.
   */
  inline bool IsConVarFlagSet(uint64_t conVarHandle, int64_t flag) {
    using IsConVarFlagSetFn = bool (*)(uint64_t, int64_t);
    static IsConVarFlagSetFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsConVarFlagSet", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle, flag);
  }

  /**
   * @brief Adds flags to a console variable.
   * @function AddConVarFlags
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param flags (int64): The flags to be added.
   */
  inline void AddConVarFlags(uint64_t conVarHandle, int64_t flags) {
    using AddConVarFlagsFn = void (*)(uint64_t, int64_t);
    static AddConVarFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddConVarFlags", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, flags);
  }

  /**
   * @brief Removes flags from a console variable.
   * @function RemoveConVarFlags
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param flags (int64): The flags to be removed.
   */
  inline void RemoveConVarFlags(uint64_t conVarHandle, int64_t flags) {
    using RemoveConVarFlagsFn = void (*)(uint64_t, int64_t);
    static RemoveConVarFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RemoveConVarFlags", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, flags);
  }

  /**
   * @brief Retrieves the current flags of a console variable.
   * @function GetConVarFlags
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return int64: The current flags set on the console variable.
   */
  inline int64_t GetConVarFlags(uint64_t conVarHandle) {
    using GetConVarFlagsFn = int64_t (*)(uint64_t);
    static GetConVarFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarFlags", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Gets the specified bound (max or min) of a console variable and stores it in the output string.
   * @function GetConVarBounds
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param max (bool): Indicates whether to get the maximum (true) or minimum (false) bound.
   * @return string: The bound value.
   */
  inline plg::string GetConVarBounds(uint64_t conVarHandle, bool max) {
    using GetConVarBoundsFn = plg::string (*)(uint64_t, bool);
    static GetConVarBoundsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarBounds", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle, max);
  }

  /**
   * @brief Sets the specified bound (max or min) for a console variable.
   * @function SetConVarBounds
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param max (bool): Indicates whether to set the maximum (true) or minimum (false) bound.
   * @param value (string): The value to set as the bound.
   */
  inline void SetConVarBounds(uint64_t conVarHandle, bool max, plg::string value) {
    using SetConVarBoundsFn = void (*)(uint64_t, bool, plg::string);
    static SetConVarBoundsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarBounds", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, max, value);
  }

  /**
   * @brief Retrieves the default value of a console variable and stores it in the output string.
   * @function GetConVarDefault
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return string: The output value in string format.
   */
  inline plg::string GetConVarDefault(uint64_t conVarHandle) {
    using GetConVarDefaultFn = plg::string (*)(uint64_t);
    static GetConVarDefaultFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarDefault", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a console variable and stores it in the output string.
   * @function GetConVarValue
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return string: The output value in string format.
   */
  inline plg::string GetConVarValue(uint64_t conVarHandle) {
    using GetConVarValueFn = plg::string (*)(uint64_t);
    static GetConVarValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarValue", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a console variable and stores it in the output.
   * @function GetConVar
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return any: The output value.
   */
  inline plg::any GetConVar(uint64_t conVarHandle) {
    using GetConVarFn = plg::any (*)(uint64_t);
    static GetConVarFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVar", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a boolean console variable.
   * @function GetConVarBool
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return bool: The current boolean value of the console variable.
   */
  inline bool GetConVarBool(uint64_t conVarHandle) {
    using GetConVarBoolFn = bool (*)(uint64_t);
    static GetConVarBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarBool", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a signed 16-bit integer console variable.
   * @function GetConVarInt16
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return int16: The current int16_t value of the console variable.
   */
  inline int16_t GetConVarInt16(uint64_t conVarHandle) {
    using GetConVarInt16Fn = int16_t (*)(uint64_t);
    static GetConVarInt16Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarInt16", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of an unsigned 16-bit integer console variable.
   * @function GetConVarUInt16
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return uint16: The current uint16_t value of the console variable.
   */
  inline uint16_t GetConVarUInt16(uint64_t conVarHandle) {
    using GetConVarUInt16Fn = uint16_t (*)(uint64_t);
    static GetConVarUInt16Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarUInt16", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a signed 32-bit integer console variable.
   * @function GetConVarInt32
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return int32: The current int32_t value of the console variable.
   */
  inline int32_t GetConVarInt32(uint64_t conVarHandle) {
    using GetConVarInt32Fn = int32_t (*)(uint64_t);
    static GetConVarInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarInt32", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of an unsigned 32-bit integer console variable.
   * @function GetConVarUInt32
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return uint32: The current uint32_t value of the console variable.
   */
  inline uint32_t GetConVarUInt32(uint64_t conVarHandle) {
    using GetConVarUInt32Fn = uint32_t (*)(uint64_t);
    static GetConVarUInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarUInt32", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a signed 64-bit integer console variable.
   * @function GetConVarInt64
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return int64: The current int64_t value of the console variable.
   */
  inline int64_t GetConVarInt64(uint64_t conVarHandle) {
    using GetConVarInt64Fn = int64_t (*)(uint64_t);
    static GetConVarInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarInt64", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of an unsigned 64-bit integer console variable.
   * @function GetConVarUInt64
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return uint64: The current uint64_t value of the console variable.
   */
  inline uint64_t GetConVarUInt64(uint64_t conVarHandle) {
    using GetConVarUInt64Fn = uint64_t (*)(uint64_t);
    static GetConVarUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarUInt64", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a float console variable.
   * @function GetConVarFloat
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return float: The current float value of the console variable.
   */
  inline float GetConVarFloat(uint64_t conVarHandle) {
    using GetConVarFloatFn = float (*)(uint64_t);
    static GetConVarFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarFloat", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a double console variable.
   * @function GetConVarDouble
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return double: The current double value of the console variable.
   */
  inline double GetConVarDouble(uint64_t conVarHandle) {
    using GetConVarDoubleFn = double (*)(uint64_t);
    static GetConVarDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarDouble", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a string console variable.
   * @function GetConVarString
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return string: The current string value of the console variable.
   */
  inline plg::string GetConVarString(uint64_t conVarHandle) {
    using GetConVarStringFn = plg::string (*)(uint64_t);
    static GetConVarStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarString", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a Color console variable.
   * @function GetConVarColor
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return int32: The current Color value of the console variable.
   */
  inline int32_t GetConVarColor(uint64_t conVarHandle) {
    using GetConVarColorFn = int32_t (*)(uint64_t);
    static GetConVarColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarColor", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a Vector2D console variable.
   * @function GetConVarVector2
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return vec2: The current Vector2D value of the console variable.
   */
  inline plg::vec2 GetConVarVector2(uint64_t conVarHandle) {
    using GetConVarVector2Fn = plg::vec2 (*)(uint64_t);
    static GetConVarVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarVector2", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a Vector console variable.
   * @function GetConVarVector
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return vec3: The current Vector value of the console variable.
   */
  inline plg::vec3 GetConVarVector(uint64_t conVarHandle) {
    using GetConVarVectorFn = plg::vec3 (*)(uint64_t);
    static GetConVarVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarVector", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a Vector4D console variable.
   * @function GetConVarVector4
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return vec4: The current Vector4D value of the console variable.
   */
  inline plg::vec4 GetConVarVector4(uint64_t conVarHandle) {
    using GetConVarVector4Fn = plg::vec4 (*)(uint64_t);
    static GetConVarVector4Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarVector4", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Retrieves the current value of a QAngle console variable.
   * @function GetConVarQAngle
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @return vec3: The current QAngle value of the console variable.
   */
  inline plg::vec3 GetConVarQAngle(uint64_t conVarHandle) {
    using GetConVarQAngleFn = plg::vec3 (*)(uint64_t);
    static GetConVarQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConVarQAngle", reinterpret_cast<void**>(&__func));
    return __func(conVarHandle);
  }

  /**
   * @brief Sets the value of a console variable.
   * @function SetConVarValue
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (string): The string value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarValue(uint64_t conVarHandle, plg::string value, bool replicate, bool notify) {
    using SetConVarValueFn = void (*)(uint64_t, plg::string, bool, bool);
    static SetConVarValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarValue", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a console variable.
   * @function SetConVar
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (any): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVar(uint64_t conVarHandle, plg::any value, bool replicate, bool notify) {
    using SetConVarFn = void (*)(uint64_t, plg::any, bool, bool);
    static SetConVarFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVar", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a boolean console variable.
   * @function SetConVarBool
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (bool): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarBool(uint64_t conVarHandle, bool value, bool replicate, bool notify) {
    using SetConVarBoolFn = void (*)(uint64_t, bool, bool, bool);
    static SetConVarBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarBool", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a signed 16-bit integer console variable.
   * @function SetConVarInt16
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (int16): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarInt16(uint64_t conVarHandle, int16_t value, bool replicate, bool notify) {
    using SetConVarInt16Fn = void (*)(uint64_t, int16_t, bool, bool);
    static SetConVarInt16Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarInt16", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of an unsigned 16-bit integer console variable.
   * @function SetConVarUInt16
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (uint16): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarUInt16(uint64_t conVarHandle, uint16_t value, bool replicate, bool notify) {
    using SetConVarUInt16Fn = void (*)(uint64_t, uint16_t, bool, bool);
    static SetConVarUInt16Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarUInt16", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a signed 32-bit integer console variable.
   * @function SetConVarInt32
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (int32): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarInt32(uint64_t conVarHandle, int32_t value, bool replicate, bool notify) {
    using SetConVarInt32Fn = void (*)(uint64_t, int32_t, bool, bool);
    static SetConVarInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarInt32", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of an unsigned 32-bit integer console variable.
   * @function SetConVarUInt32
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (uint32): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarUInt32(uint64_t conVarHandle, uint32_t value, bool replicate, bool notify) {
    using SetConVarUInt32Fn = void (*)(uint64_t, uint32_t, bool, bool);
    static SetConVarUInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarUInt32", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a signed 64-bit integer console variable.
   * @function SetConVarInt64
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (int64): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarInt64(uint64_t conVarHandle, int64_t value, bool replicate, bool notify) {
    using SetConVarInt64Fn = void (*)(uint64_t, int64_t, bool, bool);
    static SetConVarInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarInt64", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of an unsigned 64-bit integer console variable.
   * @function SetConVarUInt64
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (uint64): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarUInt64(uint64_t conVarHandle, uint64_t value, bool replicate, bool notify) {
    using SetConVarUInt64Fn = void (*)(uint64_t, uint64_t, bool, bool);
    static SetConVarUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarUInt64", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a floating-point console variable.
   * @function SetConVarFloat
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (float): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarFloat(uint64_t conVarHandle, float value, bool replicate, bool notify) {
    using SetConVarFloatFn = void (*)(uint64_t, float, bool, bool);
    static SetConVarFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarFloat", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a double-precision floating-point console variable.
   * @function SetConVarDouble
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (double): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarDouble(uint64_t conVarHandle, double value, bool replicate, bool notify) {
    using SetConVarDoubleFn = void (*)(uint64_t, double, bool, bool);
    static SetConVarDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarDouble", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a string console variable.
   * @function SetConVarString
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (string): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarString(uint64_t conVarHandle, plg::string value, bool replicate, bool notify) {
    using SetConVarStringFn = void (*)(uint64_t, plg::string, bool, bool);
    static SetConVarStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarString", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a color console variable.
   * @function SetConVarColor
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (int32): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarColor(uint64_t conVarHandle, int32_t value, bool replicate, bool notify) {
    using SetConVarColorFn = void (*)(uint64_t, int32_t, bool, bool);
    static SetConVarColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarColor", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a 2D vector console variable.
   * @function SetConVarVector2
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (vec2): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarVector2(uint64_t conVarHandle, plg::vec2 value, bool replicate, bool notify) {
    using SetConVarVector2Fn = void (*)(uint64_t, plg::vec2, bool, bool);
    static SetConVarVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarVector2", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a 3D vector console variable.
   * @function SetConVarVector3
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (vec3): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarVector3(uint64_t conVarHandle, plg::vec3 value, bool replicate, bool notify) {
    using SetConVarVector3Fn = void (*)(uint64_t, plg::vec3, bool, bool);
    static SetConVarVector3Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarVector3", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a 4D vector console variable.
   * @function SetConVarVector4
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (vec4): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarVector4(uint64_t conVarHandle, plg::vec4 value, bool replicate, bool notify) {
    using SetConVarVector4Fn = void (*)(uint64_t, plg::vec4, bool, bool);
    static SetConVarVector4Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarVector4", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Sets the value of a quaternion angle console variable.
   * @function SetConVarQAngle
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (vec3): The value to set for the console variable.
   * @param replicate (bool): If set to true, the new convar value will be set on all clients. This will only work if the convar has the FCVAR_REPLICATED flag and actually exists on clients.
   * @param notify (bool): If set to true, clients will be notified that the convar has changed. This will only work if the convar has the FCVAR_NOTIFY flag.
   */
  inline void SetConVarQAngle(uint64_t conVarHandle, plg::vec3 value, bool replicate, bool notify) {
    using SetConVarQAngleFn = void (*)(uint64_t, plg::vec3, bool, bool);
    static SetConVarQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetConVarQAngle", reinterpret_cast<void**>(&__func));
    __func(conVarHandle, value, replicate, notify);
  }

  /**
   * @brief Replicates a console variable value to a specific client. This does not change the actual console variable value.
   * @function SendConVarValue
   * @param playerSlot (int32): The index of the client to replicate the value to.
   * @param conVarHandle (uint64): The handle to the console variable data.
   * @param value (string): The value to send to the client.
   */
  inline void SendConVarValue(int32_t playerSlot, uint64_t conVarHandle, plg::string value) {
    using SendConVarValueFn = void (*)(int32_t, uint64_t, plg::string);
    static SendConVarValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SendConVarValue", reinterpret_cast<void**>(&__func));
    __func(playerSlot, conVarHandle, value);
  }

  /**
   * @brief Retrieves the value of a client's console variable and stores it in the output string.
   * @function GetClientConVarValue
   * @param playerSlot (int32): The index of the client whose console variable value is being retrieved.
   * @param convarName (string): The name of the console variable to retrieve.
   * @return string: The output string to store the client's console variable value.
   */
  inline plg::string GetClientConVarValue(int32_t playerSlot, plg::string convarName) {
    using GetClientConVarValueFn = plg::string (*)(int32_t, plg::string);
    static GetClientConVarValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetClientConVarValue", reinterpret_cast<void**>(&__func));
    return __func(playerSlot, convarName);
  }

  /**
   * @brief Replicates a console variable value to a specific fake client. This does not change the actual console variable value.
   * @function SetFakeClientConVarValue
   * @param playerSlot (int32): The index of the fake client to replicate the value to.
   * @param convarName (string): The name of the console variable.
   * @param convarValue (string): The value to set for the console variable.
   */
  inline void SetFakeClientConVarValue(int32_t playerSlot, plg::string convarName, plg::string convarValue) {
    using SetFakeClientConVarValueFn = void (*)(int32_t, plg::string, plg::string);
    static SetFakeClientConVarValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetFakeClientConVarValue", reinterpret_cast<void**>(&__func));
    __func(playerSlot, convarName, convarValue);
  }

  /**
   * @brief Starts a query to retrieve the value of a client's console variable.
   * @function QueryClientConVar
   * @param playerSlot (int32): The index of the player's slot to query the value from.
   * @param convarName (string): The name of client convar to query.
   * @param callback (function): A function to use as a callback when the query has finished.
   * @param data (any[]): Optional values to pass to the callback function.
   * @return int32: A cookie that uniquely identifies the query. Returns -1 on failure, such as when used on a bot.
   */
  inline int32_t QueryClientConVar(int32_t playerSlot, plg::string convarName, function callback, plg::vector<plg::any> data) {
    using QueryClientConVarFn = int32_t (*)(int32_t, plg::string, function, plg::vector<plg::any>);
    static QueryClientConVarFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.QueryClientConVar", reinterpret_cast<void**>(&__func));
    return __func(playerSlot, convarName, callback, data);
  }

  /**
   * @brief  Specifies that the given config file should be executed.
   * @function AutoExecConfig
   * @param conVarHandles (uint64[]): List of handles to the console variable data.
   * @param autoCreate (bool): If true, and the config file does not exist, such a config file will be automatically created and populated with information from the plugin's registered cvars.
   * @param name (string): Name of the config file, excluding the .cfg extension. Cannot be empty.
   * @param folder (string): Folder under cfg/ to use. By default this is "plugify." Can be empty.
   * @return bool: True on success, false otherwise.
   */
  inline bool AutoExecConfig(plg::vector<uint64_t> conVarHandles, bool autoCreate, plg::string name, plg::string folder) {
    using AutoExecConfigFn = bool (*)(plg::vector<uint64_t>, bool, plg::string, plg::string);
    static AutoExecConfigFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AutoExecConfig", reinterpret_cast<void**>(&__func));
    return __func(conVarHandles, autoCreate, name, folder);
  }

  /**
   * @brief Returns the current server language.
   * @function GetServerLanguage
   * @return string: The server language as a string.
   */
  inline plg::string GetServerLanguage() {
    using GetServerLanguageFn = plg::string (*)();
    static GetServerLanguageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetServerLanguage", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Finds a module by name.
   * @function FindModule
   * @param name (string): The name of the module to find.
   * @return ptr64: A pointer to the specified module.
   */
  inline void* FindModule(plg::string name) {
    using FindModuleFn = void* (*)(plg::string);
    static FindModuleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindModule", reinterpret_cast<void**>(&__func));
    return __func(name);
  }

  /**
   * @brief Finds an interface by name.
   * @function FindInterface
   * @param name (string): The name of the interface to find.
   * @return ptr64: A pointer to the interface.
   */
  inline void* FindInterface(plg::string name) {
    using FindInterfaceFn = void* (*)(plg::string);
    static FindInterfaceFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindInterface", reinterpret_cast<void**>(&__func));
    return __func(name);
  }

  /**
   * @brief Queries an interface from a specified module.
   * @function QueryInterface
   * @param module (string): The name of the module to query the interface from.
   * @param name (string): The name of the interface to find.
   * @return ptr64: A pointer to the queried interface.
   */
  inline void* QueryInterface(plg::string module, plg::string name) {
    using QueryInterfaceFn = void* (*)(plg::string, plg::string);
    static QueryInterfaceFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.QueryInterface", reinterpret_cast<void**>(&__func));
    return __func(module, name);
  }

  /**
   * @brief Returns the path of the game's directory.
   * @function GetGameDirectory
   * @return string: A reference to a string where the game directory path will be stored.
   */
  inline plg::string GetGameDirectory() {
    using GetGameDirectoryFn = plg::string (*)();
    static GetGameDirectoryFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameDirectory", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Returns the current map name.
   * @function GetCurrentMap
   * @return string: A reference to a string where the current map name will be stored.
   */
  inline plg::string GetCurrentMap() {
    using GetCurrentMapFn = plg::string (*)();
    static GetCurrentMapFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetCurrentMap", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Returns whether a specified map is valid or not.
   * @function IsMapValid
   * @param mapname (string): The name of the map to check for validity.
   * @return bool: True if the map is valid, false otherwise.
   */
  inline bool IsMapValid(plg::string mapname) {
    using IsMapValidFn = bool (*)(plg::string);
    static IsMapValidFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsMapValid", reinterpret_cast<void**>(&__func));
    return __func(mapname);
  }

  /**
   * @brief Returns the game time based on the game tick.
   * @function GetGameTime
   * @return float: The current game time.
   */
  inline float GetGameTime() {
    using GetGameTimeFn = float (*)();
    static GetGameTimeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameTime", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Returns the game's internal tick count.
   * @function GetGameTickCount
   * @return int32: The current tick count of the game.
   */
  inline int32_t GetGameTickCount() {
    using GetGameTickCountFn = int32_t (*)();
    static GetGameTickCountFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameTickCount", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Returns the time the game took processing the last frame.
   * @function GetGameFrameTime
   * @return float: The frame time of the last processed frame.
   */
  inline float GetGameFrameTime() {
    using GetGameFrameTimeFn = float (*)();
    static GetGameFrameTimeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameFrameTime", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Returns a high-precision time value for profiling the engine.
   * @function GetEngineTime
   * @return double: A high-precision time value.
   */
  inline double GetEngineTime() {
    using GetEngineTimeFn = double (*)();
    static GetEngineTimeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEngineTime", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Returns the maximum number of clients that can connect to the server.
   * @function GetMaxClients
   * @return int32: The maximum client count, or -1 if global variables are not initialized.
   */
  inline int32_t GetMaxClients() {
    using GetMaxClientsFn = int32_t (*)();
    static GetMaxClientsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetMaxClients", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Precaches a given file.
   * @function Precache
   * @param resource (string): The name of the resource to be precached.
   */
  inline void Precache(plg::string resource) {
    using PrecacheFn = void (*)(plg::string);
    static PrecacheFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Precache", reinterpret_cast<void**>(&__func));
    __func(resource);
  }

  /**
   * @brief Checks if a specified file is precached.
   * @function IsPrecached
   * @param resource (string): The name of the file to check.
   * @return bool
   */
  inline bool IsPrecached(plg::string resource) {
    using IsPrecachedFn = bool (*)(plg::string);
    static IsPrecachedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsPrecached", reinterpret_cast<void**>(&__func));
    return __func(resource);
  }

  /**
   * @brief Returns a pointer to the Economy Item System.
   * @function GetEconItemSystem
   * @return ptr64: A pointer to the Econ Item System.
   */
  inline void* GetEconItemSystem() {
    using GetEconItemSystemFn = void* (*)();
    static GetEconItemSystemFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEconItemSystem", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Checks if the server is currently paused.
   * @function IsServerPaused
   * @return bool: True if the server is paused, false otherwise.
   */
  inline bool IsServerPaused() {
    using IsServerPausedFn = bool (*)();
    static IsServerPausedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsServerPaused", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Queues a task to be executed on the next frame.
   * @function QueueTaskForNextFrame
   * @param callback (function): A callback function to be executed on the next frame.
   * @param userData (any[]): An array intended to hold user-related data, allowing for elements of any type.
   */
  inline void QueueTaskForNextFrame(function callback, plg::vector<plg::any> userData) {
    using QueueTaskForNextFrameFn = void (*)(function, plg::vector<plg::any>);
    static QueueTaskForNextFrameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.QueueTaskForNextFrame", reinterpret_cast<void**>(&__func));
    __func(callback, userData);
  }

  /**
   * @brief Queues a task to be executed on the next world update.
   * @function QueueTaskForNextWorldUpdate
   * @param callback (function): A callback function to be executed on the next world update.
   * @param userData (any[]): An array intended to hold user-related data, allowing for elements of any type.
   */
  inline void QueueTaskForNextWorldUpdate(function callback, plg::vector<plg::any> userData) {
    using QueueTaskForNextWorldUpdateFn = void (*)(function, plg::vector<plg::any>);
    static QueueTaskForNextWorldUpdateFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.QueueTaskForNextWorldUpdate", reinterpret_cast<void**>(&__func));
    __func(callback, userData);
  }

  /**
   * @brief Returns the duration of a specified sound.
   * @function GetSoundDuration
   * @param name (string): The name of the sound to check.
   * @return float: The duration of the sound in seconds.
   */
  inline float GetSoundDuration(plg::string name) {
    using GetSoundDurationFn = float (*)(plg::string);
    static GetSoundDurationFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetSoundDuration", reinterpret_cast<void**>(&__func));
    return __func(name);
  }

  /**
   * @brief Emits a sound from a specified entity.
   * @function EmitSound
   * @param entityHandle (int32): The handle of the entity that will emit the sound.
   * @param sound (string): The name of the sound to emit.
   * @param pitch (int32): The pitch of the sound.
   * @param volume (float): The volume of the sound.
   * @param delay (float): The delay before the sound is played.
   */
  inline void EmitSound(int32_t entityHandle, plg::string sound, int32_t pitch, float volume, float delay) {
    using EmitSoundFn = void (*)(int32_t, plg::string, int32_t, float, float);
    static EmitSoundFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.EmitSound", reinterpret_cast<void**>(&__func));
    __func(entityHandle, sound, pitch, volume, delay);
  }

  /**
   * @brief Stops a sound from a specified entity.
   * @function StopSound
   * @param entityHandle (int32): The handle of the entity that will stop the sound.
   * @param sound (string): The name of the sound to stop.
   */
  inline void StopSound(int32_t entityHandle, plg::string sound) {
    using StopSoundFn = void (*)(int32_t, plg::string);
    static StopSoundFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.StopSound", reinterpret_cast<void**>(&__func));
    __func(entityHandle, sound);
  }

  /**
   * @brief Emits a sound to a specific client.
   * @function EmitSoundToClient
   * @param playerSlot (int32): The index of the client to whom the sound will be emitted.
   * @param channel (int32): The channel through which the sound will be played.
   * @param sound (string): The name of the sound to emit.
   * @param volume (float): The volume of the sound.
   * @param soundLevel (int32): The level of the sound.
   * @param flags (int32): Additional flags for sound playback.
   * @param pitch (int32): The pitch of the sound.
   * @param origin (vec3): The origin of the sound in 3D space.
   * @param soundTime (float): The time at which the sound should be played.
   */
  inline void EmitSoundToClient(int32_t playerSlot, int32_t channel, plg::string sound, float volume, int32_t soundLevel, int32_t flags, int32_t pitch, plg::vec3 origin, float soundTime) {
    using EmitSoundToClientFn = void (*)(int32_t, int32_t, plg::string, float, int32_t, int32_t, int32_t, plg::vec3, float);
    static EmitSoundToClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.EmitSoundToClient", reinterpret_cast<void**>(&__func));
    __func(playerSlot, channel, sound, volume, soundLevel, flags, pitch, origin, soundTime);
  }

  /**
   * @brief Converts an entity index into an entity pointer.
   * @function EntIndexToEntPointer
   * @param entityIndex (int32): The index of the entity to convert.
   * @return ptr64: A pointer to the entity instance, or nullptr if the entity does not exist.
   */
  inline void* EntIndexToEntPointer(int32_t entityIndex) {
    using EntIndexToEntPointerFn = void* (*)(int32_t);
    static EntIndexToEntPointerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.EntIndexToEntPointer", reinterpret_cast<void**>(&__func));
    return __func(entityIndex);
  }

  /**
   * @brief Retrieves the entity index from an entity pointer.
   * @function EntPointerToEntIndex
   * @param entity (ptr64): A pointer to the entity whose index is to be retrieved.
   * @return int32: The index of the entity, or -1 if the entity is nullptr.
   */
  inline int32_t EntPointerToEntIndex(void* entity) {
    using EntPointerToEntIndexFn = int32_t (*)(void*);
    static EntPointerToEntIndexFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.EntPointerToEntIndex", reinterpret_cast<void**>(&__func));
    return __func(entity);
  }

  /**
   * @brief Converts an entity pointer into an entity handle.
   * @function EntPointerToEntHandle
   * @param entity (ptr64): A pointer to the entity to convert.
   * @return int32: The entity handle as an integer, or INVALID_EHANDLE_INDEX if the entity is nullptr.
   */
  inline int32_t EntPointerToEntHandle(void* entity) {
    using EntPointerToEntHandleFn = int32_t (*)(void*);
    static EntPointerToEntHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.EntPointerToEntHandle", reinterpret_cast<void**>(&__func));
    return __func(entity);
  }

  /**
   * @brief Retrieves the entity pointer from an entity handle.
   * @function EntHandleToEntPointer
   * @param entityHandle (int32): The entity handle to convert.
   * @return ptr64: A pointer to the entity instance, or nullptr if the handle is invalid.
   */
  inline void* EntHandleToEntPointer(int32_t entityHandle) {
    using EntHandleToEntPointerFn = void* (*)(int32_t);
    static EntHandleToEntPointerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.EntHandleToEntPointer", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Converts an entity index into an entity handle.
   * @function EntIndexToEntHandle
   * @param entityIndex (int32): The index of the entity to convert.
   * @return int32: The entity handle as an integer, or -1 if the entity index is invalid.
   */
  inline int32_t EntIndexToEntHandle(int32_t entityIndex) {
    using EntIndexToEntHandleFn = int32_t (*)(int32_t);
    static EntIndexToEntHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.EntIndexToEntHandle", reinterpret_cast<void**>(&__func));
    return __func(entityIndex);
  }

  /**
   * @brief Retrieves the entity index from an entity handle.
   * @function EntHandleToEntIndex
   * @param entityHandle (int32): The entity handle from which to retrieve the index.
   * @return int32: The index of the entity, or -1 if the handle is invalid.
   */
  inline int32_t EntHandleToEntIndex(int32_t entityHandle) {
    using EntHandleToEntIndexFn = int32_t (*)(int32_t);
    static EntHandleToEntIndexFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.EntHandleToEntIndex", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Checks if the provided entity handle is valid.
   * @function IsValidEntHandle
   * @param entityHandle (int32): The entity handle to check.
   * @return bool: True if the entity handle is valid, false otherwise.
   */
  inline bool IsValidEntHandle(int32_t entityHandle) {
    using IsValidEntHandleFn = bool (*)(int32_t);
    static IsValidEntHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsValidEntHandle", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Checks if the provided entity pointer is valid.
   * @function IsValidEntPointer
   * @param entity (ptr64): The entity pointer to check.
   * @return bool: True if the entity pointer is valid, false otherwise.
   */
  inline bool IsValidEntPointer(void* entity) {
    using IsValidEntPointerFn = bool (*)(void*);
    static IsValidEntPointerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsValidEntPointer", reinterpret_cast<void**>(&__func));
    return __func(entity);
  }

  /**
   * @brief Retrieves the pointer to the first active entity.
   * @function GetFirstActiveEntity
   * @return ptr64: A pointer to the first active entity.
   */
  inline void* GetFirstActiveEntity() {
    using GetFirstActiveEntityFn = void* (*)();
    static GetFirstActiveEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetFirstActiveEntity", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Retrieves a pointer to the concrete entity list.
   * @function GetConcreteEntityListPointer
   * @return ptr64: A pointer to the entity list structure.
   */
  inline void* GetConcreteEntityListPointer() {
    using GetConcreteEntityListPointerFn = void* (*)();
    static GetConcreteEntityListPointerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetConcreteEntityListPointer", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Adds an entity output hook on a specified entity class name.
   * @function HookEntityOutput
   * @param classname (string): The class name of the entity to hook the output for.
   * @param output (string): The output event name to hook.
   * @param callback (function): The callback function to invoke when the output is fired.
   * @param type (uint8): Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @return bool: True if the hook was successfully added, false otherwise.
   */
  inline bool HookEntityOutput(plg::string classname, plg::string output, function callback, uint8_t type) {
    using HookEntityOutputFn = bool (*)(plg::string, plg::string, function, uint8_t);
    static HookEntityOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.HookEntityOutput", reinterpret_cast<void**>(&__func));
    return __func(classname, output, callback, type);
  }

  /**
   * @brief Removes an entity output hook.
   * @function UnhookEntityOutput
   * @param classname (string): The class name of the entity from which to unhook the output.
   * @param output (string): The output event name to unhook.
   * @param callback (function): The callback function that was previously hooked.
   * @param type (uint8): Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @return bool: True if the hook was successfully removed, false otherwise.
   */
  inline bool UnhookEntityOutput(plg::string classname, plg::string output, function callback, uint8_t type) {
    using UnhookEntityOutputFn = bool (*)(plg::string, plg::string, function, uint8_t);
    static UnhookEntityOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UnhookEntityOutput", reinterpret_cast<void**>(&__func));
    return __func(classname, output, callback, type);
  }

  /**
   * @brief Finds an entity by classname within a radius with iteration.
   * @function FindEntityByClassnameWithin
   * @param startFrom (int32): The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param classname (string): The class name to search for.
   * @param origin (vec3): The center of the search sphere.
   * @param radius (float): The search radius.
   * @return int32: The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  inline int32_t FindEntityByClassnameWithin(int32_t startFrom, plg::string classname, plg::vec3 origin, float radius) {
    using FindEntityByClassnameWithinFn = int32_t (*)(int32_t, plg::string, plg::vec3, float);
    static FindEntityByClassnameWithinFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindEntityByClassnameWithin", reinterpret_cast<void**>(&__func));
    return __func(startFrom, classname, origin, radius);
  }

  /**
   * @brief Finds an entity by name with iteration.
   * @function FindEntityByName
   * @param startFrom (int32): The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param name (string): The targetname to search for.
   * @return int32: The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  inline int32_t FindEntityByName(int32_t startFrom, plg::string name) {
    using FindEntityByNameFn = int32_t (*)(int32_t, plg::string);
    static FindEntityByNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindEntityByName", reinterpret_cast<void**>(&__func));
    return __func(startFrom, name);
  }

  /**
   * @brief Finds the nearest entity by name to a point.
   * @function FindEntityByNameNearest
   * @param name (string): The targetname to search for.
   * @param origin (vec3): The point to search around.
   * @param maxRadius (float): Maximum search radius.
   * @return int32: The handle of the nearest entity, or INVALID_EHANDLE_INDEX if none found.
   */
  inline int32_t FindEntityByNameNearest(plg::string name, plg::vec3 origin, float maxRadius) {
    using FindEntityByNameNearestFn = int32_t (*)(plg::string, plg::vec3, float);
    static FindEntityByNameNearestFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindEntityByNameNearest", reinterpret_cast<void**>(&__func));
    return __func(name, origin, maxRadius);
  }

  /**
   * @brief Finds an entity by name within a radius with iteration.
   * @function FindEntityByNameWithin
   * @param startFrom (int32): The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param name (string): The targetname to search for.
   * @param origin (vec3): The center of the search sphere.
   * @param radius (float): The search radius.
   * @return int32: The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  inline int32_t FindEntityByNameWithin(int32_t startFrom, plg::string name, plg::vec3 origin, float radius) {
    using FindEntityByNameWithinFn = int32_t (*)(int32_t, plg::string, plg::vec3, float);
    static FindEntityByNameWithinFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindEntityByNameWithin", reinterpret_cast<void**>(&__func));
    return __func(startFrom, name, origin, radius);
  }

  /**
   * @brief Finds an entity by targetname with iteration.
   * @function FindEntityByTarget
   * @param startFrom (int32): The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param name (string): The targetname to search for.
   * @return int32: The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  inline int32_t FindEntityByTarget(int32_t startFrom, plg::string name) {
    using FindEntityByTargetFn = int32_t (*)(int32_t, plg::string);
    static FindEntityByTargetFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindEntityByTarget", reinterpret_cast<void**>(&__func));
    return __func(startFrom, name);
  }

  /**
   * @brief Finds an entity within a sphere with iteration.
   * @function FindEntityInSphere
   * @param startFrom (int32): The handle of the entity to start from, or INVALID_EHANDLE_INDEX to start fresh.
   * @param origin (vec3): The center of the search sphere.
   * @param radius (float): The search radius.
   * @return int32: The handle of the found entity, or INVALID_EHANDLE_INDEX if none found.
   */
  inline int32_t FindEntityInSphere(int32_t startFrom, plg::vec3 origin, float radius) {
    using FindEntityInSphereFn = int32_t (*)(int32_t, plg::vec3, float);
    static FindEntityInSphereFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FindEntityInSphere", reinterpret_cast<void**>(&__func));
    return __func(startFrom, origin, radius);
  }

  /**
   * @brief Creates an entity by classname.
   * @function SpawnEntityByName
   * @param className (string): The class name of the entity to create.
   * @return int32: The handle of the created entity, or INVALID_EHANDLE_INDEX if creation failed.
   */
  inline int32_t SpawnEntityByName(plg::string className) {
    using SpawnEntityByNameFn = int32_t (*)(plg::string);
    static SpawnEntityByNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SpawnEntityByName", reinterpret_cast<void**>(&__func));
    return __func(className);
  }

  /**
   * @brief Creates an entity by string name but does not spawn it.
   * @function CreateEntityByName
   * @param className (string): The class name of the entity to create.
   * @return int32: The entity handle of the created entity, or INVALID_EHANDLE_INDEX if the entity could not be created.
   */
  inline int32_t CreateEntityByName(plg::string className) {
    using CreateEntityByNameFn = int32_t (*)(plg::string);
    static CreateEntityByNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateEntityByName", reinterpret_cast<void**>(&__func));
    return __func(className);
  }

  /**
   * @brief Spawns an entity into the game.
   * @function DispatchSpawn
   * @param entityHandle (int32): The handle of the entity to spawn.
   */
  inline void DispatchSpawn(int32_t entityHandle) {
    using DispatchSpawnFn = void (*)(int32_t);
    static DispatchSpawnFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DispatchSpawn", reinterpret_cast<void**>(&__func));
    __func(entityHandle);
  }

  /**
   * @brief Spawns an entity into the game with key-value properties.
   * @function DispatchSpawn2
   * @param entityHandle (int32): The handle of the entity to spawn.
   * @param keys (string[]): A vector of keys representing the property names to set on the entity.
   * @param values (any[]): A vector of values corresponding to the keys, representing the property values to set on the entity.
   */
  inline void DispatchSpawn2(int32_t entityHandle, plg::vector<plg::string> keys, plg::vector<plg::any> values) {
    using DispatchSpawn2Fn = void (*)(int32_t, plg::vector<plg::string>, plg::vector<plg::any>);
    static DispatchSpawn2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DispatchSpawn2", reinterpret_cast<void**>(&__func));
    __func(entityHandle, keys, values);
  }

  /**
   * @brief Marks an entity for deletion.
   * @function RemoveEntity
   * @param entityHandle (int32): The handle of the entity to be deleted.
   */
  inline void RemoveEntity(int32_t entityHandle) {
    using RemoveEntityFn = void (*)(int32_t);
    static RemoveEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RemoveEntity", reinterpret_cast<void**>(&__func));
    __func(entityHandle);
  }

  /**
   * @brief Checks if an entity is a player controller.
   * @function IsEntityPlayerController
   * @param entityHandle (int32): The handle of the entity.
   * @return bool: True if the entity is a player controller, false otherwise.
   */
  inline bool IsEntityPlayerController(int32_t entityHandle) {
    using IsEntityPlayerControllerFn = bool (*)(int32_t);
    static IsEntityPlayerControllerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsEntityPlayerController", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Checks if an entity is a player pawn.
   * @function IsEntityPlayerPawn
   * @param entityHandle (int32): The handle of the entity.
   * @return bool: True if the entity is a player pawn, false otherwise.
   */
  inline bool IsEntityPlayerPawn(int32_t entityHandle) {
    using IsEntityPlayerPawnFn = bool (*)(int32_t);
    static IsEntityPlayerPawnFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsEntityPlayerPawn", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the class name of an entity.
   * @function GetEntityClassname
   * @param entityHandle (int32): The handle of the entity whose class name is to be retrieved.
   * @return string: A string where the class name will be stored.
   */
  inline plg::string GetEntityClassname(int32_t entityHandle) {
    using GetEntityClassnameFn = plg::string (*)(int32_t);
    static GetEntityClassnameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityClassname", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the name of an entity.
   * @function GetEntityName
   * @param entityHandle (int32): The handle of the entity whose name is to be retrieved.
   * @return string
   */
  inline plg::string GetEntityName(int32_t entityHandle) {
    using GetEntityNameFn = plg::string (*)(int32_t);
    static GetEntityNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityName", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the name of an entity.
   * @function SetEntityName
   * @param entityHandle (int32): The handle of the entity whose name is to be set.
   * @param name (string): The new name to set for the entity.
   */
  inline void SetEntityName(int32_t entityHandle, plg::string name) {
    using SetEntityNameFn = void (*)(int32_t, plg::string);
    static SetEntityNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityName", reinterpret_cast<void**>(&__func));
    __func(entityHandle, name);
  }

  /**
   * @brief Retrieves the movement type of an entity.
   * @function GetEntityMoveType
   * @param entityHandle (int32): The handle of the entity whose movement type is to be retrieved.
   * @return int32: The movement type of the entity, or 0 if the entity is invalid.
   */
  inline int32_t GetEntityMoveType(int32_t entityHandle) {
    using GetEntityMoveTypeFn = int32_t (*)(int32_t);
    static GetEntityMoveTypeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityMoveType", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the movement type of an entity.
   * @function SetEntityMoveType
   * @param entityHandle (int32): The handle of the entity whose movement type is to be set.
   * @param moveType (int32): The new movement type to set for the entity.
   */
  inline void SetEntityMoveType(int32_t entityHandle, int32_t moveType) {
    using SetEntityMoveTypeFn = void (*)(int32_t, int32_t);
    static SetEntityMoveTypeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityMoveType", reinterpret_cast<void**>(&__func));
    __func(entityHandle, moveType);
  }

  /**
   * @brief Retrieves the gravity scale of an entity.
   * @function GetEntityGravity
   * @param entityHandle (int32): The handle of the entity whose gravity scale is to be retrieved.
   * @return float: The gravity scale of the entity, or 0.0f if the entity is invalid.
   */
  inline float GetEntityGravity(int32_t entityHandle) {
    using GetEntityGravityFn = float (*)(int32_t);
    static GetEntityGravityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityGravity", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the gravity scale of an entity.
   * @function SetEntityGravity
   * @param entityHandle (int32): The handle of the entity whose gravity scale is to be set.
   * @param gravity (float): The new gravity scale to set for the entity.
   */
  inline void SetEntityGravity(int32_t entityHandle, float gravity) {
    using SetEntityGravityFn = void (*)(int32_t, float);
    static SetEntityGravityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityGravity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, gravity);
  }

  /**
   * @brief Retrieves the flags of an entity.
   * @function GetEntityFlags
   * @param entityHandle (int32): The handle of the entity whose flags are to be retrieved.
   * @return int32: The flags of the entity, or 0 if the entity is invalid.
   */
  inline int32_t GetEntityFlags(int32_t entityHandle) {
    using GetEntityFlagsFn = int32_t (*)(int32_t);
    static GetEntityFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityFlags", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the flags of an entity.
   * @function SetEntityFlags
   * @param entityHandle (int32): The handle of the entity whose flags are to be set.
   * @param flags (int32): The new flags to set for the entity.
   */
  inline void SetEntityFlags(int32_t entityHandle, int32_t flags) {
    using SetEntityFlagsFn = void (*)(int32_t, int32_t);
    static SetEntityFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityFlags", reinterpret_cast<void**>(&__func));
    __func(entityHandle, flags);
  }

  /**
   * @brief Retrieves the render color of an entity.
   * @function GetEntityRenderColor
   * @param entityHandle (int32): The handle of the entity whose render color is to be retrieved.
   * @return int32: The raw color value of the entity's render color, or 0 if the entity is invalid.
   */
  inline int32_t GetEntityRenderColor(int32_t entityHandle) {
    using GetEntityRenderColorFn = int32_t (*)(int32_t);
    static GetEntityRenderColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityRenderColor", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the render color of an entity.
   * @function SetEntityRenderColor
   * @param entityHandle (int32): The handle of the entity whose render color is to be set.
   * @param color (int32): The new raw color value to set for the entity's render color.
   */
  inline void SetEntityRenderColor(int32_t entityHandle, int32_t color) {
    using SetEntityRenderColorFn = void (*)(int32_t, int32_t);
    static SetEntityRenderColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityRenderColor", reinterpret_cast<void**>(&__func));
    __func(entityHandle, color);
  }

  /**
   * @brief Retrieves the render mode of an entity.
   * @function GetEntityRenderMode
   * @param entityHandle (int32): The handle of the entity whose render mode is to be retrieved.
   * @return uint8: The render mode of the entity, or 0 if the entity is invalid.
   */
  inline uint8_t GetEntityRenderMode(int32_t entityHandle) {
    using GetEntityRenderModeFn = uint8_t (*)(int32_t);
    static GetEntityRenderModeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityRenderMode", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the render mode of an entity.
   * @function SetEntityRenderMode
   * @param entityHandle (int32): The handle of the entity whose render mode is to be set.
   * @param renderMode (uint8): The new render mode to set for the entity.
   */
  inline void SetEntityRenderMode(int32_t entityHandle, uint8_t renderMode) {
    using SetEntityRenderModeFn = void (*)(int32_t, uint8_t);
    static SetEntityRenderModeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityRenderMode", reinterpret_cast<void**>(&__func));
    __func(entityHandle, renderMode);
  }

  /**
   * @brief Retrieves the mass of an entity.
   * @function GetEntityMass
   * @param entityHandle (int32): The handle of the entity whose mass is to be retrieved.
   * @return int32: The mass of the entity, or 0 if the entity is invalid.
   */
  inline int32_t GetEntityMass(int32_t entityHandle) {
    using GetEntityMassFn = int32_t (*)(int32_t);
    static GetEntityMassFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityMass", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the mass of an entity.
   * @function SetEntityMass
   * @param entityHandle (int32): The handle of the entity whose mass is to be set.
   * @param mass (int32): The new mass value to set for the entity.
   */
  inline void SetEntityMass(int32_t entityHandle, int32_t mass) {
    using SetEntityMassFn = void (*)(int32_t, int32_t);
    static SetEntityMassFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityMass", reinterpret_cast<void**>(&__func));
    __func(entityHandle, mass);
  }

  /**
   * @brief Retrieves the friction of an entity.
   * @function GetEntityFriction
   * @param entityHandle (int32): The handle of the entity whose friction is to be retrieved.
   * @return float: The friction of the entity, or 0 if the entity is invalid.
   */
  inline float GetEntityFriction(int32_t entityHandle) {
    using GetEntityFrictionFn = float (*)(int32_t);
    static GetEntityFrictionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityFriction", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the friction of an entity.
   * @function SetEntityFriction
   * @param entityHandle (int32): The handle of the entity whose friction is to be set.
   * @param friction (float): The new friction value to set for the entity.
   */
  inline void SetEntityFriction(int32_t entityHandle, float friction) {
    using SetEntityFrictionFn = void (*)(int32_t, float);
    static SetEntityFrictionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityFriction", reinterpret_cast<void**>(&__func));
    __func(entityHandle, friction);
  }

  /**
   * @brief Retrieves the health of an entity.
   * @function GetEntityHealth
   * @param entityHandle (int32): The handle of the entity whose health is to be retrieved.
   * @return int32: The health of the entity, or 0 if the entity is invalid.
   */
  inline int32_t GetEntityHealth(int32_t entityHandle) {
    using GetEntityHealthFn = int32_t (*)(int32_t);
    static GetEntityHealthFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityHealth", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the health of an entity.
   * @function SetEntityHealth
   * @param entityHandle (int32): The handle of the entity whose health is to be set.
   * @param health (int32): The new health value to set for the entity.
   */
  inline void SetEntityHealth(int32_t entityHandle, int32_t health) {
    using SetEntityHealthFn = void (*)(int32_t, int32_t);
    static SetEntityHealthFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityHealth", reinterpret_cast<void**>(&__func));
    __func(entityHandle, health);
  }

  /**
   * @brief Retrieves the max health of an entity.
   * @function GetEntityMaxHealth
   * @param entityHandle (int32): The handle of the entity whose max health is to be retrieved.
   * @return int32: The max health of the entity, or 0 if the entity is invalid.
   */
  inline int32_t GetEntityMaxHealth(int32_t entityHandle) {
    using GetEntityMaxHealthFn = int32_t (*)(int32_t);
    static GetEntityMaxHealthFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityMaxHealth", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the max health of an entity.
   * @function SetEntityMaxHealth
   * @param entityHandle (int32): The handle of the entity whose max health is to be set.
   * @param maxHealth (int32): The new max health value to set for the entity.
   */
  inline void SetEntityMaxHealth(int32_t entityHandle, int32_t maxHealth) {
    using SetEntityMaxHealthFn = void (*)(int32_t, int32_t);
    static SetEntityMaxHealthFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityMaxHealth", reinterpret_cast<void**>(&__func));
    __func(entityHandle, maxHealth);
  }

  /**
   * @brief Retrieves the team number of an entity.
   * @function GetEntityTeam
   * @param entityHandle (int32): The handle of the entity whose team number is to be retrieved.
   * @return int32: The team number of the entity, or 0 if the entity is invalid.
   */
  inline int32_t GetEntityTeam(int32_t entityHandle) {
    using GetEntityTeamFn = int32_t (*)(int32_t);
    static GetEntityTeamFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityTeam", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the team number of an entity.
   * @function SetEntityTeam
   * @param entityHandle (int32): The handle of the entity whose team number is to be set.
   * @param team (int32): The new team number to set for the entity.
   */
  inline void SetEntityTeam(int32_t entityHandle, int32_t team) {
    using SetEntityTeamFn = void (*)(int32_t, int32_t);
    static SetEntityTeamFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityTeam", reinterpret_cast<void**>(&__func));
    __func(entityHandle, team);
  }

  /**
   * @brief Retrieves the owner of an entity.
   * @function GetEntityOwner
   * @param entityHandle (int32): The handle of the entity whose owner is to be retrieved.
   * @return int32: The handle of the owner entity, or INVALID_EHANDLE_INDEX if the entity is invalid.
   */
  inline int32_t GetEntityOwner(int32_t entityHandle) {
    using GetEntityOwnerFn = int32_t (*)(int32_t);
    static GetEntityOwnerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityOwner", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the owner of an entity.
   * @function SetEntityOwner
   * @param entityHandle (int32): The handle of the entity whose owner is to be set.
   * @param ownerHandle (int32): The handle of the new owner entity.
   */
  inline void SetEntityOwner(int32_t entityHandle, int32_t ownerHandle) {
    using SetEntityOwnerFn = void (*)(int32_t, int32_t);
    static SetEntityOwnerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityOwner", reinterpret_cast<void**>(&__func));
    __func(entityHandle, ownerHandle);
  }

  /**
   * @brief Retrieves the parent of an entity.
   * @function GetEntityParent
   * @param entityHandle (int32): The handle of the entity whose parent is to be retrieved.
   * @return int32: The handle of the parent entity, or INVALID_EHANDLE_INDEX if the entity is invalid.
   */
  inline int32_t GetEntityParent(int32_t entityHandle) {
    using GetEntityParentFn = int32_t (*)(int32_t);
    static GetEntityParentFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityParent", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the parent of an entity.
   * @function SetEntityParent
   * @param entityHandle (int32): The handle of the entity whose parent is to be set.
   * @param parentHandle (int32): The handle of the new parent entity.
   * @param attachmentName (string): The name of the entity's attachment.
   */
  inline void SetEntityParent(int32_t entityHandle, int32_t parentHandle, plg::string attachmentName) {
    using SetEntityParentFn = void (*)(int32_t, int32_t, plg::string);
    static SetEntityParentFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityParent", reinterpret_cast<void**>(&__func));
    __func(entityHandle, parentHandle, attachmentName);
  }

  /**
   * @brief Retrieves the absolute origin of an entity.
   * @function GetEntityAbsOrigin
   * @param entityHandle (int32): The handle of the entity whose absolute origin is to be retrieved.
   * @return vec3: A vector where the absolute origin will be stored.
   */
  inline plg::vec3 GetEntityAbsOrigin(int32_t entityHandle) {
    using GetEntityAbsOriginFn = plg::vec3 (*)(int32_t);
    static GetEntityAbsOriginFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAbsOrigin", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the absolute origin of an entity.
   * @function SetEntityAbsOrigin
   * @param entityHandle (int32): The handle of the entity whose absolute origin is to be set.
   * @param origin (vec3): The new absolute origin to set for the entity.
   */
  inline void SetEntityAbsOrigin(int32_t entityHandle, plg::vec3 origin) {
    using SetEntityAbsOriginFn = void (*)(int32_t, plg::vec3);
    static SetEntityAbsOriginFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityAbsOrigin", reinterpret_cast<void**>(&__func));
    __func(entityHandle, origin);
  }

  /**
   * @brief Retrieves the absolute scale of an entity.
   * @function GetEntityAbsScale
   * @param entityHandle (int32): The handle of the entity whose absolute scale is to be retrieved.
   * @return float: A vector where the absolute scale will be stored.
   */
  inline float GetEntityAbsScale(int32_t entityHandle) {
    using GetEntityAbsScaleFn = float (*)(int32_t);
    static GetEntityAbsScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAbsScale", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the absolute scale of an entity.
   * @function SetEntityAbsScale
   * @param entityHandle (int32): The handle of the entity whose absolute scale is to be set.
   * @param scale (float): The new absolute scale to set for the entity.
   */
  inline void SetEntityAbsScale(int32_t entityHandle, float scale) {
    using SetEntityAbsScaleFn = void (*)(int32_t, float);
    static SetEntityAbsScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityAbsScale", reinterpret_cast<void**>(&__func));
    __func(entityHandle, scale);
  }

  /**
   * @brief Retrieves the angular rotation of an entity.
   * @function GetEntityAbsAngles
   * @param entityHandle (int32): The handle of the entity whose angular rotation is to be retrieved.
   * @return vec3: A QAngle where the angular rotation will be stored.
   */
  inline plg::vec3 GetEntityAbsAngles(int32_t entityHandle) {
    using GetEntityAbsAnglesFn = plg::vec3 (*)(int32_t);
    static GetEntityAbsAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAbsAngles", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the angular rotation of an entity.
   * @function SetEntityAbsAngles
   * @param entityHandle (int32): The handle of the entity whose angular rotation is to be set.
   * @param angle (vec3): The new angular rotation to set for the entity.
   */
  inline void SetEntityAbsAngles(int32_t entityHandle, plg::vec3 angle) {
    using SetEntityAbsAnglesFn = void (*)(int32_t, plg::vec3);
    static SetEntityAbsAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityAbsAngles", reinterpret_cast<void**>(&__func));
    __func(entityHandle, angle);
  }

  /**
   * @brief Retrieves the local origin of an entity.
   * @function GetEntityLocalOrigin
   * @param entityHandle (int32): The handle of the entity whose local origin is to be retrieved.
   * @return vec3: A vector where the local origin will be stored.
   */
  inline plg::vec3 GetEntityLocalOrigin(int32_t entityHandle) {
    using GetEntityLocalOriginFn = plg::vec3 (*)(int32_t);
    static GetEntityLocalOriginFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityLocalOrigin", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the local origin of an entity.
   * @function SetEntityLocalOrigin
   * @param entityHandle (int32): The handle of the entity whose local origin is to be set.
   * @param origin (vec3): The new local origin to set for the entity.
   */
  inline void SetEntityLocalOrigin(int32_t entityHandle, plg::vec3 origin) {
    using SetEntityLocalOriginFn = void (*)(int32_t, plg::vec3);
    static SetEntityLocalOriginFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityLocalOrigin", reinterpret_cast<void**>(&__func));
    __func(entityHandle, origin);
  }

  /**
   * @brief Retrieves the local scale of an entity.
   * @function GetEntityLocalScale
   * @param entityHandle (int32): The handle of the entity whose local scale is to be retrieved.
   * @return float: A vector where the local scale will be stored.
   */
  inline float GetEntityLocalScale(int32_t entityHandle) {
    using GetEntityLocalScaleFn = float (*)(int32_t);
    static GetEntityLocalScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityLocalScale", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the local scale of an entity.
   * @function SetEntityLocalScale
   * @param entityHandle (int32): The handle of the entity whose local scale is to be set.
   * @param scale (float): The new local scale to set for the entity.
   */
  inline void SetEntityLocalScale(int32_t entityHandle, float scale) {
    using SetEntityLocalScaleFn = void (*)(int32_t, float);
    static SetEntityLocalScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityLocalScale", reinterpret_cast<void**>(&__func));
    __func(entityHandle, scale);
  }

  /**
   * @brief Retrieves the angular rotation of an entity.
   * @function GetEntityLocalAngles
   * @param entityHandle (int32): The handle of the entity whose angular rotation is to be retrieved.
   * @return vec3: A QAngle where the angular rotation will be stored.
   */
  inline plg::vec3 GetEntityLocalAngles(int32_t entityHandle) {
    using GetEntityLocalAnglesFn = plg::vec3 (*)(int32_t);
    static GetEntityLocalAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityLocalAngles", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the angular rotation of an entity.
   * @function SetEntityLocalAngles
   * @param entityHandle (int32): The handle of the entity whose angular rotation is to be set.
   * @param angle (vec3): The new angular rotation to set for the entity.
   */
  inline void SetEntityLocalAngles(int32_t entityHandle, plg::vec3 angle) {
    using SetEntityLocalAnglesFn = void (*)(int32_t, plg::vec3);
    static SetEntityLocalAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityLocalAngles", reinterpret_cast<void**>(&__func));
    __func(entityHandle, angle);
  }

  /**
   * @brief Retrieves the absolute velocity of an entity.
   * @function GetEntityAbsVelocity
   * @param entityHandle (int32): The handle of the entity whose absolute velocity is to be retrieved.
   * @return vec3: A vector where the absolute velocity will be stored.
   */
  inline plg::vec3 GetEntityAbsVelocity(int32_t entityHandle) {
    using GetEntityAbsVelocityFn = plg::vec3 (*)(int32_t);
    static GetEntityAbsVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAbsVelocity", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the absolute velocity of an entity.
   * @function SetEntityAbsVelocity
   * @param entityHandle (int32): The handle of the entity whose absolute velocity is to be set.
   * @param velocity (vec3): The new absolute velocity to set for the entity.
   */
  inline void SetEntityAbsVelocity(int32_t entityHandle, plg::vec3 velocity) {
    using SetEntityAbsVelocityFn = void (*)(int32_t, plg::vec3);
    static SetEntityAbsVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityAbsVelocity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, velocity);
  }

  /**
   * @brief Retrieves the base velocity of an entity.
   * @function GetEntityBaseVelocity
   * @param entityHandle (int32): The handle of the entity whose base velocity is to be retrieved.
   * @return vec3: A vector where the base velocity will be stored.
   */
  inline plg::vec3 GetEntityBaseVelocity(int32_t entityHandle) {
    using GetEntityBaseVelocityFn = plg::vec3 (*)(int32_t);
    static GetEntityBaseVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityBaseVelocity", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the local angular velocity of an entity.
   * @function GetEntityLocalAngVelocity
   * @param entityHandle (int32): The handle of the entity whose local angular velocity is to be retrieved.
   * @return vec3: A vector where the local angular velocity will be stored.
   */
  inline plg::vec3 GetEntityLocalAngVelocity(int32_t entityHandle) {
    using GetEntityLocalAngVelocityFn = plg::vec3 (*)(int32_t);
    static GetEntityLocalAngVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityLocalAngVelocity", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the angular velocity of an entity.
   * @function GetEntityAngVelocity
   * @param entityHandle (int32): The handle of the entity whose angular velocity is to be retrieved.
   * @return vec3: A vector where the angular velocity will be stored.
   */
  inline plg::vec3 GetEntityAngVelocity(int32_t entityHandle) {
    using GetEntityAngVelocityFn = plg::vec3 (*)(int32_t);
    static GetEntityAngVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAngVelocity", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the angular velocity of an entity.
   * @function SetEntityAngVelocity
   * @param entityHandle (int32): The handle of the entity whose angular velocity is to be set.
   * @param velocity (vec3): The new angular velocity to set for the entity.
   */
  inline void SetEntityAngVelocity(int32_t entityHandle, plg::vec3 velocity) {
    using SetEntityAngVelocityFn = void (*)(int32_t, plg::vec3);
    static SetEntityAngVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityAngVelocity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, velocity);
  }

  /**
   * @brief Retrieves the local velocity of an entity.
   * @function GetEntityLocalVelocity
   * @param entityHandle (int32): The handle of the entity whose local velocity is to be retrieved.
   * @return vec3: A vector where the local velocity will be stored.
   */
  inline plg::vec3 GetEntityLocalVelocity(int32_t entityHandle) {
    using GetEntityLocalVelocityFn = plg::vec3 (*)(int32_t);
    static GetEntityLocalVelocityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityLocalVelocity", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the angular rotation of an entity.
   * @function GetEntityAngRotation
   * @param entityHandle (int32): The handle of the entity whose angular rotation is to be retrieved.
   * @return vec3: A vector where the angular rotation will be stored.
   */
  inline plg::vec3 GetEntityAngRotation(int32_t entityHandle) {
    using GetEntityAngRotationFn = plg::vec3 (*)(int32_t);
    static GetEntityAngRotationFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAngRotation", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the angular rotation of an entity.
   * @function SetEntityAngRotation
   * @param entityHandle (int32): The handle of the entity whose angular rotation is to be set.
   * @param rotation (vec3): The new angular rotation to set for the entity.
   */
  inline void SetEntityAngRotation(int32_t entityHandle, plg::vec3 rotation) {
    using SetEntityAngRotationFn = void (*)(int32_t, plg::vec3);
    static SetEntityAngRotationFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityAngRotation", reinterpret_cast<void**>(&__func));
    __func(entityHandle, rotation);
  }

  /**
   * @brief Returns the input Vector transformed from entity to world space.
   * @function TransformPointEntityToWorld
   * @param entityHandle (int32): The handle of the entity
   * @param point (vec3): Point in entity local space to transform
   * @return vec3: The point transformed to world space coordinates
   */
  inline plg::vec3 TransformPointEntityToWorld(int32_t entityHandle, plg::vec3 point) {
    using TransformPointEntityToWorldFn = plg::vec3 (*)(int32_t, plg::vec3);
    static TransformPointEntityToWorldFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TransformPointEntityToWorld", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, point);
  }

  /**
   * @brief Returns the input Vector transformed from world to entity space.
   * @function TransformPointWorldToEntity
   * @param entityHandle (int32): The handle of the entity
   * @param point (vec3): Point in world space to transform
   * @return vec3: The point transformed to entity local space coordinates
   */
  inline plg::vec3 TransformPointWorldToEntity(int32_t entityHandle, plg::vec3 point) {
    using TransformPointWorldToEntityFn = plg::vec3 (*)(int32_t, plg::vec3);
    static TransformPointWorldToEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TransformPointWorldToEntity", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, point);
  }

  /**
   * @brief Get vector to eye position - absolute coords.
   * @function GetEntityEyePosition
   * @param entityHandle (int32): The handle of the entity
   * @return vec3: Eye position in absolute/world coordinates
   */
  inline plg::vec3 GetEntityEyePosition(int32_t entityHandle) {
    using GetEntityEyePositionFn = plg::vec3 (*)(int32_t);
    static GetEntityEyePositionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityEyePosition", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Get the qangles that this entity is looking at.
   * @function GetEntityEyeAngles
   * @param entityHandle (int32): The handle of the entity
   * @return vec3: Eye angles as a vector (pitch, yaw, roll)
   */
  inline plg::vec3 GetEntityEyeAngles(int32_t entityHandle) {
    using GetEntityEyeAnglesFn = plg::vec3 (*)(int32_t);
    static GetEntityEyeAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityEyeAngles", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the forward velocity of an entity.
   * @function SetEntityForwardVector
   * @param entityHandle (int32): The handle of the entity whose forward velocity is to be set.
   * @param forward (vec3)
   */
  inline void SetEntityForwardVector(int32_t entityHandle, plg::vec3 forward) {
    using SetEntityForwardVectorFn = void (*)(int32_t, plg::vec3);
    static SetEntityForwardVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityForwardVector", reinterpret_cast<void**>(&__func));
    __func(entityHandle, forward);
  }

  /**
   * @brief Get the forward vector of the entity.
   * @function GetEntityForwardVector
   * @param entityHandle (int32): The handle of the entity to query
   * @return vec3: Forward-facing direction vector of the entity
   */
  inline plg::vec3 GetEntityForwardVector(int32_t entityHandle) {
    using GetEntityForwardVectorFn = plg::vec3 (*)(int32_t);
    static GetEntityForwardVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityForwardVector", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Get the left vector of the entity.
   * @function GetEntityLeftVector
   * @param entityHandle (int32): The handle of the entity to query
   * @return vec3: Left-facing direction vector of the entity (aligned with the y axis)
   */
  inline plg::vec3 GetEntityLeftVector(int32_t entityHandle) {
    using GetEntityLeftVectorFn = plg::vec3 (*)(int32_t);
    static GetEntityLeftVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityLeftVector", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Get the right vector of the entity.
   * @function GetEntityRightVector
   * @param entityHandle (int32): The handle of the entity to query
   * @return vec3: Right-facing direction vector of the entity
   */
  inline plg::vec3 GetEntityRightVector(int32_t entityHandle) {
    using GetEntityRightVectorFn = plg::vec3 (*)(int32_t);
    static GetEntityRightVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityRightVector", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Get the up vector of the entity.
   * @function GetEntityUpVector
   * @param entityHandle (int32): The handle of the entity to query
   * @return vec3: Up-facing direction vector of the entity
   */
  inline plg::vec3 GetEntityUpVector(int32_t entityHandle) {
    using GetEntityUpVectorFn = plg::vec3 (*)(int32_t);
    static GetEntityUpVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityUpVector", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Get the entity-to-world transformation matrix.
   * @function GetEntityTransform
   * @param entityHandle (int32): The handle of the entity to query
   * @return mat4x4: 4x4 transformation matrix representing entity's position, rotation, and scale in world space
   */
  inline plg::mat4x4 GetEntityTransform(int32_t entityHandle) {
    using GetEntityTransformFn = plg::mat4x4 (*)(int32_t);
    static GetEntityTransformFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityTransform", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the model name of an entity.
   * @function GetEntityModel
   * @param entityHandle (int32): The handle of the entity whose model name is to be retrieved.
   * @return string: A string where the model name will be stored.
   */
  inline plg::string GetEntityModel(int32_t entityHandle) {
    using GetEntityModelFn = plg::string (*)(int32_t);
    static GetEntityModelFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityModel", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Sets the model name of an entity.
   * @function SetEntityModel
   * @param entityHandle (int32): The handle of the entity whose model name is to be set.
   * @param model (string): The new model name to set for the entity.
   */
  inline void SetEntityModel(int32_t entityHandle, plg::string model) {
    using SetEntityModelFn = void (*)(int32_t, plg::string);
    static SetEntityModelFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityModel", reinterpret_cast<void**>(&__func));
    __func(entityHandle, model);
  }

  /**
   * @brief Retrieves the water level of an entity.
   * @function GetEntityWaterLevel
   * @param entityHandle (int32): The handle of the entity whose water level is to be retrieved.
   * @return float: The water level of the entity, or 0.0f if the entity is invalid.
   */
  inline float GetEntityWaterLevel(int32_t entityHandle) {
    using GetEntityWaterLevelFn = float (*)(int32_t);
    static GetEntityWaterLevelFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityWaterLevel", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the ground entity of an entity.
   * @function GetEntityGroundEntity
   * @param entityHandle (int32): The handle of the entity whose ground entity is to be retrieved.
   * @return int32: The handle of the ground entity, or INVALID_EHANDLE_INDEX if the entity is invalid.
   */
  inline int32_t GetEntityGroundEntity(int32_t entityHandle) {
    using GetEntityGroundEntityFn = int32_t (*)(int32_t);
    static GetEntityGroundEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityGroundEntity", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the effects of an entity.
   * @function GetEntityEffects
   * @param entityHandle (int32): The handle of the entity whose effects are to be retrieved.
   * @return int32: The effect flags of the entity, or 0 if the entity is invalid.
   */
  inline int32_t GetEntityEffects(int32_t entityHandle) {
    using GetEntityEffectsFn = int32_t (*)(int32_t);
    static GetEntityEffectsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityEffects", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Adds the render effect flag to an entity.
   * @function AddEntityEffects
   * @param entityHandle (int32): The handle of the entity to modify
   * @param effects (int32): Render effect flags to add
   */
  inline void AddEntityEffects(int32_t entityHandle, int32_t effects) {
    using AddEntityEffectsFn = void (*)(int32_t, int32_t);
    static AddEntityEffectsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddEntityEffects", reinterpret_cast<void**>(&__func));
    __func(entityHandle, effects);
  }

  /**
   * @brief Removes the render effect flag from an entity.
   * @function RemoveEntityEffects
   * @param entityHandle (int32): The handle of the entity to modify
   * @param effects (int32): Render effect flags to remove
   */
  inline void RemoveEntityEffects(int32_t entityHandle, int32_t effects) {
    using RemoveEntityEffectsFn = void (*)(int32_t, int32_t);
    static RemoveEntityEffectsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RemoveEntityEffects", reinterpret_cast<void**>(&__func));
    __func(entityHandle, effects);
  }

  /**
   * @brief Get a vector containing max bounds, centered on object.
   * @function GetEntityBoundingMaxs
   * @param entityHandle (int32): The handle of the entity to query
   * @return vec3: Vector containing the maximum bounds of the entity's bounding box
   */
  inline plg::vec3 GetEntityBoundingMaxs(int32_t entityHandle) {
    using GetEntityBoundingMaxsFn = plg::vec3 (*)(int32_t);
    static GetEntityBoundingMaxsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityBoundingMaxs", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Get a vector containing min bounds, centered on object.
   * @function GetEntityBoundingMins
   * @param entityHandle (int32): The handle of the entity to query
   * @return vec3: Vector containing the minimum bounds of the entity's bounding box
   */
  inline plg::vec3 GetEntityBoundingMins(int32_t entityHandle) {
    using GetEntityBoundingMinsFn = plg::vec3 (*)(int32_t);
    static GetEntityBoundingMinsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityBoundingMins", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Get vector to center of object - absolute coords.
   * @function GetEntityCenter
   * @param entityHandle (int32): The handle of the entity to query
   * @return vec3: Vector pointing to the center of the entity in absolute/world coordinates
   */
  inline plg::vec3 GetEntityCenter(int32_t entityHandle) {
    using GetEntityCenterFn = plg::vec3 (*)(int32_t);
    static GetEntityCenterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityCenter", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Teleports an entity to a specified location and orientation.
   * @function TeleportEntity
   * @param entityHandle (int32): The handle of the entity to teleport.
   * @param origin (ptr64): A pointer to a Vector representing the new absolute position. Can be nullptr.
   * @param angles (ptr64): A pointer to a QAngle representing the new orientation. Can be nullptr.
   * @param velocity (ptr64): A pointer to a Vector representing the new velocity. Can be nullptr.
   */
  inline void TeleportEntity(int32_t entityHandle, void* origin, void* angles, void* velocity) {
    using TeleportEntityFn = void (*)(int32_t, void*, void*, void*);
    static TeleportEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TeleportEntity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, origin, angles, velocity);
  }

  /**
   * @brief Apply an absolute velocity impulse to an entity.
   * @function ApplyAbsVelocityImpulseToEntity
   * @param entityHandle (int32): The handle of the entity to apply impulse to
   * @param vecImpulse (vec3): Velocity impulse vector to apply
   */
  inline void ApplyAbsVelocityImpulseToEntity(int32_t entityHandle, plg::vec3 vecImpulse) {
    using ApplyAbsVelocityImpulseToEntityFn = void (*)(int32_t, plg::vec3);
    static ApplyAbsVelocityImpulseToEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ApplyAbsVelocityImpulseToEntity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, vecImpulse);
  }

  /**
   * @brief Apply a local angular velocity impulse to an entity.
   * @function ApplyLocalAngularVelocityImpulseToEntity
   * @param entityHandle (int32): The handle of the entity to apply impulse to
   * @param angImpulse (vec3): Angular velocity impulse vector to apply
   */
  inline void ApplyLocalAngularVelocityImpulseToEntity(int32_t entityHandle, plg::vec3 angImpulse) {
    using ApplyLocalAngularVelocityImpulseToEntityFn = void (*)(int32_t, plg::vec3);
    static ApplyLocalAngularVelocityImpulseToEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ApplyLocalAngularVelocityImpulseToEntity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, angImpulse);
  }

  /**
   * @brief Invokes a named input method on a specified entity.
   * @function AcceptEntityInput
   * @param entityHandle (int32): The handle of the target entity that will receive the input.
   * @param inputName (string): The name of the input action to invoke.
   * @param activatorHandle (int32): The handle of the entity that initiated the sequence of actions.
   * @param callerHandle (int32): The handle of the entity sending this event.
   * @param value (any): The value associated with the input action.
   * @param type (int32): The type or classification of the value.
   * @param outputId (int32): An identifier for tracking the output of this operation.
   */
  inline void AcceptEntityInput(int32_t entityHandle, plg::string inputName, int32_t activatorHandle, int32_t callerHandle, plg::any value, int32_t type, int32_t outputId) {
    using AcceptEntityInputFn = void (*)(int32_t, plg::string, int32_t, int32_t, plg::any, int32_t, int32_t);
    static AcceptEntityInputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AcceptEntityInput", reinterpret_cast<void**>(&__func));
    __func(entityHandle, inputName, activatorHandle, callerHandle, value, type, outputId);
  }

  /**
   * @brief Connects a script function to an entity output.
   * @function ConnectEntityOutput
   * @param entityHandle (int32): The handle of the entity.
   * @param output (string): The name of the output to connect to.
   * @param functionName (string): The name of the script function to call.
   */
  inline void ConnectEntityOutput(int32_t entityHandle, plg::string output, plg::string functionName) {
    using ConnectEntityOutputFn = void (*)(int32_t, plg::string, plg::string);
    static ConnectEntityOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ConnectEntityOutput", reinterpret_cast<void**>(&__func));
    __func(entityHandle, output, functionName);
  }

  /**
   * @brief Disconnects a script function from an entity output.
   * @function DisconnectEntityOutput
   * @param entityHandle (int32): The handle of the entity.
   * @param output (string): The name of the output.
   * @param functionName (string): The name of the script function to disconnect.
   */
  inline void DisconnectEntityOutput(int32_t entityHandle, plg::string output, plg::string functionName) {
    using DisconnectEntityOutputFn = void (*)(int32_t, plg::string, plg::string);
    static DisconnectEntityOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DisconnectEntityOutput", reinterpret_cast<void**>(&__func));
    __func(entityHandle, output, functionName);
  }

  /**
   * @brief Disconnects a script function from an I/O event on a different entity.
   * @function DisconnectEntityRedirectedOutput
   * @param entityHandle (int32): The handle of the calling entity.
   * @param output (string): The name of the output.
   * @param functionName (string): The function name to disconnect.
   * @param targetHandle (int32): The handle of the entity whose output is being disconnected.
   */
  inline void DisconnectEntityRedirectedOutput(int32_t entityHandle, plg::string output, plg::string functionName, int32_t targetHandle) {
    using DisconnectEntityRedirectedOutputFn = void (*)(int32_t, plg::string, plg::string, int32_t);
    static DisconnectEntityRedirectedOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DisconnectEntityRedirectedOutput", reinterpret_cast<void**>(&__func));
    __func(entityHandle, output, functionName, targetHandle);
  }

  /**
   * @brief Fires an entity output.
   * @function FireEntityOutput
   * @param entityHandle (int32): The handle of the entity firing the output.
   * @param outputName (string): The name of the output to fire.
   * @param activatorHandle (int32): The entity activating the output.
   * @param callerHandle (int32): The entity that called the output.
   * @param value (any): The value associated with the input action.
   * @param type (int32): The type or classification of the value.
   * @param delay (float): Delay in seconds before firing the output.
   */
  inline void FireEntityOutput(int32_t entityHandle, plg::string outputName, int32_t activatorHandle, int32_t callerHandle, plg::any value, int32_t type, float delay) {
    using FireEntityOutputFn = void (*)(int32_t, plg::string, int32_t, int32_t, plg::any, int32_t, float);
    static FireEntityOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FireEntityOutput", reinterpret_cast<void**>(&__func));
    __func(entityHandle, outputName, activatorHandle, callerHandle, value, type, delay);
  }

  /**
   * @brief Redirects an entity output to call a function on another entity.
   * @function RedirectEntityOutput
   * @param entityHandle (int32): The handle of the entity whose output is being redirected.
   * @param output (string): The name of the output to redirect.
   * @param functionName (string): The function name to call on the target entity.
   * @param targetHandle (int32): The handle of the entity that will receive the output call.
   */
  inline void RedirectEntityOutput(int32_t entityHandle, plg::string output, plg::string functionName, int32_t targetHandle) {
    using RedirectEntityOutputFn = void (*)(int32_t, plg::string, plg::string, int32_t);
    static RedirectEntityOutputFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RedirectEntityOutput", reinterpret_cast<void**>(&__func));
    __func(entityHandle, output, functionName, targetHandle);
  }

  /**
   * @brief Makes an entity follow another entity with optional bone merging.
   * @function FollowEntity
   * @param entityHandle (int32): The handle of the entity that will follow
   * @param attachmentHandle (int32): The handle of the entity to follow
   * @param boneMerge (bool): If true, bones will be merged between entities
   */
  inline void FollowEntity(int32_t entityHandle, int32_t attachmentHandle, bool boneMerge) {
    using FollowEntityFn = void (*)(int32_t, int32_t, bool);
    static FollowEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FollowEntity", reinterpret_cast<void**>(&__func));
    __func(entityHandle, attachmentHandle, boneMerge);
  }

  /**
   * @brief Makes an entity follow another entity and merge with a specific bone or attachment.
   * @function FollowEntityMerge
   * @param entityHandle (int32): The handle of the entity that will follow
   * @param attachmentHandle (int32): The handle of the entity to follow
   * @param boneOrAttachName (string): Name of the bone or attachment point to merge with
   */
  inline void FollowEntityMerge(int32_t entityHandle, int32_t attachmentHandle, plg::string boneOrAttachName) {
    using FollowEntityMergeFn = void (*)(int32_t, int32_t, plg::string);
    static FollowEntityMergeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FollowEntityMerge", reinterpret_cast<void**>(&__func));
    __func(entityHandle, attachmentHandle, boneOrAttachName);
  }

  /**
   * @brief Apply damage to an entity.
   * @function TakeEntityDamage
   * @param entityHandle (int32): The handle of the entity receiving damage
   * @param inflictorHandle (int32): The handle of the entity inflicting damage (e.g., projectile)
   * @param attackerHandle (int32): The handle of the attacking entity
   * @param force (vec3): Direction and magnitude of force to apply
   * @param hitPos (vec3): Position where the damage hit occurred
   * @param damage (float): Amount of damage to apply
   * @param damageTypes (int32): Bitfield of damage type flags
   * @return int32: Amount of damage actually applied to the entity
   */
  inline int32_t TakeEntityDamage(int32_t entityHandle, int32_t inflictorHandle, int32_t attackerHandle, plg::vec3 force, plg::vec3 hitPos, float damage, int32_t damageTypes) {
    using TakeEntityDamageFn = int32_t (*)(int32_t, int32_t, int32_t, plg::vec3, plg::vec3, float, int32_t);
    static TakeEntityDamageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TakeEntityDamage", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, inflictorHandle, attackerHandle, force, hitPos, damage, damageTypes);
  }

  /**
   * @brief Retrieves a float attribute value from an entity.
   * @function GetEntityAttributeFloatValue
   * @param entityHandle (int32): The handle of the entity.
   * @param name (string): The name of the attribute.
   * @param defaultValue (float): The default value to return if the attribute does not exist.
   * @return float: The float value of the attribute, or the default value if missing or invalid.
   */
  inline float GetEntityAttributeFloatValue(int32_t entityHandle, plg::string name, float defaultValue) {
    using GetEntityAttributeFloatValueFn = float (*)(int32_t, plg::string, float);
    static GetEntityAttributeFloatValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAttributeFloatValue", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, name, defaultValue);
  }

  /**
   * @brief Retrieves an integer attribute value from an entity.
   * @function GetEntityAttributeIntValue
   * @param entityHandle (int32): The handle of the entity.
   * @param name (string): The name of the attribute.
   * @param defaultValue (int32): The default value to return if the attribute does not exist.
   * @return int32: The integer value of the attribute, or the default value if missing or invalid.
   */
  inline int32_t GetEntityAttributeIntValue(int32_t entityHandle, plg::string name, int32_t defaultValue) {
    using GetEntityAttributeIntValueFn = int32_t (*)(int32_t, plg::string, int32_t);
    static GetEntityAttributeIntValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAttributeIntValue", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, name, defaultValue);
  }

  /**
   * @brief Sets a float attribute value on an entity.
   * @function SetEntityAttributeFloatValue
   * @param entityHandle (int32): The handle of the entity.
   * @param name (string): The name of the attribute.
   * @param value (float): The float value to assign to the attribute.
   */
  inline void SetEntityAttributeFloatValue(int32_t entityHandle, plg::string name, float value) {
    using SetEntityAttributeFloatValueFn = void (*)(int32_t, plg::string, float);
    static SetEntityAttributeFloatValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityAttributeFloatValue", reinterpret_cast<void**>(&__func));
    __func(entityHandle, name, value);
  }

  /**
   * @brief Sets an integer attribute value on an entity.
   * @function SetEntityAttributeIntValue
   * @param entityHandle (int32): The handle of the entity.
   * @param name (string): The name of the attribute.
   * @param value (int32): The integer value to assign to the attribute.
   */
  inline void SetEntityAttributeIntValue(int32_t entityHandle, plg::string name, int32_t value) {
    using SetEntityAttributeIntValueFn = void (*)(int32_t, plg::string, int32_t);
    static SetEntityAttributeIntValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityAttributeIntValue", reinterpret_cast<void**>(&__func));
    __func(entityHandle, name, value);
  }

  /**
   * @brief Deletes an attribute from an entity.
   * @function DeleteEntityAttribute
   * @param entityHandle (int32): The handle of the entity.
   * @param name (string): The name of the attribute to delete.
   */
  inline void DeleteEntityAttribute(int32_t entityHandle, plg::string name) {
    using DeleteEntityAttributeFn = void (*)(int32_t, plg::string);
    static DeleteEntityAttributeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.DeleteEntityAttribute", reinterpret_cast<void**>(&__func));
    __func(entityHandle, name);
  }

  /**
   * @brief Checks if an entity has a specific attribute.
   * @function HasEntityAttribute
   * @param entityHandle (int32): The handle of the entity.
   * @param name (string): The name of the attribute to check.
   * @return bool: True if the attribute exists, false otherwise.
   */
  inline bool HasEntityAttribute(int32_t entityHandle, plg::string name) {
    using HasEntityAttributeFn = bool (*)(int32_t, plg::string);
    static HasEntityAttributeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.HasEntityAttribute", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, name);
  }

  /**
   * @brief Creates a hook for when a game event is fired.
   * @function HookEvent
   * @param name (string): The name of the event to hook.
   * @param callback (function): The callback function to call when the event is fired.
   * @param type (uint8): Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @return int32: An integer indicating the result of the hook operation.
   */
  inline int32_t HookEvent(plg::string name, function callback, uint8_t type) {
    using HookEventFn = int32_t (*)(plg::string, function, uint8_t);
    static HookEventFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.HookEvent", reinterpret_cast<void**>(&__func));
    return __func(name, callback, type);
  }

  /**
   * @brief Removes a hook for when a game event is fired.
   * @function UnhookEvent
   * @param name (string): The name of the event to unhook.
   * @param callback (function): The callback function to remove.
   * @param type (uint8): Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @return int32: An integer indicating the result of the unhook operation.
   */
  inline int32_t UnhookEvent(plg::string name, function callback, uint8_t type) {
    using UnhookEventFn = int32_t (*)(plg::string, function, uint8_t);
    static UnhookEventFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UnhookEvent", reinterpret_cast<void**>(&__func));
    return __func(name, callback, type);
  }

  /**
   * @brief Creates a game event to be fired later.
   * @function CreateEvent
   * @param name (string): The name of the event to create.
   * @param force (bool): A boolean indicating whether to force the creation of the event.
   * @return ptr64: A pointer to the created EventInfo structure.
   */
  inline void* CreateEvent(plg::string name, bool force) {
    using CreateEventFn = void* (*)(plg::string, bool);
    static CreateEventFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateEvent", reinterpret_cast<void**>(&__func));
    return __func(name, force);
  }

  /**
   * @brief Fires a game event.
   * @function FireEvent
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param dontBroadcast (bool): A boolean indicating whether to broadcast the event.
   */
  inline void FireEvent(void* info, bool dontBroadcast) {
    using FireEventFn = void (*)(void*, bool);
    static FireEventFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FireEvent", reinterpret_cast<void**>(&__func));
    __func(info, dontBroadcast);
  }

  /**
   * @brief Fires a game event to a specific client.
   * @function FireEventToClient
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param playerSlot (int32): The index of the client to fire the event to.
   */
  inline void FireEventToClient(void* info, int32_t playerSlot) {
    using FireEventToClientFn = void (*)(void*, int32_t);
    static FireEventToClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.FireEventToClient", reinterpret_cast<void**>(&__func));
    __func(info, playerSlot);
  }

  /**
   * @brief Cancels a previously created game event that has not been fired.
   * @function CancelCreatedEvent
   * @param info (ptr64): A pointer to the EventInfo structure of the event to cancel.
   */
  inline void CancelCreatedEvent(void* info) {
    using CancelCreatedEventFn = void (*)(void*);
    static CancelCreatedEventFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CancelCreatedEvent", reinterpret_cast<void**>(&__func));
    __func(info);
  }

  /**
   * @brief Retrieves the boolean value of a game event's key.
   * @function GetEventBool
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the boolean value.
   * @return bool: The boolean value associated with the key.
   */
  inline bool GetEventBool(void* info, plg::string key) {
    using GetEventBoolFn = bool (*)(void*, plg::string);
    static GetEventBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventBool", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the float value of a game event's key.
   * @function GetEventFloat
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the float value.
   * @return float: The float value associated with the key.
   */
  inline float GetEventFloat(void* info, plg::string key) {
    using GetEventFloatFn = float (*)(void*, plg::string);
    static GetEventFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventFloat", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the integer value of a game event's key.
   * @function GetEventInt
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the integer value.
   * @return int32: The integer value associated with the key.
   */
  inline int32_t GetEventInt(void* info, plg::string key) {
    using GetEventIntFn = int32_t (*)(void*, plg::string);
    static GetEventIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventInt", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the long integer value of a game event's key.
   * @function GetEventUInt64
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the long integer value.
   * @return uint64: The long integer value associated with the key.
   */
  inline uint64_t GetEventUInt64(void* info, plg::string key) {
    using GetEventUInt64Fn = uint64_t (*)(void*, plg::string);
    static GetEventUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventUInt64", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the string value of a game event's key.
   * @function GetEventString
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the string value.
   * @return string: A string where the result will be stored.
   */
  inline plg::string GetEventString(void* info, plg::string key) {
    using GetEventStringFn = plg::string (*)(void*, plg::string);
    static GetEventStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventString", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the pointer value of a game event's key.
   * @function GetEventPtr
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the pointer value.
   * @return ptr64: The pointer value associated with the key.
   */
  inline void* GetEventPtr(void* info, plg::string key) {
    using GetEventPtrFn = void* (*)(void*, plg::string);
    static GetEventPtrFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventPtr", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the player controller address of a game event's key.
   * @function GetEventPlayerController
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the player controller address.
   * @return ptr64: A pointer to the player controller associated with the key.
   */
  inline void* GetEventPlayerController(void* info, plg::string key) {
    using GetEventPlayerControllerFn = void* (*)(void*, plg::string);
    static GetEventPlayerControllerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventPlayerController", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the player index of a game event's key.
   * @function GetEventPlayerIndex
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the player index.
   * @return int32: The player index associated with the key.
   */
  inline int32_t GetEventPlayerIndex(void* info, plg::string key) {
    using GetEventPlayerIndexFn = int32_t (*)(void*, plg::string);
    static GetEventPlayerIndexFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventPlayerIndex", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the player pawn address of a game event's key.
   * @function GetEventPlayerPawn
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the player pawn address.
   * @return ptr64: A pointer to the player pawn associated with the key.
   */
  inline void* GetEventPlayerPawn(void* info, plg::string key) {
    using GetEventPlayerPawnFn = void* (*)(void*, plg::string);
    static GetEventPlayerPawnFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventPlayerPawn", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the entity address of a game event's key.
   * @function GetEventEntity
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the entity address.
   * @return ptr64: A pointer to the entity associated with the key.
   */
  inline void* GetEventEntity(void* info, plg::string key) {
    using GetEventEntityFn = void* (*)(void*, plg::string);
    static GetEventEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventEntity", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the entity index of a game event's key.
   * @function GetEventEntityIndex
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the entity index.
   * @return int32: The entity index associated with the key.
   */
  inline int32_t GetEventEntityIndex(void* info, plg::string key) {
    using GetEventEntityIndexFn = int32_t (*)(void*, plg::string);
    static GetEventEntityIndexFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventEntityIndex", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the entity handle of a game event's key.
   * @function GetEventEntityHandle
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to retrieve the entity handle.
   * @return int32: The entity handle associated with the key.
   */
  inline int32_t GetEventEntityHandle(void* info, plg::string key) {
    using GetEventEntityHandleFn = int32_t (*)(void*, plg::string);
    static GetEventEntityHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventEntityHandle", reinterpret_cast<void**>(&__func));
    return __func(info, key);
  }

  /**
   * @brief Retrieves the name of a game event.
   * @function GetEventName
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @return string: A string where the result will be stored.
   */
  inline plg::string GetEventName(void* info) {
    using GetEventNameFn = plg::string (*)(void*);
    static GetEventNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEventName", reinterpret_cast<void**>(&__func));
    return __func(info);
  }

  /**
   * @brief Sets the boolean value of a game event's key.
   * @function SetEventBool
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the boolean value.
   * @param value (bool): The boolean value to set.
   */
  inline void SetEventBool(void* info, plg::string key, bool value) {
    using SetEventBoolFn = void (*)(void*, plg::string, bool);
    static SetEventBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventBool", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the floating point value of a game event's key.
   * @function SetEventFloat
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the float value.
   * @param value (float): The float value to set.
   */
  inline void SetEventFloat(void* info, plg::string key, float value) {
    using SetEventFloatFn = void (*)(void*, plg::string, float);
    static SetEventFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventFloat", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the integer value of a game event's key.
   * @function SetEventInt
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the integer value.
   * @param value (int32): The integer value to set.
   */
  inline void SetEventInt(void* info, plg::string key, int32_t value) {
    using SetEventIntFn = void (*)(void*, plg::string, int32_t);
    static SetEventIntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventInt", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the long integer value of a game event's key.
   * @function SetEventUInt64
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the long integer value.
   * @param value (uint64): The long integer value to set.
   */
  inline void SetEventUInt64(void* info, plg::string key, uint64_t value) {
    using SetEventUInt64Fn = void (*)(void*, plg::string, uint64_t);
    static SetEventUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventUInt64", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the string value of a game event's key.
   * @function SetEventString
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the string value.
   * @param value (string): The string value to set.
   */
  inline void SetEventString(void* info, plg::string key, plg::string value) {
    using SetEventStringFn = void (*)(void*, plg::string, plg::string);
    static SetEventStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventString", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the pointer value of a game event's key.
   * @function SetEventPtr
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the pointer value.
   * @param value (ptr64): The pointer value to set.
   */
  inline void SetEventPtr(void* info, plg::string key, void* value) {
    using SetEventPtrFn = void (*)(void*, plg::string, void*);
    static SetEventPtrFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventPtr", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the player controller address of a game event's key.
   * @function SetEventPlayerController
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the player controller address.
   * @param value (ptr64): A pointer to the player controller to set.
   */
  inline void SetEventPlayerController(void* info, plg::string key, void* value) {
    using SetEventPlayerControllerFn = void (*)(void*, plg::string, void*);
    static SetEventPlayerControllerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventPlayerController", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the player index value of a game event's key.
   * @function SetEventPlayerIndex
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the player index value.
   * @param value (int32): The player index value to set.
   */
  inline void SetEventPlayerIndex(void* info, plg::string key, int32_t value) {
    using SetEventPlayerIndexFn = void (*)(void*, plg::string, int32_t);
    static SetEventPlayerIndexFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventPlayerIndex", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the entity address of a game event's key.
   * @function SetEventEntity
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the entity address.
   * @param value (ptr64): A pointer to the entity to set.
   */
  inline void SetEventEntity(void* info, plg::string key, void* value) {
    using SetEventEntityFn = void (*)(void*, plg::string, void*);
    static SetEventEntityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventEntity", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the entity index of a game event's key.
   * @function SetEventEntityIndex
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the entity index.
   * @param value (int32): The entity index value to set.
   */
  inline void SetEventEntityIndex(void* info, plg::string key, int32_t value) {
    using SetEventEntityIndexFn = void (*)(void*, plg::string, int32_t);
    static SetEventEntityIndexFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventEntityIndex", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets the entity handle of a game event's key.
   * @function SetEventEntityHandle
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param key (string): The key for which to set the entity handle.
   * @param value (int32): The entity handle value to set.
   */
  inline void SetEventEntityHandle(void* info, plg::string key, int32_t value) {
    using SetEventEntityHandleFn = void (*)(void*, plg::string, int32_t);
    static SetEventEntityHandleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventEntityHandle", reinterpret_cast<void**>(&__func));
    __func(info, key, value);
  }

  /**
   * @brief Sets whether an event's broadcasting will be disabled or not.
   * @function SetEventBroadcast
   * @param info (ptr64): A pointer to the EventInfo structure containing event data.
   * @param dontBroadcast (bool): A boolean indicating whether to disable broadcasting.
   */
  inline void SetEventBroadcast(void* info, bool dontBroadcast) {
    using SetEventBroadcastFn = void (*)(void*, bool);
    static SetEventBroadcastFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEventBroadcast", reinterpret_cast<void**>(&__func));
    __func(info, dontBroadcast);
  }

  /**
   * @brief Load game event descriptions from a file (e.g., "resource/gameevents.res").
   * @function LoadEventsFromFile
   * @param path (string): The path to the file containing event descriptions.
   * @param searchAll (bool): A boolean indicating whether to search all paths for the file.
   * @return int32: An integer indicating the result of the loading operation.
   */
  inline int32_t LoadEventsFromFile(plg::string path, bool searchAll) {
    using LoadEventsFromFileFn = int32_t (*)(plg::string, bool);
    static LoadEventsFromFileFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.LoadEventsFromFile", reinterpret_cast<void**>(&__func));
    return __func(path, searchAll);
  }

  /**
   * @brief Closes a game configuration file.
   * @function CloseGameConfigFile
   * @param id (uint32): An id to the game configuration to be closed.
   */
  inline void CloseGameConfigFile(uint32_t id) {
    using CloseGameConfigFileFn = void (*)(uint32_t);
    static CloseGameConfigFileFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CloseGameConfigFile", reinterpret_cast<void**>(&__func));
    __func(id);
  }

  /**
   * @brief Loads a game configuration file.
   * @function LoadGameConfigFile
   * @param paths (string[]): The paths to the game configuration file to be loaded.
   * @return uint32: A id to the loaded game configuration object, or -1 if loading fails.
   */
  inline uint32_t LoadGameConfigFile(plg::vector<plg::string> paths) {
    using LoadGameConfigFileFn = uint32_t (*)(plg::vector<plg::string>);
    static LoadGameConfigFileFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.LoadGameConfigFile", reinterpret_cast<void**>(&__func));
    return __func(paths);
  }

  /**
   * @brief Retrieves a patch associated with the game configuration.
   * @function GetGameConfigPatch
   * @param id (uint32): An id to the game configuration from which to retrieve the patch.
   * @param name (string): The name of the patch to be retrieved.
   * @return string: A string where the patch will be stored.
   */
  inline plg::string GetGameConfigPatch(uint32_t id, plg::string name) {
    using GetGameConfigPatchFn = plg::string (*)(uint32_t, plg::string);
    static GetGameConfigPatchFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameConfigPatch", reinterpret_cast<void**>(&__func));
    return __func(id, name);
  }

  /**
   * @brief Retrieves the offset associated with a name from the game configuration.
   * @function GetGameConfigOffset
   * @param id (uint32): An id to the game configuration from which to retrieve the offset.
   * @param name (string): The name whose offset is to be retrieved.
   * @return int32: The offset associated with the specified name.
   */
  inline int32_t GetGameConfigOffset(uint32_t id, plg::string name) {
    using GetGameConfigOffsetFn = int32_t (*)(uint32_t, plg::string);
    static GetGameConfigOffsetFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameConfigOffset", reinterpret_cast<void**>(&__func));
    return __func(id, name);
  }

  /**
   * @brief Retrieves the address associated with a name from the game configuration.
   * @function GetGameConfigAddress
   * @param id (uint32): An id to the game configuration from which to retrieve the address.
   * @param name (string): The name whose address is to be retrieved.
   * @return ptr64: A pointer to the address associated with the specified name.
   */
  inline void* GetGameConfigAddress(uint32_t id, plg::string name) {
    using GetGameConfigAddressFn = void* (*)(uint32_t, plg::string);
    static GetGameConfigAddressFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameConfigAddress", reinterpret_cast<void**>(&__func));
    return __func(id, name);
  }

  /**
   * @brief Retrieves a vtable associated with the game configuration.
   * @function GetGameConfigVTable
   * @param id (uint32): An id to the game configuration from which to retrieve the vtable.
   * @param name (string): The name of the vtable to be retrieved.
   * @return ptr64: A pointer to the vtable associated with the specified name
   */
  inline void* GetGameConfigVTable(uint32_t id, plg::string name) {
    using GetGameConfigVTableFn = void* (*)(uint32_t, plg::string);
    static GetGameConfigVTableFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameConfigVTable", reinterpret_cast<void**>(&__func));
    return __func(id, name);
  }

  /**
   * @brief Retrieves the signature associated with a name from the game configuration.
   * @function GetGameConfigSignature
   * @param id (uint32): An id to the game configuration from which to retrieve the signature.
   * @param name (string): The name whose signature is to be resolved and retrieved.
   * @return ptr64: A pointer to the signature associated with the specified name.
   */
  inline void* GetGameConfigSignature(uint32_t id, plg::string name) {
    using GetGameConfigSignatureFn = void* (*)(uint32_t, plg::string);
    static GetGameConfigSignatureFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameConfigSignature", reinterpret_cast<void**>(&__func));
    return __func(id, name);
  }

  /**
   * @brief Registers a new logging channel with specified properties.
   * @function RegisterLoggingChannel
   * @param name (string): The name of the logging channel.
   * @param iFlags (int32): Flags associated with the logging channel.
   * @param verbosity (int32): The verbosity level for the logging channel.
   * @param color (int32): The color for messages logged to this channel.
   * @return int32: The ID of the newly created logging channel.
   */
  inline int32_t RegisterLoggingChannel(plg::string name, int32_t iFlags, int32_t verbosity, int32_t color) {
    using RegisterLoggingChannelFn = int32_t (*)(plg::string, int32_t, int32_t, int32_t);
    static RegisterLoggingChannelFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RegisterLoggingChannel", reinterpret_cast<void**>(&__func));
    return __func(name, iFlags, verbosity, color);
  }

  /**
   * @brief Adds a tag to a specified logging channel.
   * @function AddLoggerTagToChannel
   * @param channelID (int32): The ID of the logging channel to which the tag will be added.
   * @param tagName (string): The name of the tag to add to the channel.
   */
  inline void AddLoggerTagToChannel(int32_t channelID, plg::string tagName) {
    using AddLoggerTagToChannelFn = void (*)(int32_t, plg::string);
    static AddLoggerTagToChannelFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.AddLoggerTagToChannel", reinterpret_cast<void**>(&__func));
    __func(channelID, tagName);
  }

  /**
   * @brief Checks if a specified tag exists in a logging channel.
   * @function HasLoggerTag
   * @param channelID (int32): The ID of the logging channel.
   * @param tag (string): The name of the tag to check for.
   * @return bool: True if the tag exists in the channel, otherwise false.
   */
  inline bool HasLoggerTag(int32_t channelID, plg::string tag) {
    using HasLoggerTagFn = bool (*)(int32_t, plg::string);
    static HasLoggerTagFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.HasLoggerTag", reinterpret_cast<void**>(&__func));
    return __func(channelID, tag);
  }

  /**
   * @brief Checks if a logging channel is enabled based on severity.
   * @function IsLoggerChannelEnabledBySeverity
   * @param channelID (int32): The ID of the logging channel.
   * @param severity (int32): The severity of a logging operation.
   * @return bool: True if the channel is enabled for the specified severity, otherwise false.
   */
  inline bool IsLoggerChannelEnabledBySeverity(int32_t channelID, int32_t severity) {
    using IsLoggerChannelEnabledBySeverityFn = bool (*)(int32_t, int32_t);
    static IsLoggerChannelEnabledBySeverityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsLoggerChannelEnabledBySeverity", reinterpret_cast<void**>(&__func));
    return __func(channelID, severity);
  }

  /**
   * @brief Checks if a logging channel is enabled based on verbosity.
   * @function IsLoggerChannelEnabledByVerbosity
   * @param channelID (int32): The ID of the logging channel.
   * @param verbosity (int32): The verbosity level to check.
   * @return bool: True if the channel is enabled for the specified verbosity, otherwise false.
   */
  inline bool IsLoggerChannelEnabledByVerbosity(int32_t channelID, int32_t verbosity) {
    using IsLoggerChannelEnabledByVerbosityFn = bool (*)(int32_t, int32_t);
    static IsLoggerChannelEnabledByVerbosityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsLoggerChannelEnabledByVerbosity", reinterpret_cast<void**>(&__func));
    return __func(channelID, verbosity);
  }

  /**
   * @brief Retrieves the verbosity level of a logging channel.
   * @function GetLoggerChannelVerbosity
   * @param channelID (int32): The ID of the logging channel.
   * @return int32: The verbosity level of the specified logging channel.
   */
  inline int32_t GetLoggerChannelVerbosity(int32_t channelID) {
    using GetLoggerChannelVerbosityFn = int32_t (*)(int32_t);
    static GetLoggerChannelVerbosityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetLoggerChannelVerbosity", reinterpret_cast<void**>(&__func));
    return __func(channelID);
  }

  /**
   * @brief Sets the verbosity level of a logging channel.
   * @function SetLoggerChannelVerbosity
   * @param channelID (int32): The ID of the logging channel.
   * @param verbosity (int32): The new verbosity level to set.
   */
  inline void SetLoggerChannelVerbosity(int32_t channelID, int32_t verbosity) {
    using SetLoggerChannelVerbosityFn = void (*)(int32_t, int32_t);
    static SetLoggerChannelVerbosityFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetLoggerChannelVerbosity", reinterpret_cast<void**>(&__func));
    __func(channelID, verbosity);
  }

  /**
   * @brief Sets the verbosity level of a logging channel by name.
   * @function SetLoggerChannelVerbosityByName
   * @param channelID (int32): The ID of the logging channel.
   * @param name (string): The name of the logging channel.
   * @param verbosity (int32): The new verbosity level to set.
   */
  inline void SetLoggerChannelVerbosityByName(int32_t channelID, plg::string name, int32_t verbosity) {
    using SetLoggerChannelVerbosityByNameFn = void (*)(int32_t, plg::string, int32_t);
    static SetLoggerChannelVerbosityByNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetLoggerChannelVerbosityByName", reinterpret_cast<void**>(&__func));
    __func(channelID, name, verbosity);
  }

  /**
   * @brief Sets the verbosity level of a logging channel by tag.
   * @function SetLoggerChannelVerbosityByTag
   * @param channelID (int32): The ID of the logging channel.
   * @param tag (string): The name of the tag.
   * @param verbosity (int32): The new verbosity level to set.
   */
  inline void SetLoggerChannelVerbosityByTag(int32_t channelID, plg::string tag, int32_t verbosity) {
    using SetLoggerChannelVerbosityByTagFn = void (*)(int32_t, plg::string, int32_t);
    static SetLoggerChannelVerbosityByTagFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetLoggerChannelVerbosityByTag", reinterpret_cast<void**>(&__func));
    __func(channelID, tag, verbosity);
  }

  /**
   * @brief Retrieves the color setting of a logging channel.
   * @function GetLoggerChannelColor
   * @param channelID (int32): The ID of the logging channel.
   * @return int32: The color value of the specified logging channel.
   */
  inline int32_t GetLoggerChannelColor(int32_t channelID) {
    using GetLoggerChannelColorFn = int32_t (*)(int32_t);
    static GetLoggerChannelColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetLoggerChannelColor", reinterpret_cast<void**>(&__func));
    return __func(channelID);
  }

  /**
   * @brief Sets the color setting of a logging channel.
   * @function SetLoggerChannelColor
   * @param channelID (int32): The ID of the logging channel.
   * @param color (int32): The new color value to set for the channel.
   */
  inline void SetLoggerChannelColor(int32_t channelID, int32_t color) {
    using SetLoggerChannelColorFn = void (*)(int32_t, int32_t);
    static SetLoggerChannelColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetLoggerChannelColor", reinterpret_cast<void**>(&__func));
    __func(channelID, color);
  }

  /**
   * @brief Retrieves the flags of a logging channel.
   * @function GetLoggerChannelFlags
   * @param channelID (int32): The ID of the logging channel.
   * @return int32: The flags of the specified logging channel.
   */
  inline int32_t GetLoggerChannelFlags(int32_t channelID) {
    using GetLoggerChannelFlagsFn = int32_t (*)(int32_t);
    static GetLoggerChannelFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetLoggerChannelFlags", reinterpret_cast<void**>(&__func));
    return __func(channelID);
  }

  /**
   * @brief Sets the flags of a logging channel.
   * @function SetLoggerChannelFlags
   * @param channelID (int32): The ID of the logging channel.
   * @param eFlags (int32): The new flags to set for the channel.
   */
  inline void SetLoggerChannelFlags(int32_t channelID, int32_t eFlags) {
    using SetLoggerChannelFlagsFn = void (*)(int32_t, int32_t);
    static SetLoggerChannelFlagsFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetLoggerChannelFlags", reinterpret_cast<void**>(&__func));
    __func(channelID, eFlags);
  }

  /**
   * @brief Logs a message to a specified channel with a severity level.
   * @function Log
   * @param channelID (int32): The ID of the logging channel.
   * @param severity (int32): The severity level for the log message.
   * @param message (string): The message to log.
   * @return int32: An integer indicating the result of the logging operation.
   */
  inline int32_t Log(int32_t channelID, int32_t severity, plg::string message) {
    using LogFn = int32_t (*)(int32_t, int32_t, plg::string);
    static LogFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.Log", reinterpret_cast<void**>(&__func));
    return __func(channelID, severity, message);
  }

  /**
   * @brief Logs a colored message to a specified channel with a severity level.
   * @function LogColored
   * @param channelID (int32): The ID of the logging channel.
   * @param severity (int32): The severity level for the log message.
   * @param color (int32): The color for the log message.
   * @param message (string): The message to log.
   * @return int32: An integer indicating the result of the logging operation.
   */
  inline int32_t LogColored(int32_t channelID, int32_t severity, int32_t color, plg::string message) {
    using LogColoredFn = int32_t (*)(int32_t, int32_t, int32_t, plg::string);
    static LogColoredFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.LogColored", reinterpret_cast<void**>(&__func));
    return __func(channelID, severity, color, message);
  }

  /**
   * @brief Logs a detailed message to a specified channel, including source code info.
   * @function LogFull
   * @param channelID (int32): The ID of the logging channel.
   * @param severity (int32): The severity level for the log message.
   * @param file (string): The file name where the log call occurred.
   * @param line (int32): The line number where the log call occurred.
   * @param function (string): The name of the function where the log call occurred.
   * @param message (string): The message to log.
   * @return int32: An integer indicating the result of the logging operation.
   */
  inline int32_t LogFull(int32_t channelID, int32_t severity, plg::string file, int32_t line, plg::string function, plg::string message) {
    using LogFullFn = int32_t (*)(int32_t, int32_t, plg::string, int32_t, plg::string, plg::string);
    static LogFullFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.LogFull", reinterpret_cast<void**>(&__func));
    return __func(channelID, severity, file, line, function, message);
  }

  /**
   * @brief Logs a detailed colored message to a specified channel, including source code info.
   * @function LogFullColored
   * @param channelID (int32): The ID of the logging channel.
   * @param severity (int32): The severity level for the log message.
   * @param file (string): The file name where the log call occurred.
   * @param line (int32): The line number where the log call occurred.
   * @param function (string): The name of the function where the log call occurred.
   * @param color (int32): The color for the log message.
   * @param message (string): The message to log.
   * @return int32: An integer indicating the result of the logging operation.
   */
  inline int32_t LogFullColored(int32_t channelID, int32_t severity, plg::string file, int32_t line, plg::string function, int32_t color, plg::string message) {
    using LogFullColoredFn = int32_t (*)(int32_t, int32_t, plg::string, int32_t, plg::string, int32_t, plg::string);
    static LogFullColoredFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.LogFullColored", reinterpret_cast<void**>(&__func));
    return __func(channelID, severity, file, line, function, color, message);
  }

  /**
   * @brief Retrieves the attachment angles of an entity.
   * @function GetEntityAttachmentAngles
   * @param entityHandle (int32): The handle of the entity whose attachment angles are to be retrieved.
   * @param attachmentIndex (int32): The attachment index.
   * @return vec3: A vector representing the attachment angles (pitch, yaw, roll).
   */
  inline plg::vec3 GetEntityAttachmentAngles(int32_t entityHandle, int32_t attachmentIndex) {
    using GetEntityAttachmentAnglesFn = plg::vec3 (*)(int32_t, int32_t);
    static GetEntityAttachmentAnglesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAttachmentAngles", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, attachmentIndex);
  }

  /**
   * @brief Retrieves the forward vector of an entity's attachment.
   * @function GetEntityAttachmentForward
   * @param entityHandle (int32): The handle of the entity whose attachment forward vector is to be retrieved.
   * @param attachmentIndex (int32): The attachment index.
   * @return vec3: A vector representing the forward direction of the attachment.
   */
  inline plg::vec3 GetEntityAttachmentForward(int32_t entityHandle, int32_t attachmentIndex) {
    using GetEntityAttachmentForwardFn = plg::vec3 (*)(int32_t, int32_t);
    static GetEntityAttachmentForwardFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAttachmentForward", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, attachmentIndex);
  }

  /**
   * @brief Retrieves the origin vector of an entity's attachment.
   * @function GetEntityAttachmentOrigin
   * @param entityHandle (int32): The handle of the entity whose attachment origin is to be retrieved.
   * @param attachmentIndex (int32): The attachment index.
   * @return vec3: A vector representing the origin of the attachment.
   */
  inline plg::vec3 GetEntityAttachmentOrigin(int32_t entityHandle, int32_t attachmentIndex) {
    using GetEntityAttachmentOriginFn = plg::vec3 (*)(int32_t, int32_t);
    static GetEntityAttachmentOriginFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityAttachmentOrigin", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, attachmentIndex);
  }

  /**
   * @brief Retrieves the material group hash of an entity.
   * @function GetEntityMaterialGroupHash
   * @param entityHandle (int32): The handle of the entity.
   * @return uint32: The material group hash.
   */
  inline uint32_t GetEntityMaterialGroupHash(int32_t entityHandle) {
    using GetEntityMaterialGroupHashFn = uint32_t (*)(int32_t);
    static GetEntityMaterialGroupHashFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityMaterialGroupHash", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the material group mask of an entity.
   * @function GetEntityMaterialGroupMask
   * @param entityHandle (int32): The handle of the entity.
   * @return uint64: The mesh group mask.
   */
  inline uint64_t GetEntityMaterialGroupMask(int32_t entityHandle) {
    using GetEntityMaterialGroupMaskFn = uint64_t (*)(int32_t);
    static GetEntityMaterialGroupMaskFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityMaterialGroupMask", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the model scale of an entity.
   * @function GetEntityModelScale
   * @param entityHandle (int32): The handle of the entity.
   * @return float: The model scale factor.
   */
  inline float GetEntityModelScale(int32_t entityHandle) {
    using GetEntityModelScaleFn = float (*)(int32_t);
    static GetEntityModelScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityModelScale", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the render alpha of an entity.
   * @function GetEntityRenderAlpha
   * @param entityHandle (int32): The handle of the entity.
   * @return int32: The alpha modulation value.
   */
  inline int32_t GetEntityRenderAlpha(int32_t entityHandle) {
    using GetEntityRenderAlphaFn = int32_t (*)(int32_t);
    static GetEntityRenderAlphaFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityRenderAlpha", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the render color of an entity.
   * @function GetEntityRenderColor2
   * @param entityHandle (int32): The handle of the entity.
   * @return vec3: A vector representing the render color (R, G, B).
   */
  inline plg::vec3 GetEntityRenderColor2(int32_t entityHandle) {
    using GetEntityRenderColor2Fn = plg::vec3 (*)(int32_t);
    static GetEntityRenderColor2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntityRenderColor2", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves an attachment index by name.
   * @function ScriptLookupAttachment
   * @param entityHandle (int32): The handle of the entity.
   * @param attachmentName (string): The name of the attachment.
   * @return int32: The attachment index, or -1 if not found.
   */
  inline int32_t ScriptLookupAttachment(int32_t entityHandle, plg::string attachmentName) {
    using ScriptLookupAttachmentFn = int32_t (*)(int32_t, plg::string);
    static ScriptLookupAttachmentFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ScriptLookupAttachment", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, attachmentName);
  }

  /**
   * @brief Sets a bodygroup value by index.
   * @function SetEntityBodygroup
   * @param entityHandle (int32): The handle of the entity.
   * @param group (int32): The bodygroup index.
   * @param value (int32): The new value to set for the bodygroup.
   */
  inline void SetEntityBodygroup(int32_t entityHandle, int32_t group, int32_t value) {
    using SetEntityBodygroupFn = void (*)(int32_t, int32_t, int32_t);
    static SetEntityBodygroupFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityBodygroup", reinterpret_cast<void**>(&__func));
    __func(entityHandle, group, value);
  }

  /**
   * @brief Sets a bodygroup value by name.
   * @function SetEntityBodygroupByName
   * @param entityHandle (int32): The handle of the entity.
   * @param name (string): The bodygroup name.
   * @param value (int32): The new value to set for the bodygroup.
   */
  inline void SetEntityBodygroupByName(int32_t entityHandle, plg::string name, int32_t value) {
    using SetEntityBodygroupByNameFn = void (*)(int32_t, plg::string, int32_t);
    static SetEntityBodygroupByNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityBodygroupByName", reinterpret_cast<void**>(&__func));
    __func(entityHandle, name, value);
  }

  /**
   * @brief Sets the light group of an entity.
   * @function SetEntityLightGroup
   * @param entityHandle (int32): The handle of the entity.
   * @param lightGroup (string): The light group name.
   */
  inline void SetEntityLightGroup(int32_t entityHandle, plg::string lightGroup) {
    using SetEntityLightGroupFn = void (*)(int32_t, plg::string);
    static SetEntityLightGroupFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityLightGroup", reinterpret_cast<void**>(&__func));
    __func(entityHandle, lightGroup);
  }

  /**
   * @brief Sets the material group of an entity.
   * @function SetEntityMaterialGroup
   * @param entityHandle (int32): The handle of the entity.
   * @param materialGroup (string): The material group name.
   */
  inline void SetEntityMaterialGroup(int32_t entityHandle, plg::string materialGroup) {
    using SetEntityMaterialGroupFn = void (*)(int32_t, plg::string);
    static SetEntityMaterialGroupFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityMaterialGroup", reinterpret_cast<void**>(&__func));
    __func(entityHandle, materialGroup);
  }

  /**
   * @brief Sets the material group hash of an entity.
   * @function SetEntityMaterialGroupHash
   * @param entityHandle (int32): The handle of the entity.
   * @param hash (uint32): The new material group hash to set.
   */
  inline void SetEntityMaterialGroupHash(int32_t entityHandle, uint32_t hash) {
    using SetEntityMaterialGroupHashFn = void (*)(int32_t, uint32_t);
    static SetEntityMaterialGroupHashFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityMaterialGroupHash", reinterpret_cast<void**>(&__func));
    __func(entityHandle, hash);
  }

  /**
   * @brief Sets the material group mask of an entity.
   * @function SetEntityMaterialGroupMask
   * @param entityHandle (int32): The handle of the entity.
   * @param mask (uint64): The new mesh group mask to set.
   */
  inline void SetEntityMaterialGroupMask(int32_t entityHandle, uint64_t mask) {
    using SetEntityMaterialGroupMaskFn = void (*)(int32_t, uint64_t);
    static SetEntityMaterialGroupMaskFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityMaterialGroupMask", reinterpret_cast<void**>(&__func));
    __func(entityHandle, mask);
  }

  /**
   * @brief Sets the model scale of an entity.
   * @function SetEntityModelScale
   * @param entityHandle (int32): The handle of the entity.
   * @param scale (float): The new scale factor.
   */
  inline void SetEntityModelScale(int32_t entityHandle, float scale) {
    using SetEntityModelScaleFn = void (*)(int32_t, float);
    static SetEntityModelScaleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityModelScale", reinterpret_cast<void**>(&__func));
    __func(entityHandle, scale);
  }

  /**
   * @brief Sets the render alpha of an entity.
   * @function SetEntityRenderAlpha
   * @param entityHandle (int32): The handle of the entity.
   * @param alpha (int32): The new alpha value (0â€“255).
   */
  inline void SetEntityRenderAlpha(int32_t entityHandle, int32_t alpha) {
    using SetEntityRenderAlphaFn = void (*)(int32_t, int32_t);
    static SetEntityRenderAlphaFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityRenderAlpha", reinterpret_cast<void**>(&__func));
    __func(entityHandle, alpha);
  }

  /**
   * @brief Sets the render color of an entity.
   * @function SetEntityRenderColor2
   * @param entityHandle (int32): The handle of the entity.
   * @param r (int32): The red component (0â€“255).
   * @param g (int32): The green component (0â€“255).
   * @param b (int32): The blue component (0â€“255).
   */
  inline void SetEntityRenderColor2(int32_t entityHandle, int32_t r, int32_t g, int32_t b) {
    using SetEntityRenderColor2Fn = void (*)(int32_t, int32_t, int32_t, int32_t);
    static SetEntityRenderColor2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityRenderColor2", reinterpret_cast<void**>(&__func));
    __func(entityHandle, r, g, b);
  }

  /**
   * @brief Sets the render mode of an entity.
   * @function SetEntityRenderMode2
   * @param entityHandle (int32): The handle of the entity.
   * @param mode (int32): The new render mode value.
   */
  inline void SetEntityRenderMode2(int32_t entityHandle, int32_t mode) {
    using SetEntityRenderMode2Fn = void (*)(int32_t, int32_t);
    static SetEntityRenderMode2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntityRenderMode2", reinterpret_cast<void**>(&__func));
    __func(entityHandle, mode);
  }

  /**
   * @brief Sets a single mesh group for an entity.
   * @function SetEntitySingleMeshGroup
   * @param entityHandle (int32): The handle of the entity.
   * @param meshGroupName (string): The name of the mesh group.
   */
  inline void SetEntitySingleMeshGroup(int32_t entityHandle, plg::string meshGroupName) {
    using SetEntitySingleMeshGroupFn = void (*)(int32_t, plg::string);
    static SetEntitySingleMeshGroupFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntitySingleMeshGroup", reinterpret_cast<void**>(&__func));
    __func(entityHandle, meshGroupName);
  }

  /**
   * @brief Sets the size (bounding box) of an entity.
   * @function SetEntitySize
   * @param entityHandle (int32): The handle of the entity.
   * @param mins (vec3): The minimum bounding box vector.
   * @param maxs (vec3): The maximum bounding box vector.
   */
  inline void SetEntitySize(int32_t entityHandle, plg::vec3 mins, plg::vec3 maxs) {
    using SetEntitySizeFn = void (*)(int32_t, plg::vec3, plg::vec3);
    static SetEntitySizeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntitySize", reinterpret_cast<void**>(&__func));
    __func(entityHandle, mins, maxs);
  }

  /**
   * @brief Sets the skin of an entity.
   * @function SetEntitySkin
   * @param entityHandle (int32): The handle of the entity.
   * @param skin (int32): The new skin index.
   */
  inline void SetEntitySkin(int32_t entityHandle, int32_t skin) {
    using SetEntitySkinFn = void (*)(int32_t, int32_t);
    static SetEntitySkinFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntitySkin", reinterpret_cast<void**>(&__func));
    __func(entityHandle, skin);
  }

  /**
   * @brief Start a new Yes/No vote
   * @function PanoramaSendYesNoVote
   * @param duration (double): Maximum time to leave vote active for
   * @param caller (int32): Player slot of the vote caller. Use VOTE_CALLER_SERVER for 'Server'.
   * @param voteTitle (string): Translation string to use as the vote message. (Only '#SFUI_vote' or '#Panorama_vote' strings)
   * @param detailStr (string): Extra string used in some vote translation strings.
   * @param votePassTitle (string): Translation string to use as the vote message. (Only '#SFUI_vote' or '#Panorama_vote' strings)
   * @param detailPassStr (string): Extra string used in some vote translation strings when the vote passes.
   * @param failReason (int32): Reason for the vote to fail, used in some translation strings.
   * @param filter (uint64): Recipient filter with all the clients who are allowed to participate in the vote.
   * @param result (function): Called when a menu action is completed.
   * @param handler (function): Called when the vote has finished.
   * @return bool
   */
  inline bool PanoramaSendYesNoVote(double duration, int32_t caller, plg::string voteTitle, plg::string detailStr, plg::string votePassTitle, plg::string detailPassStr, int32_t failReason, uint64_t filter, function result, function handler) {
    using PanoramaSendYesNoVoteFn = bool (*)(double, int32_t, plg::string, plg::string, plg::string, plg::string, int32_t, uint64_t, function, function);
    static PanoramaSendYesNoVoteFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PanoramaSendYesNoVote", reinterpret_cast<void**>(&__func));
    return __func(duration, caller, voteTitle, detailStr, votePassTitle, detailPassStr, failReason, filter, result, handler);
  }

  /**
   * @brief Start a new Yes/No vote with all players included
   * @function PanoramaSendYesNoVoteToAll
   * @param duration (double): Maximum time to leave vote active for
   * @param caller (int32): Player slot of the vote caller. Use VOTE_CALLER_SERVER for 'Server'.
   * @param voteTitle (string): Translation string to use as the vote message. (Only '#SFUI_vote' or '#Panorama_vote' strings)
   * @param detailStr (string): Extra string used in some vote translation strings.
   * @param votePassTitle (string): Translation string to use as the vote message. (Only '#SFUI_vote' or '#Panorama_vote' strings)
   * @param detailPassStr (string): Extra string used in some vote translation strings when the vote passes.
   * @param failReason (int32): Reason for the vote to fail, used in some translation strings.
   * @param result (function): Called when a menu action is completed.
   * @param handler (function): Called when the vote has finished.
   * @return bool
   */
  inline bool PanoramaSendYesNoVoteToAll(double duration, int32_t caller, plg::string voteTitle, plg::string detailStr, plg::string votePassTitle, plg::string detailPassStr, int32_t failReason, function result, function handler) {
    using PanoramaSendYesNoVoteToAllFn = bool (*)(double, int32_t, plg::string, plg::string, plg::string, plg::string, int32_t, function, function);
    static PanoramaSendYesNoVoteToAllFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PanoramaSendYesNoVoteToAll", reinterpret_cast<void**>(&__func));
    return __func(duration, caller, voteTitle, detailStr, votePassTitle, detailPassStr, failReason, result, handler);
  }

  /**
   * @brief Removes a player from the current vote.
   * @function PanoramaRemovePlayerFromVote
   * @param playerSlot (int32): The slot/index of the player to remove from the vote.
   */
  inline void PanoramaRemovePlayerFromVote(int32_t playerSlot) {
    using PanoramaRemovePlayerFromVoteFn = void (*)(int32_t);
    static PanoramaRemovePlayerFromVoteFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PanoramaRemovePlayerFromVote", reinterpret_cast<void**>(&__func));
    __func(playerSlot);
  }

  /**
   * @brief Checks if a player is in the vote pool.
   * @function PanoramaIsPlayerInVotePool
   * @param playerSlot (int32): The slot/index of the player to check.
   * @return bool: true if the player is in the vote pool, false otherwise.
   */
  inline bool PanoramaIsPlayerInVotePool(int32_t playerSlot) {
    using PanoramaIsPlayerInVotePoolFn = bool (*)(int32_t);
    static PanoramaIsPlayerInVotePoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PanoramaIsPlayerInVotePool", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Redraws the vote UI to a specific player client.
   * @function PanoramaRedrawVoteToClient
   * @param playerSlot (int32): The slot/index of the player to update.
   * @return bool: true if the vote UI was successfully redrawn, false otherwise.
   */
  inline bool PanoramaRedrawVoteToClient(int32_t playerSlot) {
    using PanoramaRedrawVoteToClientFn = bool (*)(int32_t);
    static PanoramaRedrawVoteToClientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PanoramaRedrawVoteToClient", reinterpret_cast<void**>(&__func));
    return __func(playerSlot);
  }

  /**
   * @brief Checks if a vote is currently in progress.
   * @function PanoramaIsVoteInProgress
   * @return bool: true if a vote is active, false otherwise.
   */
  inline bool PanoramaIsVoteInProgress() {
    using PanoramaIsVoteInProgressFn = bool (*)();
    static PanoramaIsVoteInProgressFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PanoramaIsVoteInProgress", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Ends the current vote with a specified reason.
   * @function PanoramaEndVote
   * @param reason (int32): The reason for ending the vote.
   */
  inline void PanoramaEndVote(int32_t reason) {
    using PanoramaEndVoteFn = void (*)(int32_t);
    static PanoramaEndVoteFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PanoramaEndVote", reinterpret_cast<void**>(&__func));
    __func(reason);
  }

  /**
   * @brief Get the offset of a member in a given schema class.
   * @function GetSchemaOffset
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the member whose offset is to be retrieved.
   * @return int32: The offset of the member in the class.
   */
  inline int32_t GetSchemaOffset(plg::string className, plg::string memberName) {
    using GetSchemaOffsetFn = int32_t (*)(plg::string, plg::string);
    static GetSchemaOffsetFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetSchemaOffset", reinterpret_cast<void**>(&__func));
    return __func(className, memberName);
  }

  /**
   * @brief Get the offset of a chain in a given schema class.
   * @function GetSchemaChainOffset
   * @param className (string): The name of the class.
   * @return int32: The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline int32_t GetSchemaChainOffset(plg::string className) {
    using GetSchemaChainOffsetFn = int32_t (*)(plg::string);
    static GetSchemaChainOffsetFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetSchemaChainOffset", reinterpret_cast<void**>(&__func));
    return __func(className);
  }

  /**
   * @brief Check if a schema field is networked.
   * @function IsSchemaFieldNetworked
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the member to check.
   * @return bool: True if the member is networked, false otherwise.
   */
  inline bool IsSchemaFieldNetworked(plg::string className, plg::string memberName) {
    using IsSchemaFieldNetworkedFn = bool (*)(plg::string, plg::string);
    static IsSchemaFieldNetworkedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.IsSchemaFieldNetworked", reinterpret_cast<void**>(&__func));
    return __func(className, memberName);
  }

  /**
   * @brief Get the size of a schema class.
   * @function GetSchemaClassSize
   * @param className (string): The name of the class.
   * @return int32: The size of the class in bytes, or -1 if the class is not found.
   */
  inline int32_t GetSchemaClassSize(plg::string className) {
    using GetSchemaClassSizeFn = int32_t (*)(plg::string);
    static GetSchemaClassSizeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetSchemaClassSize", reinterpret_cast<void**>(&__func));
    return __func(className);
  }

  /**
   * @brief Peeks into an entity's object schema and retrieves the integer value at the given offset.
   * @function GetEntData2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @param size (int32): Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @return int64: The integer value at the given memory location.
   */
  inline int64_t GetEntData2(void* entity, int32_t offset, int32_t size) {
    using GetEntData2Fn = int64_t (*)(void*, int32_t, int32_t);
    static GetEntData2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntData2", reinterpret_cast<void**>(&__func));
    return __func(entity, offset, size);
  }

  /**
   * @brief Peeks into an entity's object data and sets the integer value at the given offset.
   * @function SetEntData2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @param value (int64): The integer value to set.
   * @param size (int32): Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntData2(void* entity, int32_t offset, int64_t value, int32_t size, bool changeState, int32_t chainOffset) {
    using SetEntData2Fn = void (*)(void*, int32_t, int64_t, int32_t, bool, int32_t);
    static SetEntData2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntData2", reinterpret_cast<void**>(&__func));
    __func(entity, offset, value, size, changeState, chainOffset);
  }

  /**
   * @brief Peeks into an entity's object schema and retrieves the float value at the given offset.
   * @function GetEntDataFloat2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @param size (int32): Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @return double: The float value at the given memory location.
   */
  inline double GetEntDataFloat2(void* entity, int32_t offset, int32_t size) {
    using GetEntDataFloat2Fn = double (*)(void*, int32_t, int32_t);
    static GetEntDataFloat2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntDataFloat2", reinterpret_cast<void**>(&__func));
    return __func(entity, offset, size);
  }

  /**
   * @brief Peeks into an entity's object data and sets the float value at the given offset.
   * @function SetEntDataFloat2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @param value (double): The float value to set.
   * @param size (int32): Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntDataFloat2(void* entity, int32_t offset, double value, int32_t size, bool changeState, int32_t chainOffset) {
    using SetEntDataFloat2Fn = void (*)(void*, int32_t, double, int32_t, bool, int32_t);
    static SetEntDataFloat2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntDataFloat2", reinterpret_cast<void**>(&__func));
    __func(entity, offset, value, size, changeState, chainOffset);
  }

  /**
   * @brief Peeks into an entity's object schema and retrieves the string value at the given offset.
   * @function GetEntDataString2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @return string: The string value at the given memory location.
   */
  inline plg::string GetEntDataString2(void* entity, int32_t offset) {
    using GetEntDataString2Fn = plg::string (*)(void*, int32_t);
    static GetEntDataString2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntDataString2", reinterpret_cast<void**>(&__func));
    return __func(entity, offset);
  }

  /**
   * @brief Peeks into an entity's object data and sets the string at the given offset.
   * @function SetEntDataString2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @param value (string): The string value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntDataString2(void* entity, int32_t offset, plg::string value, bool changeState, int32_t chainOffset) {
    using SetEntDataString2Fn = void (*)(void*, int32_t, plg::string, bool, int32_t);
    static SetEntDataString2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntDataString2", reinterpret_cast<void**>(&__func));
    __func(entity, offset, value, changeState, chainOffset);
  }

  /**
   * @brief Peeks into an entity's object schema and retrieves the vector value at the given offset.
   * @function GetEntDataVector2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @return vec3: The vector value at the given memory location.
   */
  inline plg::vec3 GetEntDataVector2(void* entity, int32_t offset) {
    using GetEntDataVector2Fn = plg::vec3 (*)(void*, int32_t);
    static GetEntDataVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntDataVector2", reinterpret_cast<void**>(&__func));
    return __func(entity, offset);
  }

  /**
   * @brief Peeks into an entity's object data and sets the vector at the given offset.
   * @function SetEntDataVector2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @param value (vec3): The vector value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntDataVector2(void* entity, int32_t offset, plg::vec3 value, bool changeState, int32_t chainOffset) {
    using SetEntDataVector2Fn = void (*)(void*, int32_t, plg::vec3, bool, int32_t);
    static SetEntDataVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntDataVector2", reinterpret_cast<void**>(&__func));
    __func(entity, offset, value, changeState, chainOffset);
  }

  /**
   * @brief Peeks into an entity's object data and retrieves the entity handle at the given offset.
   * @function GetEntDataEnt2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @return int32: The entity handle at the given memory location.
   */
  inline int32_t GetEntDataEnt2(void* entity, int32_t offset) {
    using GetEntDataEnt2Fn = int32_t (*)(void*, int32_t);
    static GetEntDataEnt2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntDataEnt2", reinterpret_cast<void**>(&__func));
    return __func(entity, offset);
  }

  /**
   * @brief Peeks into an entity's object data and sets the entity handle at the given offset.
   * @function SetEntDataEnt2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @param value (int32): The entity handle to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntDataEnt2(void* entity, int32_t offset, int32_t value, bool changeState, int32_t chainOffset) {
    using SetEntDataEnt2Fn = void (*)(void*, int32_t, int32_t, bool, int32_t);
    static SetEntDataEnt2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntDataEnt2", reinterpret_cast<void**>(&__func));
    __func(entity, offset, value, changeState, chainOffset);
  }

  /**
   * @brief Updates the networked state of a schema field for a given entity pointer.
   * @function ChangeEntityState2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param offset (int32): The offset of the schema to use.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void ChangeEntityState2(void* entity, int32_t offset, int32_t chainOffset) {
    using ChangeEntityState2Fn = void (*)(void*, int32_t, int32_t);
    static ChangeEntityState2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ChangeEntityState2", reinterpret_cast<void**>(&__func));
    __func(entity, offset, chainOffset);
  }

  /**
   * @brief Peeks into an entity's object schema and retrieves the integer value at the given offset.
   * @function GetEntData
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @param size (int32): Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @return int64: The integer value at the given memory location.
   */
  inline int64_t GetEntData(int32_t entityHandle, int32_t offset, int32_t size) {
    using GetEntDataFn = int64_t (*)(int32_t, int32_t, int32_t);
    static GetEntDataFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntData", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, offset, size);
  }

  /**
   * @brief Peeks into an entity's object data and sets the integer value at the given offset.
   * @function SetEntData
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @param value (int64): The integer value to set.
   * @param size (int32): Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntData(int32_t entityHandle, int32_t offset, int64_t value, int32_t size, bool changeState, int32_t chainOffset) {
    using SetEntDataFn = void (*)(int32_t, int32_t, int64_t, int32_t, bool, int32_t);
    static SetEntDataFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntData", reinterpret_cast<void**>(&__func));
    __func(entityHandle, offset, value, size, changeState, chainOffset);
  }

  /**
   * @brief Peeks into an entity's object schema and retrieves the float value at the given offset.
   * @function GetEntDataFloat
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @param size (int32): Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @return double: The float value at the given memory location.
   */
  inline double GetEntDataFloat(int32_t entityHandle, int32_t offset, int32_t size) {
    using GetEntDataFloatFn = double (*)(int32_t, int32_t, int32_t);
    static GetEntDataFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntDataFloat", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, offset, size);
  }

  /**
   * @brief Peeks into an entity's object data and sets the float value at the given offset.
   * @function SetEntDataFloat
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @param value (double): The float value to set.
   * @param size (int32): Number of bytes to write (valid values are 1, 2, 4 or 8).
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntDataFloat(int32_t entityHandle, int32_t offset, double value, int32_t size, bool changeState, int32_t chainOffset) {
    using SetEntDataFloatFn = void (*)(int32_t, int32_t, double, int32_t, bool, int32_t);
    static SetEntDataFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntDataFloat", reinterpret_cast<void**>(&__func));
    __func(entityHandle, offset, value, size, changeState, chainOffset);
  }

  /**
   * @brief Peeks into an entity's object schema and retrieves the string value at the given offset.
   * @function GetEntDataString
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @return string: The string value at the given memory location.
   */
  inline plg::string GetEntDataString(int32_t entityHandle, int32_t offset) {
    using GetEntDataStringFn = plg::string (*)(int32_t, int32_t);
    static GetEntDataStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntDataString", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, offset);
  }

  /**
   * @brief Peeks into an entity's object data and sets the string at the given offset.
   * @function SetEntDataString
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @param value (string): The string value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntDataString(int32_t entityHandle, int32_t offset, plg::string value, bool changeState, int32_t chainOffset) {
    using SetEntDataStringFn = void (*)(int32_t, int32_t, plg::string, bool, int32_t);
    static SetEntDataStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntDataString", reinterpret_cast<void**>(&__func));
    __func(entityHandle, offset, value, changeState, chainOffset);
  }

  /**
   * @brief Peeks into an entity's object schema and retrieves the vector value at the given offset.
   * @function GetEntDataVector
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @return vec3: The vector value at the given memory location.
   */
  inline plg::vec3 GetEntDataVector(int32_t entityHandle, int32_t offset) {
    using GetEntDataVectorFn = plg::vec3 (*)(int32_t, int32_t);
    static GetEntDataVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntDataVector", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, offset);
  }

  /**
   * @brief Peeks into an entity's object data and sets the vector at the given offset.
   * @function SetEntDataVector
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @param value (vec3): The vector value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntDataVector(int32_t entityHandle, int32_t offset, plg::vec3 value, bool changeState, int32_t chainOffset) {
    using SetEntDataVectorFn = void (*)(int32_t, int32_t, plg::vec3, bool, int32_t);
    static SetEntDataVectorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntDataVector", reinterpret_cast<void**>(&__func));
    __func(entityHandle, offset, value, changeState, chainOffset);
  }

  /**
   * @brief Peeks into an entity's object data and retrieves the entity handle at the given offset.
   * @function GetEntDataEnt
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @return int32: The entity handle at the given memory location.
   */
  inline int32_t GetEntDataEnt(int32_t entityHandle, int32_t offset) {
    using GetEntDataEntFn = int32_t (*)(int32_t, int32_t);
    static GetEntDataEntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntDataEnt", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, offset);
  }

  /**
   * @brief Peeks into an entity's object data and sets the entity handle at the given offset.
   * @function SetEntDataEnt
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @param value (int32): The entity handle to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void SetEntDataEnt(int32_t entityHandle, int32_t offset, int32_t value, bool changeState, int32_t chainOffset) {
    using SetEntDataEntFn = void (*)(int32_t, int32_t, int32_t, bool, int32_t);
    static SetEntDataEntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntDataEnt", reinterpret_cast<void**>(&__func));
    __func(entityHandle, offset, value, changeState, chainOffset);
  }

  /**
   * @brief Updates the networked state of a schema field for a given entity handle.
   * @function ChangeEntityState
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param offset (int32): The offset of the schema to use.
   * @param chainOffset (int32): The offset of the chain entity in the class (-1 for non-entity classes).
   */
  inline void ChangeEntityState(int32_t entityHandle, int32_t offset, int32_t chainOffset) {
    using ChangeEntityStateFn = void (*)(int32_t, int32_t, int32_t);
    static ChangeEntityStateFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.ChangeEntityState", reinterpret_cast<void**>(&__func));
    __func(entityHandle, offset, chainOffset);
  }

  /**
   * @brief Retrieves the count of values that an entity schema's array can store.
   * @function GetEntSchemaArraySize2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @return int32: Size of array (in elements) or 0 if schema is not an array.
   */
  inline int32_t GetEntSchemaArraySize2(void* entity, plg::string className, plg::string memberName) {
    using GetEntSchemaArraySize2Fn = int32_t (*)(void*, plg::string, plg::string);
    static GetEntSchemaArraySize2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaArraySize2", reinterpret_cast<void**>(&__func));
    return __func(entity, className, memberName);
  }

  /**
   * @brief Retrieves an integer value from an entity's schema.
   * @function GetEntSchema2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return int64: An integer value at the given schema offset.
   */
  inline int64_t GetEntSchema2(void* entity, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchema2Fn = int64_t (*)(void*, plg::string, plg::string, int32_t);
    static GetEntSchema2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchema2", reinterpret_cast<void**>(&__func));
    return __func(entity, className, memberName, element);
  }

  /**
   * @brief Sets an integer value in an entity's schema.
   * @function SetEntSchema2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (int64): The integer value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchema2(void* entity, plg::string className, plg::string memberName, int64_t value, bool changeState, int32_t element) {
    using SetEntSchema2Fn = void (*)(void*, plg::string, plg::string, int64_t, bool, int32_t);
    static SetEntSchema2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchema2", reinterpret_cast<void**>(&__func));
    __func(entity, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a float value from an entity's schema.
   * @function GetEntSchemaFloat2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return double: A float value at the given schema offset.
   */
  inline double GetEntSchemaFloat2(void* entity, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaFloat2Fn = double (*)(void*, plg::string, plg::string, int32_t);
    static GetEntSchemaFloat2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaFloat2", reinterpret_cast<void**>(&__func));
    return __func(entity, className, memberName, element);
  }

  /**
   * @brief Sets a float value in an entity's schema.
   * @function SetEntSchemaFloat2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (double): The float value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaFloat2(void* entity, plg::string className, plg::string memberName, double value, bool changeState, int32_t element) {
    using SetEntSchemaFloat2Fn = void (*)(void*, plg::string, plg::string, double, bool, int32_t);
    static SetEntSchemaFloat2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaFloat2", reinterpret_cast<void**>(&__func));
    __func(entity, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a string value from an entity's schema.
   * @function GetEntSchemaString2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return string: A string value at the given schema offset.
   */
  inline plg::string GetEntSchemaString2(void* entity, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaString2Fn = plg::string (*)(void*, plg::string, plg::string, int32_t);
    static GetEntSchemaString2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaString2", reinterpret_cast<void**>(&__func));
    return __func(entity, className, memberName, element);
  }

  /**
   * @brief Sets a string value in an entity's schema.
   * @function SetEntSchemaString2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (string): The string value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaString2(void* entity, plg::string className, plg::string memberName, plg::string value, bool changeState, int32_t element) {
    using SetEntSchemaString2Fn = void (*)(void*, plg::string, plg::string, plg::string, bool, int32_t);
    static SetEntSchemaString2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaString2", reinterpret_cast<void**>(&__func));
    __func(entity, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a vector value from an entity's schema.
   * @function GetEntSchemaVector3D2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return vec3: A vector value at the given schema offset.
   */
  inline plg::vec3 GetEntSchemaVector3D2(void* entity, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaVector3D2Fn = plg::vec3 (*)(void*, plg::string, plg::string, int32_t);
    static GetEntSchemaVector3D2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaVector3D2", reinterpret_cast<void**>(&__func));
    return __func(entity, className, memberName, element);
  }

  /**
   * @brief Sets a vector value in an entity's schema.
   * @function SetEntSchemaVector3D2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (vec3): The vector value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaVector3D2(void* entity, plg::string className, plg::string memberName, plg::vec3 value, bool changeState, int32_t element) {
    using SetEntSchemaVector3D2Fn = void (*)(void*, plg::string, plg::string, plg::vec3, bool, int32_t);
    static SetEntSchemaVector3D2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaVector3D2", reinterpret_cast<void**>(&__func));
    __func(entity, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a vector value from an entity's schema.
   * @function GetEntSchemaVector2D2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return vec2: A vector value at the given schema offset.
   */
  inline plg::vec2 GetEntSchemaVector2D2(void* entity, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaVector2D2Fn = plg::vec2 (*)(void*, plg::string, plg::string, int32_t);
    static GetEntSchemaVector2D2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaVector2D2", reinterpret_cast<void**>(&__func));
    return __func(entity, className, memberName, element);
  }

  /**
   * @brief Sets a vector value in an entity's schema.
   * @function SetEntSchemaVector2D2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (vec2): The vector value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaVector2D2(void* entity, plg::string className, plg::string memberName, plg::vec2 value, bool changeState, int32_t element) {
    using SetEntSchemaVector2D2Fn = void (*)(void*, plg::string, plg::string, plg::vec2, bool, int32_t);
    static SetEntSchemaVector2D2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaVector2D2", reinterpret_cast<void**>(&__func));
    __func(entity, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a vector value from an entity's schema.
   * @function GetEntSchemaVector4D2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return vec4: A vector value at the given schema offset.
   */
  inline plg::vec4 GetEntSchemaVector4D2(void* entity, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaVector4D2Fn = plg::vec4 (*)(void*, plg::string, plg::string, int32_t);
    static GetEntSchemaVector4D2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaVector4D2", reinterpret_cast<void**>(&__func));
    return __func(entity, className, memberName, element);
  }

  /**
   * @brief Sets a vector value in an entity's schema.
   * @function SetEntSchemaVector4D2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (vec4): The vector value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaVector4D2(void* entity, plg::string className, plg::string memberName, plg::vec4 value, bool changeState, int32_t element) {
    using SetEntSchemaVector4D2Fn = void (*)(void*, plg::string, plg::string, plg::vec4, bool, int32_t);
    static SetEntSchemaVector4D2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaVector4D2", reinterpret_cast<void**>(&__func));
    __func(entity, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves an entity handle from an entity's schema.
   * @function GetEntSchemaEnt2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return int32: A string value at the given schema offset.
   */
  inline int32_t GetEntSchemaEnt2(void* entity, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaEnt2Fn = int32_t (*)(void*, plg::string, plg::string, int32_t);
    static GetEntSchemaEnt2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaEnt2", reinterpret_cast<void**>(&__func));
    return __func(entity, className, memberName, element);
  }

  /**
   * @brief Sets an entity handle in an entity's schema.
   * @function SetEntSchemaEnt2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (int32): The entity handle to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaEnt2(void* entity, plg::string className, plg::string memberName, int32_t value, bool changeState, int32_t element) {
    using SetEntSchemaEnt2Fn = void (*)(void*, plg::string, plg::string, int32_t, bool, int32_t);
    static SetEntSchemaEnt2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaEnt2", reinterpret_cast<void**>(&__func));
    __func(entity, className, memberName, value, changeState, element);
  }

  /**
   * @brief Updates the networked state of a schema field for a given entity pointer.
   * @function NetworkStateChanged2
   * @param entity (ptr64): Pointer to the instance of the class where the value is to be set.
   * @param className (string): The name of the class that contains the member.
   * @param memberName (string): The name of the member to be set.
   */
  inline void NetworkStateChanged2(void* entity, plg::string className, plg::string memberName) {
    using NetworkStateChanged2Fn = void (*)(void*, plg::string, plg::string);
    static NetworkStateChanged2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.NetworkStateChanged2", reinterpret_cast<void**>(&__func));
    __func(entity, className, memberName);
  }

  /**
   * @brief Retrieves the count of values that an entity schema's array can store.
   * @function GetEntSchemaArraySize
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @return int32: Size of array (in elements) or 0 if schema is not an array.
   */
  inline int32_t GetEntSchemaArraySize(int32_t entityHandle, plg::string className, plg::string memberName) {
    using GetEntSchemaArraySizeFn = int32_t (*)(int32_t, plg::string, plg::string);
    static GetEntSchemaArraySizeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaArraySize", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, className, memberName);
  }

  /**
   * @brief Retrieves an integer value from an entity's schema.
   * @function GetEntSchema
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return int64: An integer value at the given schema offset.
   */
  inline int64_t GetEntSchema(int32_t entityHandle, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaFn = int64_t (*)(int32_t, plg::string, plg::string, int32_t);
    static GetEntSchemaFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchema", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, className, memberName, element);
  }

  /**
   * @brief Sets an integer value in an entity's schema.
   * @function SetEntSchema
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (int64): The integer value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchema(int32_t entityHandle, plg::string className, plg::string memberName, int64_t value, bool changeState, int32_t element) {
    using SetEntSchemaFn = void (*)(int32_t, plg::string, plg::string, int64_t, bool, int32_t);
    static SetEntSchemaFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchema", reinterpret_cast<void**>(&__func));
    __func(entityHandle, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a float value from an entity's schema.
   * @function GetEntSchemaFloat
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return double: A float value at the given schema offset.
   */
  inline double GetEntSchemaFloat(int32_t entityHandle, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaFloatFn = double (*)(int32_t, plg::string, plg::string, int32_t);
    static GetEntSchemaFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaFloat", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, className, memberName, element);
  }

  /**
   * @brief Sets a float value in an entity's schema.
   * @function SetEntSchemaFloat
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (double): The float value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaFloat(int32_t entityHandle, plg::string className, plg::string memberName, double value, bool changeState, int32_t element) {
    using SetEntSchemaFloatFn = void (*)(int32_t, plg::string, plg::string, double, bool, int32_t);
    static SetEntSchemaFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaFloat", reinterpret_cast<void**>(&__func));
    __func(entityHandle, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a string value from an entity's schema.
   * @function GetEntSchemaString
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return string: A string value at the given schema offset.
   */
  inline plg::string GetEntSchemaString(int32_t entityHandle, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaStringFn = plg::string (*)(int32_t, plg::string, plg::string, int32_t);
    static GetEntSchemaStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaString", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, className, memberName, element);
  }

  /**
   * @brief Sets a string value in an entity's schema.
   * @function SetEntSchemaString
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (string): The string value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaString(int32_t entityHandle, plg::string className, plg::string memberName, plg::string value, bool changeState, int32_t element) {
    using SetEntSchemaStringFn = void (*)(int32_t, plg::string, plg::string, plg::string, bool, int32_t);
    static SetEntSchemaStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaString", reinterpret_cast<void**>(&__func));
    __func(entityHandle, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a vector value from an entity's schema.
   * @function GetEntSchemaVector3D
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return vec3: A string value at the given schema offset.
   */
  inline plg::vec3 GetEntSchemaVector3D(int32_t entityHandle, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaVector3DFn = plg::vec3 (*)(int32_t, plg::string, plg::string, int32_t);
    static GetEntSchemaVector3DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaVector3D", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, className, memberName, element);
  }

  /**
   * @brief Sets a vector value in an entity's schema.
   * @function SetEntSchemaVector3D
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (vec3): The vector value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaVector3D(int32_t entityHandle, plg::string className, plg::string memberName, plg::vec3 value, bool changeState, int32_t element) {
    using SetEntSchemaVector3DFn = void (*)(int32_t, plg::string, plg::string, plg::vec3, bool, int32_t);
    static SetEntSchemaVector3DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaVector3D", reinterpret_cast<void**>(&__func));
    __func(entityHandle, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a vector value from an entity's schema.
   * @function GetEntSchemaVector2D
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return vec2: A string value at the given schema offset.
   */
  inline plg::vec2 GetEntSchemaVector2D(int32_t entityHandle, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaVector2DFn = plg::vec2 (*)(int32_t, plg::string, plg::string, int32_t);
    static GetEntSchemaVector2DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaVector2D", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, className, memberName, element);
  }

  /**
   * @brief Sets a vector value in an entity's schema.
   * @function SetEntSchemaVector2D
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (vec2): The vector value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaVector2D(int32_t entityHandle, plg::string className, plg::string memberName, plg::vec2 value, bool changeState, int32_t element) {
    using SetEntSchemaVector2DFn = void (*)(int32_t, plg::string, plg::string, plg::vec2, bool, int32_t);
    static SetEntSchemaVector2DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaVector2D", reinterpret_cast<void**>(&__func));
    __func(entityHandle, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves a vector value from an entity's schema.
   * @function GetEntSchemaVector4D
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return vec4: A string value at the given schema offset.
   */
  inline plg::vec4 GetEntSchemaVector4D(int32_t entityHandle, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaVector4DFn = plg::vec4 (*)(int32_t, plg::string, plg::string, int32_t);
    static GetEntSchemaVector4DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaVector4D", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, className, memberName, element);
  }

  /**
   * @brief Sets a vector value in an entity's schema.
   * @function SetEntSchemaVector4D
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (vec4): The vector value to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaVector4D(int32_t entityHandle, plg::string className, plg::string memberName, plg::vec4 value, bool changeState, int32_t element) {
    using SetEntSchemaVector4DFn = void (*)(int32_t, plg::string, plg::string, plg::vec4, bool, int32_t);
    static SetEntSchemaVector4DFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaVector4D", reinterpret_cast<void**>(&__func));
    __func(entityHandle, className, memberName, value, changeState, element);
  }

  /**
   * @brief Retrieves an entity handle from an entity's schema.
   * @function GetEntSchemaEnt
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   * @return int32: A string value at the given schema offset.
   */
  inline int32_t GetEntSchemaEnt(int32_t entityHandle, plg::string className, plg::string memberName, int32_t element) {
    using GetEntSchemaEntFn = int32_t (*)(int32_t, plg::string, plg::string, int32_t);
    static GetEntSchemaEntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetEntSchemaEnt", reinterpret_cast<void**>(&__func));
    return __func(entityHandle, className, memberName, element);
  }

  /**
   * @brief Sets an entity handle in an entity's schema.
   * @function SetEntSchemaEnt
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class.
   * @param memberName (string): The name of the schema member.
   * @param value (int32): The entity handle to set.
   * @param changeState (bool): If true, change will be sent over the network.
   * @param element (int32): Element # (starting from 0) if schema is an array.
   */
  inline void SetEntSchemaEnt(int32_t entityHandle, plg::string className, plg::string memberName, int32_t value, bool changeState, int32_t element) {
    using SetEntSchemaEntFn = void (*)(int32_t, plg::string, plg::string, int32_t, bool, int32_t);
    static SetEntSchemaEntFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.SetEntSchemaEnt", reinterpret_cast<void**>(&__func));
    __func(entityHandle, className, memberName, value, changeState, element);
  }

  /**
   * @brief Updates the networked state of a schema field for a given entity handle.
   * @function NetworkStateChanged
   * @param entityHandle (int32): The handle of the entity from which the value is to be retrieved.
   * @param className (string): The name of the class that contains the member.
   * @param memberName (string): The name of the member to be set.
   */
  inline void NetworkStateChanged(int32_t entityHandle, plg::string className, plg::string memberName) {
    using NetworkStateChangedFn = void (*)(int32_t, plg::string, plg::string);
    static NetworkStateChangedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.NetworkStateChanged", reinterpret_cast<void**>(&__func));
    __func(entityHandle, className, memberName);
  }

  /**
   * @brief Creates a new timer that executes a callback function at specified delays.
   * @function CreateTimer
   * @param delay (double): The time delay in seconds between each callback execution.
   * @param callback (function): The function to be called when the timer expires.
   * @param flags (int32): Flags that modify the behavior of the timer (e.g., no-map change, repeating).
   * @param userData (any[]): An array intended to hold user-related data, allowing for elements of any type.
   * @return uint32: A id to the newly created Timer object, or -1 if the timer could not be created.
   */
  inline uint32_t CreateTimer(double delay, function callback, int32_t flags, plg::vector<plg::any> userData) {
    using CreateTimerFn = uint32_t (*)(double, function, int32_t, plg::vector<plg::any>);
    static CreateTimerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.CreateTimer", reinterpret_cast<void**>(&__func));
    return __func(delay, callback, flags, userData);
  }

  /**
   * @brief Stops and removes an existing timer.
   * @function KillsTimer
   * @param timer (uint32): A id of the Timer object to be stopped and removed.
   */
  inline void KillsTimer(uint32_t timer) {
    using KillsTimerFn = void (*)(uint32_t);
    static KillsTimerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.KillsTimer", reinterpret_cast<void**>(&__func));
    __func(timer);
  }

  /**
   * @brief Reschedules an existing timer with a new delay.
   * @function RescheduleTimer
   * @param timer (uint32): A id of the Timer object to be stopped and removed.
   * @param newDaly (double): The new delay in seconds between each callback execution.
   */
  inline void RescheduleTimer(uint32_t timer, double newDaly) {
    using RescheduleTimerFn = void (*)(uint32_t, double);
    static RescheduleTimerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.RescheduleTimer", reinterpret_cast<void**>(&__func));
    __func(timer, newDaly);
  }

  /**
   * @brief Returns the number of seconds in between game server ticks.
   * @function GetTickInterval
   * @return double: The tick interval value.
   */
  inline double GetTickInterval() {
    using GetTickIntervalFn = double (*)();
    static GetTickIntervalFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetTickInterval", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Returns the simulated game time.
   * @function GetTickedTime
   * @return double: The ticked time value.
   */
  inline double GetTickedTime() {
    using GetTickedTimeFn = double (*)();
    static GetTickedTimeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetTickedTime", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Register callback to event.
   * @function OnClientConnect_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientConnect_Register(function callback) {
    using OnClientConnect_RegisterFn = void (*)(function);
    static OnClientConnect_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientConnect_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientConnect_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientConnect_Unregister(function callback) {
    using OnClientConnect_UnregisterFn = void (*)(function);
    static OnClientConnect_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientConnect_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnClientConnect_Post_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientConnect_Post_Register(function callback) {
    using OnClientConnect_Post_RegisterFn = void (*)(function);
    static OnClientConnect_Post_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientConnect_Post_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientConnect_Post_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientConnect_Post_Unregister(function callback) {
    using OnClientConnect_Post_UnregisterFn = void (*)(function);
    static OnClientConnect_Post_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientConnect_Post_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnClientConnected_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientConnected_Register(function callback) {
    using OnClientConnected_RegisterFn = void (*)(function);
    static OnClientConnected_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientConnected_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientConnected_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientConnected_Unregister(function callback) {
    using OnClientConnected_UnregisterFn = void (*)(function);
    static OnClientConnected_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientConnected_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnClientPutInServer_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientPutInServer_Register(function callback) {
    using OnClientPutInServer_RegisterFn = void (*)(function);
    static OnClientPutInServer_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientPutInServer_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientPutInServer_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientPutInServer_Unregister(function callback) {
    using OnClientPutInServer_UnregisterFn = void (*)(function);
    static OnClientPutInServer_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientPutInServer_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnClientDisconnect_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientDisconnect_Register(function callback) {
    using OnClientDisconnect_RegisterFn = void (*)(function);
    static OnClientDisconnect_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientDisconnect_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientDisconnect_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientDisconnect_Unregister(function callback) {
    using OnClientDisconnect_UnregisterFn = void (*)(function);
    static OnClientDisconnect_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientDisconnect_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnClientDisconnect_Post_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientDisconnect_Post_Register(function callback) {
    using OnClientDisconnect_Post_RegisterFn = void (*)(function);
    static OnClientDisconnect_Post_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientDisconnect_Post_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientDisconnect_Post_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientDisconnect_Post_Unregister(function callback) {
    using OnClientDisconnect_Post_UnregisterFn = void (*)(function);
    static OnClientDisconnect_Post_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientDisconnect_Post_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnClientActive_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientActive_Register(function callback) {
    using OnClientActive_RegisterFn = void (*)(function);
    static OnClientActive_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientActive_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientActive_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientActive_Unregister(function callback) {
    using OnClientActive_UnregisterFn = void (*)(function);
    static OnClientActive_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientActive_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnClientFullyConnect_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientFullyConnect_Register(function callback) {
    using OnClientFullyConnect_RegisterFn = void (*)(function);
    static OnClientFullyConnect_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientFullyConnect_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientFullyConnect_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientFullyConnect_Unregister(function callback) {
    using OnClientFullyConnect_UnregisterFn = void (*)(function);
    static OnClientFullyConnect_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientFullyConnect_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnClientSettingsChanged_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientSettingsChanged_Register(function callback) {
    using OnClientSettingsChanged_RegisterFn = void (*)(function);
    static OnClientSettingsChanged_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientSettingsChanged_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientSettingsChanged_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientSettingsChanged_Unregister(function callback) {
    using OnClientSettingsChanged_UnregisterFn = void (*)(function);
    static OnClientSettingsChanged_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientSettingsChanged_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnClientAuthenticated_Register
   * @param callback (function): Function callback.
   */
  inline void OnClientAuthenticated_Register(function callback) {
    using OnClientAuthenticated_RegisterFn = void (*)(function);
    static OnClientAuthenticated_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientAuthenticated_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnClientAuthenticated_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnClientAuthenticated_Unregister(function callback) {
    using OnClientAuthenticated_UnregisterFn = void (*)(function);
    static OnClientAuthenticated_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnClientAuthenticated_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnRoundTerminated_Register
   * @param callback (function): Function callback.
   */
  inline void OnRoundTerminated_Register(function callback) {
    using OnRoundTerminated_RegisterFn = void (*)(function);
    static OnRoundTerminated_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnRoundTerminated_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnRoundTerminated_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnRoundTerminated_Unregister(function callback) {
    using OnRoundTerminated_UnregisterFn = void (*)(function);
    static OnRoundTerminated_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnRoundTerminated_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnEntityCreated_Register
   * @param callback (function): Function callback.
   */
  inline void OnEntityCreated_Register(function callback) {
    using OnEntityCreated_RegisterFn = void (*)(function);
    static OnEntityCreated_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnEntityCreated_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnEntityCreated_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnEntityCreated_Unregister(function callback) {
    using OnEntityCreated_UnregisterFn = void (*)(function);
    static OnEntityCreated_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnEntityCreated_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnEntityDeleted_Register
   * @param callback (function): Function callback.
   */
  inline void OnEntityDeleted_Register(function callback) {
    using OnEntityDeleted_RegisterFn = void (*)(function);
    static OnEntityDeleted_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnEntityDeleted_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnEntityDeleted_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnEntityDeleted_Unregister(function callback) {
    using OnEntityDeleted_UnregisterFn = void (*)(function);
    static OnEntityDeleted_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnEntityDeleted_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnEntityParentChanged_Register
   * @param callback (function): Function callback.
   */
  inline void OnEntityParentChanged_Register(function callback) {
    using OnEntityParentChanged_RegisterFn = void (*)(function);
    static OnEntityParentChanged_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnEntityParentChanged_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnEntityParentChanged_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnEntityParentChanged_Unregister(function callback) {
    using OnEntityParentChanged_UnregisterFn = void (*)(function);
    static OnEntityParentChanged_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnEntityParentChanged_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnServerStartup_Register
   * @param callback (function): Function callback.
   */
  inline void OnServerStartup_Register(function callback) {
    using OnServerStartup_RegisterFn = void (*)(function);
    static OnServerStartup_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnServerStartup_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnServerStartup_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnServerStartup_Unregister(function callback) {
    using OnServerStartup_UnregisterFn = void (*)(function);
    static OnServerStartup_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnServerStartup_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnServerActivate_Register
   * @param callback (function): Function callback.
   */
  inline void OnServerActivate_Register(function callback) {
    using OnServerActivate_RegisterFn = void (*)(function);
    static OnServerActivate_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnServerActivate_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnServerActivate_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnServerActivate_Unregister(function callback) {
    using OnServerActivate_UnregisterFn = void (*)(function);
    static OnServerActivate_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnServerActivate_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnServerSpawn_Register
   * @param callback (function): Function callback.
   */
  inline void OnServerSpawn_Register(function callback) {
    using OnServerSpawn_RegisterFn = void (*)(function);
    static OnServerSpawn_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnServerSpawn_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnServerSpawn_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnServerSpawn_Unregister(function callback) {
    using OnServerSpawn_UnregisterFn = void (*)(function);
    static OnServerSpawn_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnServerSpawn_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnServerStarted_Register
   * @param callback (function): Function callback.
   */
  inline void OnServerStarted_Register(function callback) {
    using OnServerStarted_RegisterFn = void (*)(function);
    static OnServerStarted_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnServerStarted_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnServerStarted_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnServerStarted_Unregister(function callback) {
    using OnServerStarted_UnregisterFn = void (*)(function);
    static OnServerStarted_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnServerStarted_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnMapStart_Register
   * @param callback (function): Function callback.
   */
  inline void OnMapStart_Register(function callback) {
    using OnMapStart_RegisterFn = void (*)(function);
    static OnMapStart_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnMapStart_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnMapStart_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnMapStart_Unregister(function callback) {
    using OnMapStart_UnregisterFn = void (*)(function);
    static OnMapStart_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnMapStart_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnMapEnd_Register
   * @param callback (function): Function callback.
   */
  inline void OnMapEnd_Register(function callback) {
    using OnMapEnd_RegisterFn = void (*)(function);
    static OnMapEnd_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnMapEnd_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnMapEnd_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnMapEnd_Unregister(function callback) {
    using OnMapEnd_UnregisterFn = void (*)(function);
    static OnMapEnd_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnMapEnd_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnGameFrame_Register
   * @param callback (function): Function callback.
   */
  inline void OnGameFrame_Register(function callback) {
    using OnGameFrame_RegisterFn = void (*)(function);
    static OnGameFrame_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnGameFrame_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnGameFrame_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnGameFrame_Unregister(function callback) {
    using OnGameFrame_UnregisterFn = void (*)(function);
    static OnGameFrame_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnGameFrame_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnUpdateWhenNotInGame_Register
   * @param callback (function): Function callback.
   */
  inline void OnUpdateWhenNotInGame_Register(function callback) {
    using OnUpdateWhenNotInGame_RegisterFn = void (*)(function);
    static OnUpdateWhenNotInGame_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnUpdateWhenNotInGame_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnUpdateWhenNotInGame_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnUpdateWhenNotInGame_Unregister(function callback) {
    using OnUpdateWhenNotInGame_UnregisterFn = void (*)(function);
    static OnUpdateWhenNotInGame_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnUpdateWhenNotInGame_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Register callback to event.
   * @function OnPreWorldUpdate_Register
   * @param callback (function): Function callback.
   */
  inline void OnPreWorldUpdate_Register(function callback) {
    using OnPreWorldUpdate_RegisterFn = void (*)(function);
    static OnPreWorldUpdate_RegisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnPreWorldUpdate_Register", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Unregister callback to event.
   * @function OnPreWorldUpdate_Unregister
   * @param callback (function): Function callback.
   */
  inline void OnPreWorldUpdate_Unregister(function callback) {
    using OnPreWorldUpdate_UnregisterFn = void (*)(function);
    static OnPreWorldUpdate_UnregisterFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.OnPreWorldUpdate_Unregister", reinterpret_cast<void**>(&__func));
    __func(callback);
  }

  /**
   * @brief Retrieves the pointer to the current game rules proxy instance.
   * @function GetGameRulesProxy
   * @return ptr64: A pointer to the game rules entity instance.
   */
  inline void* GetGameRulesProxy() {
    using GetGameRulesProxyFn = void* (*)();
    static GetGameRulesProxyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameRulesProxy", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Retrieves the pointer to the current game rules instance.
   * @function GetGameRules
   * @return ptr64: A pointer to the game rules object.
   */
  inline void* GetGameRules() {
    using GetGameRulesFn = void* (*)();
    static GetGameRulesFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameRules", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Retrieves the team manager instance for a specified team.
   * @function GetGameTeamManager
   * @param team (int32): The numeric identifier of the team.
   * @return ptr64: A pointer to the corresponding CTeam instance, or nullptr if the team was not found.
   */
  inline void* GetGameTeamManager(int32_t team) {
    using GetGameTeamManagerFn = void* (*)(int32_t);
    static GetGameTeamManagerFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameTeamManager", reinterpret_cast<void**>(&__func));
    return __func(team);
  }

  /**
   * @brief Retrieves the current score of a specified team.
   * @function GetGameTeamScore
   * @param team (int32): The numeric identifier of the team.
   * @return int32: The current score of the team, or -1 if the team could not be found.
   */
  inline int32_t GetGameTeamScore(int32_t team) {
    using GetGameTeamScoreFn = int32_t (*)(int32_t);
    static GetGameTeamScoreFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameTeamScore", reinterpret_cast<void**>(&__func));
    return __func(team);
  }

  /**
   * @brief Retrieves the number of players on a specified team.
   * @function GetGamePlayerCount
   * @param team (int32): The numeric identifier of the team (e.g., CS_TEAM_T, CS_TEAM_CT, CS_TEAM_SPECTATOR).
   * @return int32: The number of players on the team, or -1 if game rules are unavailable.
   */
  inline int32_t GetGamePlayerCount(int32_t team) {
    using GetGamePlayerCountFn = int32_t (*)(int32_t);
    static GetGamePlayerCountFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGamePlayerCount", reinterpret_cast<void**>(&__func));
    return __func(team);
  }

  /**
   * @brief Returns the total number of rounds played in the current match.
   * @function GetGameTotalRoundsPlayed
   * @return int32: The total number of rounds played, or -1 if the game rules are unavailable.
   */
  inline int32_t GetGameTotalRoundsPlayed() {
    using GetGameTotalRoundsPlayedFn = int32_t (*)();
    static GetGameTotalRoundsPlayedFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetGameTotalRoundsPlayed", reinterpret_cast<void**>(&__func));
    return __func();
  }

  /**
   * @brief Forces the round to end with a specified reason and delay.
   * @function TerminateRound
   * @param delay (float): Time (in seconds) to delay before the next round starts.
   * @param reason (int32): The reason for ending the round, defined by the CSRoundEndReason enum.
   */
  inline void TerminateRound(float delay, int32_t reason) {
    using TerminateRoundFn = void (*)(float, int32_t);
    static TerminateRoundFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.TerminateRound", reinterpret_cast<void**>(&__func));
    __func(delay, reason);
  }

  /**
   * @brief Hooks a user message with a callback.
   * @function HookUserMessage
   * @param messageId (int16): The ID of the message to hook.
   * @param callback (function): The callback function to invoke when the message is received.
   * @param mode (uint8): Whether to hook the message in the post mode (after processing) or pre mode (before processing).
   * @return bool: True if the hook was successfully added, false otherwise.
   */
  inline bool HookUserMessage(int16_t messageId, function callback, uint8_t mode) {
    using HookUserMessageFn = bool (*)(int16_t, function, uint8_t);
    static HookUserMessageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.HookUserMessage", reinterpret_cast<void**>(&__func));
    return __func(messageId, callback, mode);
  }

  /**
   * @brief Unhooks a previously hooked user message.
   * @function UnhookUserMessage
   * @param messageId (int16): The ID of the message to unhook.
   * @param callback (function): The callback function to remove.
   * @param mode (uint8): Whether the hook was in post mode (after processing) or pre mode (before processing).
   * @return bool: True if the hook was successfully removed, false otherwise.
   */
  inline bool UnhookUserMessage(int16_t messageId, function callback, uint8_t mode) {
    using UnhookUserMessageFn = bool (*)(int16_t, function, uint8_t);
    static UnhookUserMessageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UnhookUserMessage", reinterpret_cast<void**>(&__func));
    return __func(messageId, callback, mode);
  }

  /**
   * @brief Creates a UserMessage from a serializable message.
   * @function UserMessageCreateFromSerializable
   * @param msgSerializable (ptr64): The serializable message.
   * @param message (ptr64): The network message.
   * @param recipientMask (uint64): The recipient mask.
   * @return ptr64: A pointer to the newly created UserMessage.
   */
  inline void* UserMessageCreateFromSerializable(void* msgSerializable, void* message, uint64_t recipientMask) {
    using UserMessageCreateFromSerializableFn = void* (*)(void*, void*, uint64_t);
    static UserMessageCreateFromSerializableFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageCreateFromSerializable", reinterpret_cast<void**>(&__func));
    return __func(msgSerializable, message, recipientMask);
  }

  /**
   * @brief Creates a UserMessage from a message name.
   * @function UserMessageCreateFromName
   * @param messageName (string): The name of the message.
   * @return ptr64: A pointer to the newly created UserMessage.
   */
  inline void* UserMessageCreateFromName(plg::string messageName) {
    using UserMessageCreateFromNameFn = void* (*)(plg::string);
    static UserMessageCreateFromNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageCreateFromName", reinterpret_cast<void**>(&__func));
    return __func(messageName);
  }

  /**
   * @brief Creates a UserMessage from a message ID.
   * @function UserMessageCreateFromId
   * @param messageId (int16): The ID of the message.
   * @return ptr64: A pointer to the newly created UserMessage.
   */
  inline void* UserMessageCreateFromId(int16_t messageId) {
    using UserMessageCreateFromIdFn = void* (*)(int16_t);
    static UserMessageCreateFromIdFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageCreateFromId", reinterpret_cast<void**>(&__func));
    return __func(messageId);
  }

  /**
   * @brief Destroys a UserMessage and frees its memory.
   * @function UserMessageDestroy
   * @param userMessage (ptr64): The UserMessage to destroy.
   */
  inline void UserMessageDestroy(void* userMessage) {
    using UserMessageDestroyFn = void (*)(void*);
    static UserMessageDestroyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageDestroy", reinterpret_cast<void**>(&__func));
    __func(userMessage);
  }

  /**
   * @brief Sends a UserMessage to the specified recipients.
   * @function UserMessageSend
   * @param userMessage (ptr64): The UserMessage to send.
   */
  inline void UserMessageSend(void* userMessage) {
    using UserMessageSendFn = void (*)(void*);
    static UserMessageSendFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageSend", reinterpret_cast<void**>(&__func));
    __func(userMessage);
  }

  /**
   * @brief Gets the name of the message.
   * @function UserMessageGetMessageName
   * @param userMessage (ptr64): The UserMessage instance.
   * @return string: The name of the message as a string.
   */
  inline plg::string UserMessageGetMessageName(void* userMessage) {
    using UserMessageGetMessageNameFn = plg::string (*)(void*);
    static UserMessageGetMessageNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageGetMessageName", reinterpret_cast<void**>(&__func));
    return __func(userMessage);
  }

  /**
   * @brief Gets the ID of the message.
   * @function UserMessageGetMessageID
   * @param userMessage (ptr64): The UserMessage instance.
   * @return int16: The ID of the message.
   */
  inline int16_t UserMessageGetMessageID(void* userMessage) {
    using UserMessageGetMessageIDFn = int16_t (*)(void*);
    static UserMessageGetMessageIDFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageGetMessageID", reinterpret_cast<void**>(&__func));
    return __func(userMessage);
  }

  /**
   * @brief Checks if the message has a specific field.
   * @function UserMessageHasField
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field to check.
   * @return bool: True if the field exists, false otherwise.
   */
  inline bool UserMessageHasField(void* userMessage, plg::string fieldName) {
    using UserMessageHasFieldFn = bool (*)(void*, plg::string);
    static UserMessageHasFieldFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageHasField", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName);
  }

  /**
   * @brief Gets the protobuf message associated with the UserMessage.
   * @function UserMessageGetProtobufMessage
   * @param userMessage (ptr64): The UserMessage instance.
   * @return ptr64: A pointer to the protobuf message.
   */
  inline void* UserMessageGetProtobufMessage(void* userMessage) {
    using UserMessageGetProtobufMessageFn = void* (*)(void*);
    static UserMessageGetProtobufMessageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageGetProtobufMessage", reinterpret_cast<void**>(&__func));
    return __func(userMessage);
  }

  /**
   * @brief Gets the serializable message associated with the UserMessage.
   * @function UserMessageGetSerializableMessage
   * @param userMessage (ptr64): The UserMessage instance.
   * @return ptr64: A pointer to the serializable message.
   */
  inline void* UserMessageGetSerializableMessage(void* userMessage) {
    using UserMessageGetSerializableMessageFn = void* (*)(void*);
    static UserMessageGetSerializableMessageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageGetSerializableMessage", reinterpret_cast<void**>(&__func));
    return __func(userMessage);
  }

  /**
   * @brief Finds a message ID by its name.
   * @function UserMessageFindMessageIdByName
   * @param messageName (string): The name of the message.
   * @return int16: The ID of the message, or 0 if the message was not found.
   */
  inline int16_t UserMessageFindMessageIdByName(plg::string messageName) {
    using UserMessageFindMessageIdByNameFn = int16_t (*)(plg::string);
    static UserMessageFindMessageIdByNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageFindMessageIdByName", reinterpret_cast<void**>(&__func));
    return __func(messageName);
  }

  /**
   * @brief Gets the recipient mask for the UserMessage.
   * @function UserMessageGetRecipientMask
   * @param userMessage (ptr64): The UserMessage instance.
   * @return uint64: The recipient mask.
   */
  inline uint64_t UserMessageGetRecipientMask(void* userMessage) {
    using UserMessageGetRecipientMaskFn = uint64_t (*)(void*);
    static UserMessageGetRecipientMaskFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageGetRecipientMask", reinterpret_cast<void**>(&__func));
    return __func(userMessage);
  }

  /**
   * @brief Adds a single recipient (player) to the UserMessage.
   * @function UserMessageAddRecipient
   * @param userMessage (ptr64): The UserMessage instance.
   * @param playerSlot (int32): The slot index of the player to add as a recipient.
   */
  inline void UserMessageAddRecipient(void* userMessage, int32_t playerSlot) {
    using UserMessageAddRecipientFn = void (*)(void*, int32_t);
    static UserMessageAddRecipientFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageAddRecipient", reinterpret_cast<void**>(&__func));
    __func(userMessage, playerSlot);
  }

  /**
   * @brief Adds all connected players as recipients to the UserMessage.
   * @function UserMessageAddAllPlayers
   * @param userMessage (ptr64): The UserMessage instance.
   */
  inline void UserMessageAddAllPlayers(void* userMessage) {
    using UserMessageAddAllPlayersFn = void (*)(void*);
    static UserMessageAddAllPlayersFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageAddAllPlayers", reinterpret_cast<void**>(&__func));
    __func(userMessage);
  }

  /**
   * @brief Sets the recipient mask for the UserMessage.
   * @function UserMessageSetRecipientMask
   * @param userMessage (ptr64): The UserMessage instance.
   * @param mask (uint64): The recipient mask to set.
   */
  inline void UserMessageSetRecipientMask(void* userMessage, uint64_t mask) {
    using UserMessageSetRecipientMaskFn = void (*)(void*, uint64_t);
    static UserMessageSetRecipientMaskFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageSetRecipientMask", reinterpret_cast<void**>(&__func));
    __func(userMessage, mask);
  }

  /**
   * @brief Gets a nested message from a field in the UserMessage.
   * @function UserMessageGetMessage
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param message (ptr64): A pointer to store the retrieved message.
   * @return bool: True if the message was successfully retrieved, false otherwise.
   */
  inline bool UserMessageGetMessage(void* userMessage, plg::string fieldName, void* message) {
    using UserMessageGetMessageFn = bool (*)(void*, plg::string, void*);
    static UserMessageGetMessageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageGetMessage", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, message);
  }

  /**
   * @brief Gets a repeated nested message from a field in the UserMessage.
   * @function UserMessageGetRepeatedMessage
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param message (ptr64): A pointer to store the retrieved message.
   * @return bool: True if the message was successfully retrieved, false otherwise.
   */
  inline bool UserMessageGetRepeatedMessage(void* userMessage, plg::string fieldName, int32_t index, void* message) {
    using UserMessageGetRepeatedMessageFn = bool (*)(void*, plg::string, int32_t, void*);
    static UserMessageGetRepeatedMessageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageGetRepeatedMessage", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, message);
  }

  /**
   * @brief Adds a nested message to a repeated field in the UserMessage.
   * @function UserMessageAddMessage
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param message (ptr64): A pointer to the message to add.
   * @return bool: True if the message was successfully added, false otherwise.
   */
  inline bool UserMessageAddMessage(void* userMessage, plg::string fieldName, void* message) {
    using UserMessageAddMessageFn = bool (*)(void*, plg::string, void*);
    static UserMessageAddMessageFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageAddMessage", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, message);
  }

  /**
   * @brief Gets the count of repeated fields in a field of the UserMessage.
   * @function UserMessageGetRepeatedFieldCount
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @return int32: The count of repeated fields, or -1 if the field is not repeated or does not exist.
   */
  inline int32_t UserMessageGetRepeatedFieldCount(void* userMessage, plg::string fieldName) {
    using UserMessageGetRepeatedFieldCountFn = int32_t (*)(void*, plg::string);
    static UserMessageGetRepeatedFieldCountFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageGetRepeatedFieldCount", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName);
  }

  /**
   * @brief Removes a value from a repeated field in the UserMessage.
   * @function UserMessageRemoveRepeatedFieldValue
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the value to remove.
   * @return bool: True if the value was successfully removed, false otherwise.
   */
  inline bool UserMessageRemoveRepeatedFieldValue(void* userMessage, plg::string fieldName, int32_t index) {
    using UserMessageRemoveRepeatedFieldValueFn = bool (*)(void*, plg::string, int32_t);
    static UserMessageRemoveRepeatedFieldValueFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageRemoveRepeatedFieldValue", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Gets the debug string representation of the UserMessage.
   * @function UserMessageGetDebugString
   * @param userMessage (ptr64): The UserMessage instance.
   * @return string: The debug string as a string.
   */
  inline plg::string UserMessageGetDebugString(void* userMessage) {
    using UserMessageGetDebugStringFn = plg::string (*)(void*);
    static UserMessageGetDebugStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.UserMessageGetDebugString", reinterpret_cast<void**>(&__func));
    return __func(userMessage);
  }

  /**
   * @brief Reads an enum value from a UserMessage.
   * @function PbReadEnum
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return int32: The integer representation of the enum value, or 0 if invalid.
   */
  inline int32_t PbReadEnum(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadEnumFn = int32_t (*)(void*, plg::string, int32_t);
    static PbReadEnumFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadEnum", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a 32-bit integer from a UserMessage.
   * @function PbReadInt32
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return int32: The int32_t value read, or 0 if invalid.
   */
  inline int32_t PbReadInt32(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadInt32Fn = int32_t (*)(void*, plg::string, int32_t);
    static PbReadInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a 64-bit integer from a UserMessage.
   * @function PbReadInt64
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return int64: The int64_t value read, or 0 if invalid.
   */
  inline int64_t PbReadInt64(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadInt64Fn = int64_t (*)(void*, plg::string, int32_t);
    static PbReadInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads an unsigned 32-bit integer from a UserMessage.
   * @function PbReadUInt32
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return uint32: The uint32_t value read, or 0 if invalid.
   */
  inline uint32_t PbReadUInt32(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadUInt32Fn = uint32_t (*)(void*, plg::string, int32_t);
    static PbReadUInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadUInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads an unsigned 64-bit integer from a UserMessage.
   * @function PbReadUInt64
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return uint64: The uint64_t value read, or 0 if invalid.
   */
  inline uint64_t PbReadUInt64(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadUInt64Fn = uint64_t (*)(void*, plg::string, int32_t);
    static PbReadUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadUInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a floating-point value from a UserMessage.
   * @function PbReadFloat
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return float: The float value read, or 0.0 if invalid.
   */
  inline float PbReadFloat(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadFloatFn = float (*)(void*, plg::string, int32_t);
    static PbReadFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadFloat", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a double-precision floating-point value from a UserMessage.
   * @function PbReadDouble
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return double: The double value read, or 0.0 if invalid.
   */
  inline double PbReadDouble(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadDoubleFn = double (*)(void*, plg::string, int32_t);
    static PbReadDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadDouble", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a boolean value from a UserMessage.
   * @function PbReadBool
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return bool: The boolean value read, or false if invalid.
   */
  inline bool PbReadBool(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadBoolFn = bool (*)(void*, plg::string, int32_t);
    static PbReadBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadBool", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a string from a UserMessage.
   * @function PbReadString
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return string: The string value read, or an empty string if invalid.
   */
  inline plg::string PbReadString(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadStringFn = plg::string (*)(void*, plg::string, int32_t);
    static PbReadStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadString", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a color value from a UserMessage.
   * @function PbReadColor
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return int32: The color value read, or an empty value if invalid.
   */
  inline int32_t PbReadColor(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadColorFn = int32_t (*)(void*, plg::string, int32_t);
    static PbReadColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadColor", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a 2D vector from a UserMessage.
   * @function PbReadVector2
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return vec2: The 2D vector value read, or an empty value if invalid.
   */
  inline plg::vec2 PbReadVector2(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadVector2Fn = plg::vec2 (*)(void*, plg::string, int32_t);
    static PbReadVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadVector2", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a 3D vector from a UserMessage.
   * @function PbReadVector3
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return vec3: The 3D vector value read, or an empty value if invalid.
   */
  inline plg::vec3 PbReadVector3(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadVector3Fn = plg::vec3 (*)(void*, plg::string, int32_t);
    static PbReadVector3Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadVector3", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Reads a QAngle (rotation vector) from a UserMessage.
   * @function PbReadQAngle
   * @param userMessage (ptr64): Pointer to the UserMessage object.
   * @param fieldName (string): Name of the field to read.
   * @param index (int32): Index of the repeated field (use -1 for non-repeated fields).
   * @return vec3: The QAngle value read, or an empty value if invalid.
   */
  inline plg::vec3 PbReadQAngle(void* userMessage, plg::string fieldName, int32_t index) {
    using PbReadQAngleFn = plg::vec3 (*)(void*, plg::string, int32_t);
    static PbReadQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbReadQAngle", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index);
  }

  /**
   * @brief Gets a enum value from a field in the UserMessage.
   * @function PbGetEnum
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetEnum(void* userMessage, plg::string fieldName, void* out) {
    using PbGetEnumFn = bool (*)(void*, plg::string, void*);
    static PbGetEnumFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetEnum", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a enum value for a field in the UserMessage.
   * @function PbSetEnum
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (int32): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetEnum(void* userMessage, plg::string fieldName, int32_t value) {
    using PbSetEnumFn = bool (*)(void*, plg::string, int32_t);
    static PbSetEnumFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetEnum", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a 32-bit integer value from a field in the UserMessage.
   * @function PbGetInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetInt32(void* userMessage, plg::string fieldName, void* out) {
    using PbGetInt32Fn = bool (*)(void*, plg::string, void*);
    static PbGetInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a 32-bit integer value for a field in the UserMessage.
   * @function PbSetInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (int32): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetInt32(void* userMessage, plg::string fieldName, int32_t value) {
    using PbSetInt32Fn = bool (*)(void*, plg::string, int32_t);
    static PbSetInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a 64-bit integer value from a field in the UserMessage.
   * @function PbGetInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetInt64(void* userMessage, plg::string fieldName, void* out) {
    using PbGetInt64Fn = bool (*)(void*, plg::string, void*);
    static PbGetInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a 64-bit integer value for a field in the UserMessage.
   * @function PbSetInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (int64): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetInt64(void* userMessage, plg::string fieldName, int64_t value) {
    using PbSetInt64Fn = bool (*)(void*, plg::string, int64_t);
    static PbSetInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets an unsigned 32-bit integer value from a field in the UserMessage.
   * @function PbGetUInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetUInt32(void* userMessage, plg::string fieldName, void* out) {
    using PbGetUInt32Fn = bool (*)(void*, plg::string, void*);
    static PbGetUInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetUInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets an unsigned 32-bit integer value for a field in the UserMessage.
   * @function PbSetUInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (uint32): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetUInt32(void* userMessage, plg::string fieldName, uint32_t value) {
    using PbSetUInt32Fn = bool (*)(void*, plg::string, uint32_t);
    static PbSetUInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetUInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets an unsigned 64-bit integer value from a field in the UserMessage.
   * @function PbGetUInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetUInt64(void* userMessage, plg::string fieldName, void* out) {
    using PbGetUInt64Fn = bool (*)(void*, plg::string, void*);
    static PbGetUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetUInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets an unsigned 64-bit integer value for a field in the UserMessage.
   * @function PbSetUInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (uint64): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetUInt64(void* userMessage, plg::string fieldName, uint64_t value) {
    using PbSetUInt64Fn = bool (*)(void*, plg::string, uint64_t);
    static PbSetUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetUInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a bool value from a field in the UserMessage.
   * @function PbGetBool
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetBool(void* userMessage, plg::string fieldName, void* out) {
    using PbGetBoolFn = bool (*)(void*, plg::string, void*);
    static PbGetBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetBool", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a bool value for a field in the UserMessage.
   * @function PbSetBool
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (bool): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetBool(void* userMessage, plg::string fieldName, bool value) {
    using PbSetBoolFn = bool (*)(void*, plg::string, bool);
    static PbSetBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetBool", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a float value from a field in the UserMessage.
   * @function PbGetFloat
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetFloat(void* userMessage, plg::string fieldName, void* out) {
    using PbGetFloatFn = bool (*)(void*, plg::string, void*);
    static PbGetFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetFloat", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a float value for a field in the UserMessage.
   * @function PbSetFloat
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (float): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetFloat(void* userMessage, plg::string fieldName, float value) {
    using PbSetFloatFn = bool (*)(void*, plg::string, float);
    static PbSetFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetFloat", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a double value from a field in the UserMessage.
   * @function PbGetDouble
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetDouble(void* userMessage, plg::string fieldName, void* out) {
    using PbGetDoubleFn = bool (*)(void*, plg::string, void*);
    static PbGetDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetDouble", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a double value for a field in the UserMessage.
   * @function PbSetDouble
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (double): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetDouble(void* userMessage, plg::string fieldName, double value) {
    using PbSetDoubleFn = bool (*)(void*, plg::string, double);
    static PbSetDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetDouble", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a string value from a field in the UserMessage.
   * @function PbGetString
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (string&): The output string.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetString(void* userMessage, plg::string fieldName, const plg::string& out) {
    using PbGetStringFn = bool (*)(void*, plg::string, const plg::string&);
    static PbGetStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetString", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a string value for a field in the UserMessage.
   * @function PbSetString
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (string): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetString(void* userMessage, plg::string fieldName, plg::string value) {
    using PbSetStringFn = bool (*)(void*, plg::string, plg::string);
    static PbSetStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetString", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a color value from a field in the UserMessage.
   * @function PbGetColor
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output string.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetColor(void* userMessage, plg::string fieldName, void* out) {
    using PbGetColorFn = bool (*)(void*, plg::string, void*);
    static PbGetColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetColor", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a color value for a field in the UserMessage.
   * @function PbSetColor
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (int32): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetColor(void* userMessage, plg::string fieldName, int32_t value) {
    using PbSetColorFn = bool (*)(void*, plg::string, int32_t);
    static PbSetColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetColor", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a Vector2 value from a field in the UserMessage.
   * @function PbGetVector2
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output string.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetVector2(void* userMessage, plg::string fieldName, void* out) {
    using PbGetVector2Fn = bool (*)(void*, plg::string, void*);
    static PbGetVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetVector2", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a Vector2 value for a field in the UserMessage.
   * @function PbSetVector2
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (vec2): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetVector2(void* userMessage, plg::string fieldName, plg::vec2 value) {
    using PbSetVector2Fn = bool (*)(void*, plg::string, plg::vec2);
    static PbSetVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetVector2", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a Vector3 value from a field in the UserMessage.
   * @function PbGetVector3
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output string.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetVector3(void* userMessage, plg::string fieldName, void* out) {
    using PbGetVector3Fn = bool (*)(void*, plg::string, void*);
    static PbGetVector3Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetVector3", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a Vector3 value for a field in the UserMessage.
   * @function PbSetVector3
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (vec3): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetVector3(void* userMessage, plg::string fieldName, plg::vec3 value) {
    using PbSetVector3Fn = bool (*)(void*, plg::string, plg::vec3);
    static PbSetVector3Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetVector3", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a QAngle value from a field in the UserMessage.
   * @function PbGetQAngle
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param out (ptr64): The output string.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetQAngle(void* userMessage, plg::string fieldName, void* out) {
    using PbGetQAngleFn = bool (*)(void*, plg::string, void*);
    static PbGetQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetQAngle", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, out);
  }

  /**
   * @brief Sets a QAngle value for a field in the UserMessage.
   * @function PbSetQAngle
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (vec3): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetQAngle(void* userMessage, plg::string fieldName, plg::vec3 value) {
    using PbSetQAngleFn = bool (*)(void*, plg::string, plg::vec3);
    static PbSetQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetQAngle", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated enum value from a field in the UserMessage.
   * @function PbGetRepeatedEnum
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedEnum(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedEnumFn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedEnumFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedEnum", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated enum value for a field in the UserMessage.
   * @function PbSetRepeatedEnum
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (int32): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedEnum(void* userMessage, plg::string fieldName, int32_t index, int32_t value) {
    using PbSetRepeatedEnumFn = bool (*)(void*, plg::string, int32_t, int32_t);
    static PbSetRepeatedEnumFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedEnum", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a enum value to a repeated field in the UserMessage.
   * @function PbAddEnum
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (int32): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddEnum(void* userMessage, plg::string fieldName, int32_t value) {
    using PbAddEnumFn = bool (*)(void*, plg::string, int32_t);
    static PbAddEnumFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddEnum", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated int32_t value from a field in the UserMessage.
   * @function PbGetRepeatedInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedInt32(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedInt32Fn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated int32_t value for a field in the UserMessage.
   * @function PbSetRepeatedInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (int32): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedInt32(void* userMessage, plg::string fieldName, int32_t index, int32_t value) {
    using PbSetRepeatedInt32Fn = bool (*)(void*, plg::string, int32_t, int32_t);
    static PbSetRepeatedInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a 32-bit integer value to a repeated field in the UserMessage.
   * @function PbAddInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (int32): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddInt32(void* userMessage, plg::string fieldName, int32_t value) {
    using PbAddInt32Fn = bool (*)(void*, plg::string, int32_t);
    static PbAddInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated int64_t value from a field in the UserMessage.
   * @function PbGetRepeatedInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedInt64(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedInt64Fn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated int64_t value for a field in the UserMessage.
   * @function PbSetRepeatedInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (int64): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedInt64(void* userMessage, plg::string fieldName, int32_t index, int64_t value) {
    using PbSetRepeatedInt64Fn = bool (*)(void*, plg::string, int32_t, int64_t);
    static PbSetRepeatedInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a 64-bit integer value to a repeated field in the UserMessage.
   * @function PbAddInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (int64): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddInt64(void* userMessage, plg::string fieldName, int64_t value) {
    using PbAddInt64Fn = bool (*)(void*, plg::string, int64_t);
    static PbAddInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated uint32_t value from a field in the UserMessage.
   * @function PbGetRepeatedUInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedUInt32(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedUInt32Fn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedUInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedUInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated uint32_t value for a field in the UserMessage.
   * @function PbSetRepeatedUInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (uint32): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedUInt32(void* userMessage, plg::string fieldName, int32_t index, uint32_t value) {
    using PbSetRepeatedUInt32Fn = bool (*)(void*, plg::string, int32_t, uint32_t);
    static PbSetRepeatedUInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedUInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds an unsigned 32-bit integer value to a repeated field in the UserMessage.
   * @function PbAddUInt32
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (uint32): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddUInt32(void* userMessage, plg::string fieldName, uint32_t value) {
    using PbAddUInt32Fn = bool (*)(void*, plg::string, uint32_t);
    static PbAddUInt32Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddUInt32", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated uint64_t value from a field in the UserMessage.
   * @function PbGetRepeatedUInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedUInt64(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedUInt64Fn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedUInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated uint64_t value for a field in the UserMessage.
   * @function PbSetRepeatedUInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (uint64): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedUInt64(void* userMessage, plg::string fieldName, int32_t index, uint64_t value) {
    using PbSetRepeatedUInt64Fn = bool (*)(void*, plg::string, int32_t, uint64_t);
    static PbSetRepeatedUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedUInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds an unsigned 64-bit integer value to a repeated field in the UserMessage.
   * @function PbAddUInt64
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (uint64): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddUInt64(void* userMessage, plg::string fieldName, uint64_t value) {
    using PbAddUInt64Fn = bool (*)(void*, plg::string, uint64_t);
    static PbAddUInt64Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddUInt64", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated bool value from a field in the UserMessage.
   * @function PbGetRepeatedBool
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedBool(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedBoolFn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedBool", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated bool value for a field in the UserMessage.
   * @function PbSetRepeatedBool
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (bool): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedBool(void* userMessage, plg::string fieldName, int32_t index, bool value) {
    using PbSetRepeatedBoolFn = bool (*)(void*, plg::string, int32_t, bool);
    static PbSetRepeatedBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedBool", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a bool value to a repeated field in the UserMessage.
   * @function PbAddBool
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (bool): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddBool(void* userMessage, plg::string fieldName, bool value) {
    using PbAddBoolFn = bool (*)(void*, plg::string, bool);
    static PbAddBoolFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddBool", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated float value from a field in the UserMessage.
   * @function PbGetRepeatedFloat
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedFloat(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedFloatFn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedFloat", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated float value for a field in the UserMessage.
   * @function PbSetRepeatedFloat
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (float): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedFloat(void* userMessage, plg::string fieldName, int32_t index, float value) {
    using PbSetRepeatedFloatFn = bool (*)(void*, plg::string, int32_t, float);
    static PbSetRepeatedFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedFloat", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a float value to a repeated field in the UserMessage.
   * @function PbAddFloat
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (float): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddFloat(void* userMessage, plg::string fieldName, float value) {
    using PbAddFloatFn = bool (*)(void*, plg::string, float);
    static PbAddFloatFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddFloat", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated double value from a field in the UserMessage.
   * @function PbGetRepeatedDouble
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output value.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedDouble(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedDoubleFn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedDouble", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated double value for a field in the UserMessage.
   * @function PbSetRepeatedDouble
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (double): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedDouble(void* userMessage, plg::string fieldName, int32_t index, double value) {
    using PbSetRepeatedDoubleFn = bool (*)(void*, plg::string, int32_t, double);
    static PbSetRepeatedDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedDouble", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a double value to a repeated field in the UserMessage.
   * @function PbAddDouble
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (double): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddDouble(void* userMessage, plg::string fieldName, double value) {
    using PbAddDoubleFn = bool (*)(void*, plg::string, double);
    static PbAddDoubleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddDouble", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated string value from a field in the UserMessage.
   * @function PbGetRepeatedString
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (string&): The output string.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedString(void* userMessage, plg::string fieldName, int32_t index, const plg::string& out) {
    using PbGetRepeatedStringFn = bool (*)(void*, plg::string, int32_t, const plg::string&);
    static PbGetRepeatedStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedString", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated string value for a field in the UserMessage.
   * @function PbSetRepeatedString
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (string): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedString(void* userMessage, plg::string fieldName, int32_t index, plg::string value) {
    using PbSetRepeatedStringFn = bool (*)(void*, plg::string, int32_t, plg::string);
    static PbSetRepeatedStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedString", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a string value to a repeated field in the UserMessage.
   * @function PbAddString
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (string): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddString(void* userMessage, plg::string fieldName, plg::string value) {
    using PbAddStringFn = bool (*)(void*, plg::string, plg::string);
    static PbAddStringFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddString", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated color value from a field in the UserMessage.
   * @function PbGetRepeatedColor
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output color.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedColor(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedColorFn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedColor", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated color value for a field in the UserMessage.
   * @function PbSetRepeatedColor
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (int32): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedColor(void* userMessage, plg::string fieldName, int32_t index, int32_t value) {
    using PbSetRepeatedColorFn = bool (*)(void*, plg::string, int32_t, int32_t);
    static PbSetRepeatedColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedColor", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a color value to a repeated field in the UserMessage.
   * @function PbAddColor
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (int32): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddColor(void* userMessage, plg::string fieldName, int32_t value) {
    using PbAddColorFn = bool (*)(void*, plg::string, int32_t);
    static PbAddColorFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddColor", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated Vector2 value from a field in the UserMessage.
   * @function PbGetRepeatedVector2
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output vector2.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedVector2(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedVector2Fn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedVector2", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated Vector2 value for a field in the UserMessage.
   * @function PbSetRepeatedVector2
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (vec2): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedVector2(void* userMessage, plg::string fieldName, int32_t index, plg::vec2 value) {
    using PbSetRepeatedVector2Fn = bool (*)(void*, plg::string, int32_t, plg::vec2);
    static PbSetRepeatedVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedVector2", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a Vector2 value to a repeated field in the UserMessage.
   * @function PbAddVector2
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (vec2): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddVector2(void* userMessage, plg::string fieldName, plg::vec2 value) {
    using PbAddVector2Fn = bool (*)(void*, plg::string, plg::vec2);
    static PbAddVector2Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddVector2", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated Vector3 value from a field in the UserMessage.
   * @function PbGetRepeatedVector3
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output vector2.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedVector3(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedVector3Fn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedVector3Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedVector3", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated Vector3 value for a field in the UserMessage.
   * @function PbSetRepeatedVector3
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (vec3): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedVector3(void* userMessage, plg::string fieldName, int32_t index, plg::vec3 value) {
    using PbSetRepeatedVector3Fn = bool (*)(void*, plg::string, int32_t, plg::vec3);
    static PbSetRepeatedVector3Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedVector3", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a Vector3 value to a repeated field in the UserMessage.
   * @function PbAddVector3
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (vec3): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddVector3(void* userMessage, plg::string fieldName, plg::vec3 value) {
    using PbAddVector3Fn = bool (*)(void*, plg::string, plg::vec3);
    static PbAddVector3Fn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddVector3", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Gets a repeated QAngle value from a field in the UserMessage.
   * @function PbGetRepeatedQAngle
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param out (ptr64): The output vector2.
   * @return bool: True if the field was successfully retrieved, false otherwise.
   */
  inline bool PbGetRepeatedQAngle(void* userMessage, plg::string fieldName, int32_t index, void* out) {
    using PbGetRepeatedQAngleFn = bool (*)(void*, plg::string, int32_t, void*);
    static PbGetRepeatedQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbGetRepeatedQAngle", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, out);
  }

  /**
   * @brief Sets a repeated QAngle value for a field in the UserMessage.
   * @function PbSetRepeatedQAngle
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param index (int32): The index of the repeated field.
   * @param value (vec3): The value to set.
   * @return bool: True if the field was successfully set, false otherwise.
   */
  inline bool PbSetRepeatedQAngle(void* userMessage, plg::string fieldName, int32_t index, plg::vec3 value) {
    using PbSetRepeatedQAngleFn = bool (*)(void*, plg::string, int32_t, plg::vec3);
    static PbSetRepeatedQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbSetRepeatedQAngle", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, index, value);
  }

  /**
   * @brief Adds a QAngle value to a repeated field in the UserMessage.
   * @function PbAddQAngle
   * @param userMessage (ptr64): The UserMessage instance.
   * @param fieldName (string): The name of the field.
   * @param value (vec3): The value to add.
   * @return bool: True if the value was successfully added, false otherwise.
   */
  inline bool PbAddQAngle(void* userMessage, plg::string fieldName, plg::vec3 value) {
    using PbAddQAngleFn = bool (*)(void*, plg::string, plg::vec3);
    static PbAddQAngleFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.PbAddQAngle", reinterpret_cast<void**>(&__func));
    return __func(userMessage, fieldName, value);
  }

  /**
   * @brief Retrieves the weapon VData for a given weapon name.
   * @function GetWeaponVDataFromKey
   * @param name (string): The name of the weapon.
   * @return ptr64: A pointer to the `CCSWeaponBaseVData` if the entity handle is valid and represents a player weapon; otherwise, nullptr.
   */
  inline void* GetWeaponVDataFromKey(plg::string name) {
    using GetWeaponVDataFromKeyFn = void* (*)(plg::string);
    static GetWeaponVDataFromKeyFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetWeaponVDataFromKey", reinterpret_cast<void**>(&__func));
    return __func(name);
  }

  /**
   * @brief Retrieves the weapon VData for a given weapon.
   * @function GetWeaponVData
   * @param entityHandle (int32): The handle of the entity from which to retrieve the weapon VData.
   * @return ptr64: A pointer to the `CCSWeaponBaseVData` if the entity handle is valid and represents a player weapon; otherwise, nullptr.
   */
  inline void* GetWeaponVData(int32_t entityHandle) {
    using GetWeaponVDataFn = void* (*)(int32_t);
    static GetWeaponVDataFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetWeaponVData", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the weapon type of a given entity.
   * @function GetWeaponType
   * @param entityHandle (int32): The handle of the entity (weapon).
   * @return uint32: The type of the weapon, or WEAPONTYPE_UNKNOWN if the entity is invalid.
   */
  inline uint32_t GetWeaponType(int32_t entityHandle) {
    using GetWeaponTypeFn = uint32_t (*)(int32_t);
    static GetWeaponTypeFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetWeaponType", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the weapon category of a given entity.
   * @function GetWeaponCategory
   * @param entityHandle (int32): The handle of the entity (weapon).
   * @return uint32: The category of the weapon, or WEAPONCATEGORY_OTHER if the entity is invalid.
   */
  inline uint32_t GetWeaponCategory(int32_t entityHandle) {
    using GetWeaponCategoryFn = uint32_t (*)(int32_t);
    static GetWeaponCategoryFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetWeaponCategory", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the gear slot of a given weapon entity.
   * @function GetWeaponGearSlot
   * @param entityHandle (int32): The handle of the entity (weapon).
   * @return uint32: The gear slot of the weapon, or GEAR_SLOT_INVALID if the entity is invalid.
   */
  inline uint32_t GetWeaponGearSlot(int32_t entityHandle) {
    using GetWeaponGearSlotFn = uint32_t (*)(int32_t);
    static GetWeaponGearSlotFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetWeaponGearSlot", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the weapon definition index for a given entity handle.
   * @function GetWeaponItemDefinition
   * @param entityHandle (int32): The handle of the entity from which to retrieve the weapon def index.
   * @return uint16: The weapon definition index as a `uint16_t`, or 0 if the entity handle is invalid.
   */
  inline uint16_t GetWeaponItemDefinition(int32_t entityHandle) {
    using GetWeaponItemDefinitionFn = uint16_t (*)(int32_t);
    static GetWeaponItemDefinitionFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetWeaponItemDefinition", reinterpret_cast<void**>(&__func));
    return __func(entityHandle);
  }

  /**
   * @brief Retrieves the item definition index associated with a given item name.
   * @function GetWeaponItemDefinitionByName
   * @param itemName (string): The name of the item.
   * @return uint16: The weapon definition index as a `uint16_t`, or 0 if the entity handle is invalid.
   */
  inline uint16_t GetWeaponItemDefinitionByName(plg::string itemName) {
    using GetWeaponItemDefinitionByNameFn = uint16_t (*)(plg::string);
    static GetWeaponItemDefinitionByNameFn __func = nullptr;
    if (__func == nullptr) plg::GetMethodPtr2("${S2SDK_PACKAGE}.GetWeaponItemDefinitionByName", reinterpret_cast<void**>(&__func));
    return __func(itemName);
  }

} // namespace ${S2SDK_PACKAGE}
