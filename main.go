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
	if s2.IsClientInGame(playerSlot) && !s2.IsFakeClient(playerSlot) && !s2.IsClientSourceTV(playerSlot) {
		playerHandle := s2.PlayerSlotToEntHandle(playerSlot)

		notify := false

		if s2.GetClientKills(playerSlot) != 0 {
			s2.SetClientKills(playerSlot, 0)
			notify = true
		}

		if s2.GetClientAssists(playerSlot) != 0 {
			s2.SetClientAssists(playerSlot, 0)
			notify = true
		}

		if s2.GetClientDeaths(playerSlot) != 0 {
			s2.SetClientDeaths(playerSlot, 0)
			notify = true
		}

		if s2.GetClientDamage(playerSlot) != 0 {
			s2.SetClientDamage(playerSlot, 0)
			notify = true
		}

		if s2.GetEntSchema(playerHandle, "CCSPlayerController", "m_iMVPs", 0) != 0 {
			s2.SetEntSchema(playerHandle, "CCSPlayerController", "m_iMVPs", 0, true, 0)
			notify = true
		}

		if s2.GetEntSchema(playerHandle, "CCSPlayerController", "m_iScore", 0) != 0 {
			s2.SetEntSchema(playerHandle, "CCSPlayerController", "m_iScore", 0, true, 0)
			notify = true
		}

		if notify {
			s2.PrintToChat(playerSlot, " "+t.GetTranslation(playerSlot, "success_reset"))
		} else {
			s2.PrintToChat(playerSlot, " "+t.GetTranslation(playerSlot, "fail_reset"))
		}
	}

	return s2.ResultType_Continue
}

func main() {}
