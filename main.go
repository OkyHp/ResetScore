package main

import (
	"runtime/debug"

	t "github.com/OkyHp/plg_utils/translation"
	s2 "github.com/fr0nch/go-plugify-s2sdk/v2"
	"github.com/untrustedmodders/go-plugify"
)

type ResetScorePlugin struct {
}

func initDefaultValues() *ResetScorePlugin {
	return &ResetScorePlugin{}
}

var Plugin *ResetScorePlugin

func init() {
	//utils.CreateManifest("ResetScore", "1.0.0", "OkyHek", []string{"s2sdk"})

	Plugin = initDefaultValues()

	plugify.OnPluginStart(Plugin.OnPluginStart)
	plugify.OnPluginEnd(Plugin.OnPluginEnd)
	plugify.OnPluginPanic(Plugin.OnPluginPanic)
}

func (rs *ResetScorePlugin) OnPluginStart() {
	// ConVarFlag_LinkedConcommand - Declarate commad
	// FCVAR_SERVER_CAN_EXECUTE - the server is allowed to execute this command on clients via ClientCommand/NET_StringCmd/CBaseClientState::ProcessStringCmd.
	// ConVarFlag_Release - Cvars tagged with this are the only cvars avaliable to customers
	var flags = s2.ConVarFlag_LinkedConcommand | s2.ConVarFlag_Release | s2.ConVarFlag_ClientCanExecute
	s2.AddConsoleCommand("rs", "", flags, rs.onResetScore, s2.HookMode_Post)
	s2.AddConsoleCommand("кі", "", flags, rs.onResetScore, s2.HookMode_Post)
	s2.AddConsoleCommand("кы", "", flags, rs.onResetScore, s2.HookMode_Post)

	err := t.LoadTranslation("reset_score")
	if err != nil {
		panic(err)
	}

	s2.OnServerActivate_Register(rs.OnServerActivate)

	Debug("Plugin ResetScore successfully loaded.")
}

func (rs *ResetScorePlugin) OnPluginEnd() {
	s2.RemoveCommand("rs", rs.onResetScore)
	s2.RemoveCommand("кы", rs.onResetScore)
	s2.RemoveCommand("кі", rs.onResetScore)

	s2.OnServerActivate_Unregister(rs.OnServerActivate)

	Debug("Plugin ResetScore stopped.")
}

func (rs *ResetScorePlugin) OnPluginPanic() []byte {

	return debug.Stack() // workaround for could not import runtime/debug inside plugify package
}

func (rs *ResetScorePlugin) OnServerActivate() { // it`s OnMapStart
	s2.ServerCommand("mp_backup_round_file \"\"")
	s2.ServerCommand("mp_backup_round_file_last \"\"")
	s2.ServerCommand("mp_backup_round_file_pattern \"\"")
	s2.ServerCommand("mp_backup_round_auto 0")
}

func (rs *ResetScorePlugin) onResetScore(playerSlot int32, context s2.CommandCallingContext, arguments []string) s2.ResultType {
	if !s2.IsClientInGame(playerSlot) || s2.IsFakeClient(playerSlot) || s2.IsClientSourceTV(playerSlot) {
		return s2.ResultType_Continue
	}

	playerHandle := s2.PlayerSlotToEntHandle(playerSlot)
	notify := false

	resetIfNonZero := func(get func(int32) int32, set func(int32, int32)) {
		if get(playerSlot) != 0 {
			set(playerSlot, 0)
			notify = true
		}
	}

	resetIfNonZero(s2.GetClientKills, s2.SetClientKills)
	resetIfNonZero(s2.GetClientAssists, s2.SetClientAssists)
	resetIfNonZero(s2.GetClientDeaths, s2.SetClientDeaths)
	resetIfNonZero(s2.GetClientDamage, s2.SetClientDamage)

	resetSchemaIfNonZero := func(field string) {
		if s2.GetEntSchema(playerHandle, "CCSPlayerController", field, 0) != 0 {
			s2.SetEntSchema(playerHandle, "CCSPlayerController", field, 0, true, 0)
			notify = true
		}
	}

	resetSchemaIfNonZero("m_iMVPs")
	resetSchemaIfNonZero("m_iScore")

	// Сообщение
	msgKey := "fail_reset"
	if notify {
		msgKey = "success_reset"
	}
	s2.PrintToChat(playerSlot, " "+t.GetTranslation(playerSlot, msgKey))

	return s2.ResultType_Continue
}

func main() {}
