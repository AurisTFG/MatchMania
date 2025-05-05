import { useState } from 'react';
import { useUpdateTeam } from 'api/hooks/teamsHooks';
import EditButton from 'components/buttons/UpdateButton';
import withAuth from 'hocs/withAuth';
import { UpdateTeamDto } from 'types/dtos/requests/teams/updateTeamDto';
import { TeamDto } from 'types/dtos/responses/teams/teamDto';
import { Permission } from 'types/enums/permission';
import BaseTeamMutateDialog from './BaseTeamMutateDialog';

type UpdateTeamButtonProps = {
  team: TeamDto;
};

function UpdateTeamButton({ team }: UpdateTeamButtonProps) {
  const [open, setOpen] = useState(false);

  const { mutateAsync: updateTeamAsync } = useUpdateTeam();

  const handleSubmitAsync = async (payload: UpdateTeamDto) => {
    await updateTeamAsync({ teamId: team.id, payload });
  };

  const remappedTeam = {
    ...team,
    leagueIds: team.leagues.map((league) => league.id),
    playerIds: team.players.map((player) => player.id),
  } as UpdateTeamDto;

  return (
    <>
      <EditButton
        onClick={() => {
          setOpen(true);
        }}
      />

      <BaseTeamMutateDialog
        title="Edit Team"
        buttonText="Save Changes"
        team={remappedTeam}
        submitAsync={handleSubmitAsync}
        open={open}
        onClose={() => {
          setOpen(false);
        }}
      />
    </>
  );
}

export default withAuth(UpdateTeamButton, {
  permission: Permission.ManageTeam,
  dataOwnerUserId: (props) => (props as UpdateTeamButtonProps).team.user.id,
  redirect: false,
});
