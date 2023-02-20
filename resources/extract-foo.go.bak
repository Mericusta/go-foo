//go:build none

package resources

import (
	extractorfoo "go-foo/extractor-foo"

	"github.com/Cyinx/einx"
)

type renameInt int

type ExtractGoStruct struct {
	i   int
	f   float64
	s   string
	sr  []rune
	msb map[string]byte
	ps  *ExtractGoStruct
	ip  *extractorfoo.GoVariableDefinition
	// as  struct {
	// 	*ExtractGoStruct
	// 	*extractorfoo.GoStructMemberDefinition
	// }
}

type EmptyExtractGoStruct struct {
}

type OneLineEmptyExtractGoStruct struct{}

var commentAndDivide1 int = 1 / 2
var commentAndDivide2 int = 1 / 2

type GoInterfaceDeclaration interface {
	FunctionDeclaration1()
	FunctionDeclaration2() int32
	FunctionDeclaration3(int) int32
	FunctionDeclaration4(int, int64)
	FunctionDeclaration5(int, int64) int
	FunctionDeclaration6(int, int64) (int, int64)
	FunctionDeclaration7(int, int64) (int, int64)
	FunctionDeclaration8(int, func() int32) (int, int64)
	FunctionDeclaration9(*extractorfoo.GoVariableDefinition, func(int, int) int32) (*extractorfoo.GoVariableDefinition, int64)
	FunctionDeclaration0()
	// FunctionDeclaration9(*extractorfoo.GoStructMemberDefinition, func() int32) (*extractorfoo.GoStructMemberDefinition, int64)
}

type DBModule = einx.Module
type WeaponItemInfo = cargo_def.WeaponItemInfo
type TaskData = cargo_def.TaskData
type ItemInfo = cargo_def.ItemInfo
type RoleItemInfo = cargo_def.RoleItemInfo
type WeaponGrowCostItem = msg_def.WeaponGrowCostItem
type ElfInfo = cargo_def.ElfInfo
type ElfEquip = cargo_def.ElfEquipInfo
type ElfExpCostItems = msg_def.ElfExpCostItem
type WorldCfgData = gameCfg.WorldCfgData
type FightRoleInfo = cargo_def.RoleFightItemInfo

type IAccount interface {
	SendMsg(interface{})
	RouteRpc(int, ...interface{})
	GetAccountID() uint64
	GetGateGUID() string
	DBService() DBModule
	CenterService() DBModule
}

type IEnvParams interface {
	Add(int32, interface{})
	Get(int32) (interface{}, bool)
	Remove(int32)
	Clear()
}

type IPlayer interface {
	GetID() uint64
	GetAccount() IAccount
	GetPlayerData() *cargo_def.PlayerData
	MarshalPlayerData() *cargo_def.PlayerData
	IncrVersion() int64
	SendMsg(interface{})
	ForceSendMsg(interface{})
	GetDataComponent() IPlayerDataComponent
	GetTask() ITaskComponent
	GetBag() IBagComponent
	GetReportMail() IReportMailComponent
	GetKV() IKVComponent
	GetBBS() IBBSComponent
	GetWeapon() IWeaponComponent
	GetRole() IRoleComponent
	GetChip() IChipComponent
	GetRecords() IRecordsComponent
	GetGP() IGamePhoneComponent
	GetMR() IDungeonComponent
	GetReward() IRewardComponent
	GetPrestige() IPrestigeComponent
	GetActivity() IActivityComponent
	GetAchievement() IAchievementComponent
	GetScene() ISceneComponent
	GetShop() IShopComponent
	GetMail() IMailComponent
	GetELF() IELFComponent
	GetAsyncInteraction() IAsyncInteractionComponent
	GetMapScene() IMapComponent
	GetGameStateComponent() IGameStateComponent
	GetServerNodeComponent() IServerNodeComponent
	GetTime() ITimeComponent
	GetFriendComponent() IFriendComponent
	GetCollect() ICollectInfo
	GetBook() IBooksComponent
	GetFlashGame() IFlashGame
	GetPray() IPrayComponent

	LogBI() *LogAction
	OnRequest()
	UpdateDirty()
	DataDirty() bool
	GenGUID(int32) uint64
	GetLastResult() (int32, bool)
	LastResult(int32)
	ClearResult()
	SaveData()

	GetEnv() IEnvParams
	ExitEnv(IEnvParams)

	// AddTimer delay millisecond
	AddTimer(uint64, GameTimerHandler, ...interface{}) uint64
	RemoveTimer(uint64) bool
}

