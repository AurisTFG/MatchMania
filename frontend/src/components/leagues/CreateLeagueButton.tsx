import { useState } from 'react';
import { useCreateLeague } from 'api/hooks/leaguesHooks';
import CreateButton from 'components/buttons/CreateButton';
import withAuth from 'hocs/withAuth';
import { CreateLeagueDto } from 'types/dtos/requests/leagues/createLeagueDto';
import { Permission } from 'types/enums/permission';
import { getStartOfDay } from 'utils/dateUtils';
import BaseLeagueMutateDialog from './BaseLeagueMutateDialog';

const initialLeague: CreateLeagueDto = {
  name: '',
  startDate: getStartOfDay(),
  endDate: getStartOfDay(7),
  trackIds: [],
};

function CreateLeagueButton() {
  const [open, setOpen] = useState(false);

  const { mutateAsync: createLeagueAsync } = useCreateLeague();

  const handleSubmitAsync = async (payload: CreateLeagueDto) => {
    await createLeagueAsync(payload);
  };

  return (
    <>
      <CreateButton
        title="Create League"
        canCreate
        onClick={() => {
          setOpen(true);
        }}
      />

      <BaseLeagueMutateDialog
        title="Create a new League"
        buttonText="Create"
        league={initialLeague}
        submitAsync={handleSubmitAsync}
        open={open}
        onClose={() => {
          setOpen(false);
        }}
      />
    </>
  );
}

export default withAuth(CreateLeagueButton, {
  permission: Permission.ManageLeague,
  redirect: false,
});
