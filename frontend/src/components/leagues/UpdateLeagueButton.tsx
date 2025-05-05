import dayjs from 'dayjs';
import { useState } from 'react';
import { useUpdateLeague } from 'api/hooks/leaguesHooks';
import UpdateButton from 'components/buttons/UpdateButton';
import withAuth from 'hocs/withAuth';
import { UpdateLeagueDto } from 'types/dtos/requests/leagues/updateLeagueDto';
import { LeagueDto } from 'types/dtos/responses/leagues/leagueDto';
import { Permission } from 'types/enums/permission';
import BaseLeagueMutateDialog from './BaseLeagueMutateDialog';

type UpdateLeagueButtonProps = {
  league: LeagueDto;
};

function UpdateLeagueButton({ league }: UpdateLeagueButtonProps) {
  const [open, setOpen] = useState(false);

  const { mutateAsync: updateLeagueAsync } = useUpdateLeague();

  const handleSubmitAsync = async (payload: UpdateLeagueDto) => {
    await updateLeagueAsync({ leagueId: league.id, payload });
  };

  const remappedLeague = {
    ...league,
    startDate: dayjs(league.startDate).toDate(),
    endDate: dayjs(league.endDate).toDate(),
    trackIds: league.tracks.map((track) => track.id),
  } as UpdateLeagueDto;

  return (
    <>
      <UpdateButton
        onClick={() => {
          setOpen(true);
        }}
      />

      <BaseLeagueMutateDialog
        title="Edit League"
        buttonText="Save Changes"
        league={remappedLeague}
        submitAsync={handleSubmitAsync}
        open={open}
        onClose={() => {
          setOpen(false);
        }}
      />
    </>
  );
}

export default withAuth(UpdateLeagueButton, {
  permission: Permission.ManageLeague,
  dataOwnerUserId: (props) => (props as UpdateLeagueButtonProps).league.user.id,
  redirect: false,
});
