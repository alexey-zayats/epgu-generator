-- Author: Eremenko
-- Change: {{.Change}} rollback
-- Description:
begin
 pgu.pkg_service2.UnLink_eservice2target(
    i_eservice => {{.ServiceFormCode}},
    i_target => {{.ServiceTargetID}});
commit;
end;
/
