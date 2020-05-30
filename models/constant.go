package models

const (
	AuthLogin                            = "auth.login"
	AuthLogout                           = "auth.logout"
	AuthTokenList                        = "auth.token_list"
	AuthTokenAdd                         = "auth.token_add"
	AuthTokenGenerate                    = "auth.token_generate"
	AuthTokenRemove                      = "auth.token_remove"
	ConsoleCreate                        = "console.create"
	ConsoleList                          = "console.list"
	ConsoleDestroy                       = "console.destroy"
	ConsoleRead                          = "console.read"
	ConsoleWrite                         = "console.write"
	ConsoleTabs                          = "console.tabs"
	ConsoleSessionKill                   = "console.session_kill"
	ConsoleSessionDetach                 = "console.session_detach"
	CoreVersion                          = "core.version"
	CoreStop                             = "core.stop"
	CoreSetG                             = "core.setg"
	CoreUnsetG                           = "core.unsetg"
	CoreSave                             = "core.save"
	CoreReloadModules                    = "core.reload_modules"
	CoreModuleStats                      = "core.module_stats"
	CoreAddModulePath                    = "core.add_module_path"
	CoreThreadList                       = "core.thread_list"
	CoreThreadKill                       = "core.thread_kill"
	DbHosts                              = "db.hosts"
	DbServices                           = "db.services"
	DbVulns                              = "db.vulns"
	DbWorkspaces                         = "db.workspaces"
	DbCurrentWorkspace                   = "db.current_workspace"
	DbGetWorkspace                       = "db.get_workspace"
	DbSetWorkspace                       = "db.set_workspace"
	DbDelWorkspace                       = "db.del_workspace"
	DbAddWorkspace                       = "db.add_workspace"
	DbGetHost                            = "db.get_host"
	DbReportHost                         = "db.report_host"
	DbReportService                      = "db.report_service"
	DbGetService                         = "db.get_service"
	DbGetNote                            = "db.get_note"
	DbGetClient                          = "db.get_client"
	DbReportClient                       = "db.report_client"
	DbReportNote                         = "db.report_note"
	DbNotes                              = "db.notes"
	DbReportAuthInfo                     = "db.report_auth_info"
	DbGetAuthInfo                        = "db.get_auth_info"
	DbGetRef                             = "db.get_ref"
	DbDelVuln                            = "db.del_vuln"
	DbDelNote                            = "db.del_note"
	DbDelService                         = "db.del_service"
	DbDelHost                            = "db.del_host"
	DbReportVuln                         = "db.report_vuln"
	DbEvents                             = "db.events"
	DbReportEvent                        = "db.report_event"
	DbReportLoot                         = "db.report_loot"
	DbLoots                              = "db.loots"
	DbReportCred                         = "db.report_cred"
	DbCreds                              = "db.creds"
	DbImportData                         = "db.import_data"
	DbGetVuln                            = "db.get_vuln"
	DbClients                            = "db.clients"
	DbDelClient                          = "db.del_client"
	DbDriver                             = "db.driver"
	DbConnect                            = "db.connect"
	DbStatus                             = "db.status"
	DbDisconnect                         = "db.disconnect"
	JobList                              = "job.list"
	JobStop                              = "job.stop"
	JobInfo                              = "job.info"
	ModuleExploits                       = "module.exploits"
	ModuleAuxiliary                      = "module.auxiliary"
	ModulePayloads                       = "module.payloads"
	ModuleEncoders                       = "module.encoders"
	ModuleNops                           = "module.nops"
	ModulePost                           = "module.post"
	ModuleInfo                           = "module.info"
	ModuleCompatiblePayloads             = "module.compatible_payloads"
	ModuleCompatibleSessions             = "module.compatible_sessions"
	ModuleTargetCompatiblePayloads       = "module.target_compatible_payloads"
	ModuleOptions                        = "module.options"
	ModuleExecute                        = "module.execute"
	ModuleEncodeFormats                  = "module.encode_formats"
	ModuleEncode                         = "module.encode"
	PluginLoad                           = "plugin.load"
	PluginUnload                         = "plugin.unload"
	PluginLoaded                         = "plugin.loaded"
	SessionList                          = "session.list"
	SessionStop                          = "session.stop"
	SessionShellRead                     = "session.shell_read"
	SessionShellWrite                    = "session.shell_write"
	SessionShellUpgrade                  = "session.shell_upgrade"
	SessionMeterpreterRead               = "session.meterpreter_read"
	SessionRingRead                      = "session.ring_read"
	SessionRingPut                       = "session.ring_put"
	SessionRingLast                      = "session.ring_last"
	SessionRingClear                     = "session.ring_clear"
	SessionMeterpreterWrite              = "session.meterpreter_write"
	SessionMeterpreterSessionDetach      = "session.meterpreter_session_detach"
	SessionMeterpreterSessionKill        = "session.meterpreter_session_kill"
	SessionMeterpreterTabs               = "session.meterpreter_tabs"
	SessionMeterpreterRunSingle          = "session.meterpreter_run_single"
	SessionMeterpreterScript             = "session.meterpreter_script"
	SessionMeterpreterDirectorySeparator = "session.meterpreter_directory_separator"
	SessionCompatibleModules             = "session.compatible_modules"
)