type IBagComponent interface {
	MultiAddItem(map[int32]int64) int32
	AddItem(int32, int64) (map[int32]int64, int32)
	UseItem(int32, int64) (map[int32]int64, int32)
	RemoveItem(int32, int64) int32
	GetItemCount(int32) int64
	GetItemCountByGUID(uint64) int64
	RemoveItemByGUID(uint64, int64) int32
	GetItemInfoByGUID(uint64) (*ItemInfo, bool)
	UpdateDirtyMsg()
	GetItemsByClassID(int32) []*ItemInfo
	ResolveItem(uint64, int64) int32
	GetItemRecord(int32) int64
}

type ITaskInfo interface {
	GetTaskID() int32
	GetTaskVar(int32, int32) int64
	SetPolicy(int32, int32) bool
	GetPolicyID() (int32, int32)
	CompleteTask()
}

type ITaskComponent interface {
	GetMataShim() IMataShim
	AcceptTask(int32, int32) int32
	AutoAccept(int32) int32
	IsTaskPolicyEqualCondition(int32, int64, int32) bool
	DeliveryTask(int32) int32
	DeleteDeliveryTask(int32) int32
	ForceDeliveryTask(int32) int32
	IsTaskComplete(int32) bool
	ForceClearTask(int32)
	SetTaskPolicy(int32, int32, int32)
	HaveTaskAccepted(int32) bool
	DailyClean()
	WeeklyClean()
	GetCycleTaskCount(int32) (int32, int32)
	ClearTaskRelatedFlag(int32) int32
	ForceDeliveryAndCompleteTask(int32) int32
	ModifyTaskTraceID(int32)
	IsTaskRunning(int32) bool
	TaskSkip(int32, int32) int32
}

type IPlayerDataComponent interface {
	GetRewardMeta(int32) int32
	IncrRewardMeta(int32, int32)
	IsRewardMeta(int32, int32) bool
	ClearRewardMeta(int32)

	EventTrigger(string) int32

	AddExp(int32) int32
	GetLevelExp() (int32, int32)

	GetEnergy() int32
	AddEnergy(int32) int32
	PayEnergy(int32) int32

	GetMoney() int64
	PayMoney(int64) int32
	AddMoney(int64) int32

	AddRankDate(int32, int64)

	SavePlayerModel(map[int32]int32) int32
	DirtyHeadIcon()

	CompleteDrama(int32) int32
	IsDramaComplete(int32) bool

	PhotoPicture(int32) int32
	IsPhotoPicture(int32) bool
}

type ITimeComponent interface {
	ModifyPlayTime(int32)
	GetPlayTime() int32
	SetCdkTime(int32)
	GetExchangeCDkTime() int32
}

type IReportMailComponent interface {
	PushMail(int32) int32
	OnMailReport(uint64) int32
	OpMail(uint64, int32) int32
}

type IKVComponent interface {
	StoreByID(int32, int32) int32
	RemoveByID(int32) int32
	GetValueByID(int32) (int32, int32)
	SyncDirtyFlag()
	ExecEvents(int32) int32
	NewCounter(int32) int32
	RemoveCounter(int32) int32
	AddCounterValue(int32, int32) int32
	GetCounterValue(int32) (int32, bool)
	SyncDirtyCounter()
}

type IGamePhoneComponent interface {
	AddGp(int32, int32) int32
	RemoveGP(int32)
	AddGpReply(int32, int32, int32) (int32, int32)
	SetTop(int32, bool) int32
	SetGPStatus(int32, int32) int32
	IsComplete(int32) bool
	IsRead(int32) bool
}

