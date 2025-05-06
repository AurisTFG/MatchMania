import { useState } from 'react';
import { useCreateTeam } from 'api/hooks/teamsHooks';
import CreateButton from 'components/buttons/CreateButton';
import withAuth from 'hocs/withAuth';
import { CreateTeamDto } from 'types/dtos/requests/teams/createTeamDto';
import { Permission } from 'types/enums/permission';
import BaseTeamMutateDialog from './BaseTeamMutateDialog';

function CreateTeamButton() {
  const [open, setOpen] = useState(false);

  const { mutateAsync: createTeamAsync } = useCreateTeam();

  const handleSubmitAsync = async (payload: CreateTeamDto) => {
    await createTeamAsync(payload);
  };

  const initialTeam: CreateTeamDto = {
    name: '',
    logoUrl: '',
    leagueIds: [],
    playerIds: [],
  };

  return (
    <>
      <CreateButton
        title="Create Team"
        canCreate
        onClick={() => {
          setOpen(true);
        }}
      />

      <BaseTeamMutateDialog
        title="Create a new Team"
        buttonText="Create"
        team={initialTeam}
        submitAsync={handleSubmitAsync}
        open={open}
        onClose={() => {
          setOpen(false);
        }}
      />
    </>
  );
}

export default withAuth(CreateTeamButton, {
  permission: Permission.ManageTeam,
  redirect: false,
});
