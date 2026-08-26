package internal

import (
	"fmt"

	s2sdk "github.com/fr0nch/go-plugify-s2sdk/v2"
	translations "github.com/fr0nch/go-plugify-translations"
	"github.com/fr0nch/logger"
	"github.com/untrustedmodders/go-plugify"
)

type ResetScorePlugin struct {
	Plugin plugify.Plugin

	mvpsOffset  int32
	scoreOffset int32
	chainOffset int32

	log *logger.Logger
}

func NewResetScorePlugin() *ResetScorePlugin {
	return &ResetScorePlugin{}
}

func (r *ResetScorePlugin) OnPluginStart() error {
	var err error

	r.log, err = logger.NewWithOptions(logger.Options{PluginName: r.Plugin.Name()})
	if err != nil {
		return err
	}

	var flags = s2sdk.ConVarFlag_LinkedConcommand | s2sdk.ConVarFlag_Release | s2sdk.ConVarFlag_ClientCanExecute

	s2sdk.AddConsoleCommand("rs", "", flags, r.onResetScore, s2sdk.HookMode_Post)
	s2sdk.AddConsoleCommand("кі", "", flags, r.onResetScore, s2sdk.HookMode_Post)
	s2sdk.AddConsoleCommand("кы", "", flags, r.onResetScore, s2sdk.HookMode_Post)

	errString := translations.LoadTranslation([]string{
		r.Plugin.Location() + "/resetscore.yml",
		plugify.ConfigsDir() + "/resetscore.yml",
		plugify.DataDir() + "/resetscore.yml",
	})
	if errString != "" {
		return fmt.Errorf("translations err: %s", errString)
	}

	s2sdk.OnServerActivate_Register(r.OnServerActivate)

	r.log.Debug("Plugin ResetScore successfully loaded.")

	return nil
}

func (r *ResetScorePlugin) OnPluginUpdate(dt float32) error {
	return nil
}

func (r *ResetScorePlugin) OnPluginEnd() error {
	r.log.Debug("Plugin ResetScore stopped.")
	return nil
}

func (r *ResetScorePlugin) OnServerActivate() {
	r.mvpsOffset = s2sdk.GetSchemaOffset("CCSPlayerController", "m_iMVPs")
	r.scoreOffset = s2sdk.GetSchemaOffset("CCSPlayerController", "m_iScore")
	r.chainOffset = s2sdk.GetSchemaChainOffset("CCSPlayerController")

	s2sdk.ServerCommand("mp_backup_round_file \"\"")
	s2sdk.ServerCommand("mp_backup_round_file_last \"\"")
	s2sdk.ServerCommand("mp_backup_round_file_pattern \"\"")
	s2sdk.ServerCommand("mp_backup_round_auto 0")
}

func (r *ResetScorePlugin) onResetScore(playerSlot int32, context s2sdk.ConCommandContext, arguments []string) s2sdk.ResultType {
	if !s2sdk.IsClientInGame(playerSlot) || s2sdk.IsFakeClient(playerSlot) || s2sdk.IsClientSourceTV(playerSlot) {
		return s2sdk.ResultType_Handled
	}

	playerHandle := s2sdk.PlayerSlotToEntHandle(playerSlot)
	notify := false

	reset := func(get func(int32) int32, set func(int32, int32)) {
		if get(playerSlot) != 0 {
			set(playerSlot, 0)
			notify = true
		}
	}

	reset(s2sdk.GetClientKills, s2sdk.SetClientKills)
	reset(s2sdk.GetClientAssists, s2sdk.SetClientAssists)
	reset(s2sdk.GetClientDeaths, s2sdk.SetClientDeaths)
	reset(s2sdk.GetClientDamage, s2sdk.SetClientDamage)

	if s2sdk.GetEntData(playerHandle, r.mvpsOffset, 4) != 0 {
		s2sdk.SetEntData(playerHandle, r.mvpsOffset, 0, 4, true, r.chainOffset)
		notify = true
	}
	if s2sdk.GetEntData(playerHandle, r.scoreOffset, 4) != 0 {
		s2sdk.SetEntData(playerHandle, r.scoreOffset, 0, 4, true, r.chainOffset)
		notify = true
	}

	var msg string
	lang := s2sdk.GetClientLanguageCode(playerSlot)

	if notify {
		msg = translations.Translate(lang, "success_reset")
	} else {
		msg = translations.Translate(lang, "fail_reset")
	}

	s2sdk.PrintToChatColored(playerSlot, msg)

	return s2sdk.ResultType_Continue
}