type IWeaponComponent interface {
	UpdateDirtyMsg()
	CreateWeaPonItem(int32) (*WeaponItemInfo, bool)
	WeaponsGrow(uint64, []*WeaponGrowCostItem) int32
	WeaponsStar(uint64) int32
	RemoveWeapon(uint64)
	WeaponsAwaken(uint64, []uint64) int32
	WeaponsReplace(uint64, uint64, int32) int32
	GetWeaponAttribute(uint64, *AttrCalc)
	WeaponsLock(uint64) int32
	GetWeapon(uint64) (*WeaponItemInfo, bool)
	GetFightWeapon(map[int32]uint64, *cargo_def.RoleFightItemInfo)
}

type IBBSComponent interface {
	AddArticleData(int32) int32
	RemoveArticle(int32) int32
	CommentArticle(int32, int32, int32) (*cargo_def.BBSCommentData, int32)
	// ScanTask(int32) (int32, int32)
	Star(int32) (int32, int32)
	// TriggerEvent(int32) int32
	MarkArticleRead(int32) int32
	// IsArticleComplete(int32) bool
	GetBBSDataInfo(int32) (*cargo_def.BBSArticleData, bool)
}

type IRoleComponent interface {
	AddRole(int32) (uint64, int32)
	AddFightRoleExp(int32)
	RoleBreak(uint64) int32
	RoleAwaken(uint64) int32
	// SuitUpWeapon(uint64, int32, uint64) int32
	DirtyUpdate()
	GetRole(uint64) (*RoleItemInfo, bool)
	AddExpByGUID(uint64, int32) int32
	ExchangeWeapon(uint64, uint64, int32) int32
	GetRoleChips(uint64, int32) (uint64, bool)
	Check(uint64) int32
	MultiCheck([]uint64) int32
	PrepareFightRole(int32, uint32, string, []uint64, bool) int32
	GetRoleFightGroup(int32) ([]uint64, bool)
	RangeRoleByClass(int32, func(*RoleItemInfo) bool)
	ForeachRole(func(*RoleItemInfo) bool)
	//CalcAttr(*RoleItemInfo)
	SetNowFightRole(int32)
	// AddFightRoleELFTrustExp(int32, int32)
	AddFightRoleELFExp(int32)
	UseExpItem(uint64, map[uint64]int64) int32
	GetFightRole(int32) []*FightRoleInfo
	GetEndMissionRoleData() map[uint64]int32
	ClearFightRoles()
	ActiveRoleTalent(uint64, int32) int32
	InstallELF(uint64, int32) int32
	InstallChip(uint64, int32, uint64) int32
	UninstallChip(uint64, int32) int32
	GetRoleByClass(int32) (*RoleItemInfo, bool)
	UnlockRole(int32) int32
}

type IRecordsComponent interface {
	AddTaskMark(int32, int32, int32, int32)
}

type IDungeonComponent interface {
	BeginDungeon(int32) int32
	EndDungeon(int32, map[int32]int64, map[uint64]*cargo_def.DungeonEndRoleInfo) (map[int32]*cargo_def.RewardItem, []bool, int32, int32)
	GetEndDungeonConditionValue(int32) (int32, int64)
	ForeachEndDungeonRoleInfo() (int32, map[uint64]*cargo_def.DungeonEndRoleInfo)
	EndDungeonIsFinish() (int32, int32)
	GetDungeonSectionData(int32) *cargo_def.DungeonBaseInfo
	ClearDungeonData(int32) int32
}

type IChipComponent interface {
	Check(uint64) int32
	MultiCheck([]uint64) int32
	Create(int32) (uint64, int32)
	GetState(...uint64) map[uint64]IChipState
	GetInfo(uint64) (IChipInfo, int32)
	MarkUse(uint64, uint64) int32
	MarkUnuse(uint64) int32
	Lock(uint64) int32
	Unlock(uint64) int32
	CalculateGrow(uint64, []uint64, map[uint64]int64) (int32, int32, int32, int64, []uint64, map[int32]int64)
	Destroy(...uint64) int32
	SetLevelExp(uint64, int32, int32) int32
	SetSkill(uint64, []int32) int32
	GetChipAttribute(uint64, *AttrCalc)
	SyncDirtyChip()
	GetChipData(uint64) (*cargo_def.ChipItemInfo, int32)
	GetFightChip(map[int32]uint64, *cargo_def.RoleFightItemInfo)
}

