import { useState } from 'react';
import { useCreateResult } from 'api/hooks/resultsHooks';
import CreateButton from 'components/buttons/CreateButton';
import withAuth from 'hocs/withAuth';
import { CreateResultDto } from 'types/dtos/requests/results/createResultDto';
import { Permission } from 'types/enums/permission';
import { getToday } from 'utils/dateUtils';
import BaseResultMutateDialog from './BaseResultMutateDialog';

const initialResult: CreateResultDto = {
  leagueId: '',
  startDate: getToday(),
  endDate: getToday(),
  teamId: '',
  opponentTeamId: '',
  score: '',
  opponentScore: '',
};

function CreateResultButton() {
  const [open, setOpen] = useState(false);

  const { mutateAsync: createResultAsync } = useCreateResult();

  const handleSubmitAsync = async (payload: CreateResultDto) => {
    await createResultAsync(payload);
  };

  return (
    <>
      <CreateButton
        title="Create Result"
        canCreate
        onClick={() => {
          setOpen(true);
        }}
      />

      <BaseResultMutateDialog
        title="Create a new Result"
        buttonText="Create"
        result={initialResult}
        submitAsync={handleSubmitAsync}
        open={open}
        onClose={() => {
          setOpen(false);
        }}
      />
    </>
  );
}

export default withAuth(CreateResultButton, {
  permission: Permission.ManageResult,
  redirect: false,
});
