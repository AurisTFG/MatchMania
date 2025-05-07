import dayjs from 'dayjs';
import { useState } from 'react';
import { useUpdateResult } from 'api/hooks/resultsHooks';
import UpdateButton from 'components/buttons/UpdateButton';
import withAuth from 'hocs/withAuth';
import { UpdateResultDto } from 'types/dtos/requests/results/updateResultDto';
import { ResultDto } from 'types/dtos/responses/results/resultDto';
import { Permission } from 'types/enums/permission';
import BaseResultMutateDialog from './BaseResultMutateDialog';

type UpdateResultButtonProps = {
  result: ResultDto;
};

function UpdateResultButton({ result }: UpdateResultButtonProps) {
  const [open, setOpen] = useState(false);

  const { mutateAsync: updateResultAsync } = useUpdateResult();

  const handleSubmitAsync = async (payload: UpdateResultDto) => {
    await updateResultAsync({ resultId: result.id, payload });
  };

  const remappedResult = {
    ...result,
    startDate: dayjs(result.startDate).toDate(),
    endDate: dayjs(result.endDate).toDate(),
    leagueId: result.league.id,
    teamId: result.team.id,
    opponentTeamId: result.opponentTeam.id,
    score: result.score.toString(),
    opponentScore: result.opponentScore.toString(),
  } as UpdateResultDto;

  return (
    <>
      <UpdateButton
        onClick={() => {
          setOpen(true);
        }}
      />

      <BaseResultMutateDialog
        title="Edit Result"
        buttonText="Save Changes"
        result={remappedResult}
        submitAsync={handleSubmitAsync}
        open={open}
        onClose={() => {
          setOpen(false);
        }}
      />
    </>
  );
}

export default withAuth(UpdateResultButton, {
  permission: Permission.ManageResult,
  dataOwnerUserId: (props) => (props as UpdateResultButtonProps).result.user.id,
  redirect: false,
});