type IRewardComponent interface {
	GetRewardItemByPoolID(int32) (map[int32]int64, int32)
}

type IPrestigeComponent interface {
	CreatePrestige(int32) int32
	GetPrestigeLevelExp(int32, bool) (int32, int32, int32)

	CreateExtraPrestige(int32, int64) int32
	RemoveExtraPrestige(int32, int64) int32

	AddExp(int32, int32) int32
	GetPrestigeMoney(int32) int64
	UsePrestigeMoney(int32, int64) int32

	SyncDirtyPrestige()
}

type IActivityComponent interface {
	IsCanJoinActivity(int32) int32
	GetRewardBySignIn(int32) int32
	SignIn(int32) int32
}

type IAchievementComponent interface {
	GetRewardByAchievement(int32) int32
	GetChapterReward(int32) int32
	UnlockAchieveRecord(int32, int32) int32
	ServerProcessAchievement(int32, int32)
	ClientProcessAchievement(int32, int32)
}

type ISceneComponent interface {
	Create(int32, int32) int32
	OnLoginScene(string) int32
}

type IShopComponent interface {
	UpdateShop(int32) (int32, bool)
	GetShopData(int32) (*cargo_def.ShopData, int32)
	BuyInShop(int32, int32, int32, int32, uint64) int32
	SyncDirtyShop()
}

type IMailComponent interface {
	OnLoadMail([]*cargo_def.MailData)
	LoadMail()
	OpMail(int64, int32) int32
}

type IELFComponent interface {
	GetELFEntity(int32) (*ElfInfo, bool)
	GetELFDebrisCount(int32) int32
	AddELFDebrisCount(int32, int32)

	CreateElf(int32) int32
	CreateChip(int32) uint64
	GetELFAttribute(int32, *AttrCalc)
	GetElf(int32) (*ElfInfo, bool)
	GetElfEquip(uint64) (*ElfEquip, bool)
	// AddElfTrustValue(int32, int32)
	AddElfExpValue(int32, int32) int32
	InstallChip(int32, uint64, int32) int32
	UninstallChip(int32, uint64, int32) int32
	HandlerAddElfExpValue(int32, []*ElfExpCostItems) int32
	ELFAwakenUp(int32) int32
	ResolveElfEquip(uint64) int32
	UpdateDirtyElfMsg()
	UpdateDirtyElfEquipMsg()
	GetELFTrustLevel(int32) int32
	GetFightElf(int32, *cargo_def.RoleFightItemInfo)
	ElfEquipGrow(uint64, []*msg_def.ElfEquipGrowCostItem) int32
	ElfEquipStar(uint64, uint64) int32
	ElfEquipLock(uint64, bool) int32
	ElfMarkUse(int32, uint64) int32
	ElfMarkUnuse(int32) int32
}

type IAsyncInteractionComponent interface {
	GenerateAsyncInteractionA(map[int32]int32, uint64) (*cargo_def.AsyncInteractionA, int32)
	AfterStoreAsyncInteractionA(*cargo_def.AsyncInteractionA) int32
	AcceptAsyncInteractionA(uint64) int32
	AfterAcceptAsyncInteractionA(*cargo_def.AsyncInteractionA) int32
	CompleteAsyncInteractionA(uint64) int32
	AfterCompleteAsyncInteractionA(*cargo_def.AsyncInteractionA) int32
}

type ISceneCtx interface {
	GetOrNewWeatherGameTime(int32, int32) int32
}

type IScene interface {
	Init(IPlayer, ISceneCtx, int32, *WorldCfgData)
	UnInit()
	GetID() int32
	GetType() int32
	OnEnter() int32
	OnLeave() int32
	GetWeather() IWeatherComponent
}

