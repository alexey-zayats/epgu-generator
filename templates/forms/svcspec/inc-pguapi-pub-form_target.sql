-- Author: Eremenko
-- Change: {{.Change}}
-- Description:
begin

{{- $targetID := .ServiceTargetID}}
{{- range .UnlinkService}}{{$len := len .}}{{ if ne $len 0}}
pgu.pkg_service2.UnLink_eservice2target(
    i_eservice => {{.}},
    i_target => {{$targetID}});
{{- end}}{{end}}

  PGU.PKG_SRVFUTILS.create_service(
    extid                  => '{{.DepartmentCode}}',
    categories             => '{{.ApplicantType}}',
    es_id                  => {{.ServiceFormCode}},
    targets                => PGU.PKG_SRVFUTILS.T$TARGETS({{.ServiceTargetID}}),
    fnames                 => PGU.PKG_SRVFUTILS.T$NAMES('{{.ServiceTargetName}}'),
    fshortnames            => PGU.PKG_SRVFUTILS.T$NAMES('{{.ServiceTargetName}}'),
    fp                     => 'pguformsbeta',
    edsstatus_code         => '{{.Signature}}',
    eservice_online        => '0',
    active                 => 'Y',
    OKATO                  => '60000000000', --просьба проверить корректность кода ОКАТО
    published              => 1,
    edstype                => 'V_3_0',
    spec_code              => 'MR',
    spec_type              => 'CALL',
    soap_action            => null,
    ftl_template           => null,
    ftl_response           => null,
    transport              => 'SOAP_11',
    uddikey                => '{{.ServiceFormCode}}',
    bundle_code            => 'top10',
    uddikey_region_search  => 'N',
    asynch_response        => 'N',
    status_notchange       => 'N',
    log_status             => 'Y',
    icon_name              => null,
    org_status_list        => null
  );
commit;
end;
/