type IMapComponent interface {
	OnLoginComplete()
	RequestEnter(int32) int32
	InitMap(int32)
	GetScene() IScene
	// NPC
	SetNPCMood(int32, int32) (int32, int32)
	GetNPCMood(int32) (int32, int32)
	AddNPCMood(int32, int32) (int32, int32)
	SubNPCMood(int32, int32) (int32, int32)
	StartNPCFollow(int32) int32
	StopNPCFollow() int32
	SyncDirtyNPC()
}

type IWeatherComponent interface {
	AddWeather(int32) int32
	RemoveWeather(int32) int32
	Add(int32, int32, int32) int32
	Clear(int32, int32) int32
	AddIntensity(int32, int32, int32) int32
}

type IGameStateComponent interface {
	EnterSafeCity(int32) int32
	EnterDungeon(int32, int32, uint32, bool) int32
	EnterBirth(int32) int32
	EnterFight(int32) int32
	EnterBarrage(int32) int32
	EnterDiffChecker(int32) int32
	EnterStory(int32) int32
	EnterInstanceCity(int32) int32
	EnterPuzzleGame(int32) int32
	StackTop() (GameWorldState, bool)
	LeaveTop() int32
	SetFlag(int32, int32) int32
	RemoveFlag(int32) int32
	SyncDirtyFlag()
}

type IServerNodeComponent interface {
	ExecuteServerNode(int64, int32, int32) int32
}

type IFriendComponent interface {
	AddFriendRequest([]uint64) int32
	AddFriend(uint64) int32
	RemoveFriend(uint64) int32
	BlockPlayer(uint64) (int32, bool)
	UnblockPlayer(uint64) (int32, bool)
	OnLoadFriendData([]uint64, []uint64)
	OnLoadRequestMap([]uint64)
	IsInBlockList(uint64) bool
	IsInFriendList(uint64) bool
}

type IChatComponent interface {
}

type ICollectInfo interface {
	AddCollection(int32)
	IsCompleteCollectId(int32) bool
	GetCollectGroup(int32) int32
}

type IBooksComponent interface {
	AddBooks(int32, bool) int32
	GetBooksReward(int32) int32
	ReadBook(int32)
}

type GWStackMgr interface {
	LocalPlayer() IPlayer
	StackCount() int32
	StackTop() (GameWorldState, bool)
	ForeachStack(func(int, GameWorldState) bool)
	StackEntry(int, GameWorldState) int32
}

type GameWorldState interface {
	Init(GWStackMgr)
	ID() uint64
	Prepare() bool
	OnEnter() int32
	ReEnter() int32
	OnComplete()
	OnExit()
	Type() int32
	Transition(GameWorldState) (int32, int32)
	IsBlocked() (int32, int32)
	SetFlag(int32, int32) int32
	RemoveFlag(int32) int32
	SyncDirtyFlag()
}

type IFlashGame interface {
	SaveBarrageRecord(int32, int32)
	SavePuzzleRecord(int32, int32)
	GetPuzzleRecord(int32) (int32, bool)
	SaveDiffCheckerRecord(int32)
	GetDiffCheckerRecord(int32) int32
	AddQteRecord(int32, bool)
	GetQteRecord(int32) int32
}

type IPrayComponent interface {
	CheckPrayState([]int32) []int32
	GetMissCountAndNextGuaranteeNumber(int32) (int32, int32, int32)
	Pray(int32, int32) ([]*msg_def.SinglePray, int32)
	GetPrayGroupCount(int32) int32
}

func ExtractFunction(a, b int, c []string, d *msg_def.SinglePray, e func(int, int) map[string]interface{}, f func() map[string]interface{}) ([]*msg_def.SinglePray, func(struct{ v interface{} }) interface{}, []interface{}) {
	return nil, nil
}

func (this *ThisStruct) ExtractMemberFunction(a, b int, c []string, d *msg_def.SinglePray) func(struct{ v interface{} }) interface{} {
	return nil
}
