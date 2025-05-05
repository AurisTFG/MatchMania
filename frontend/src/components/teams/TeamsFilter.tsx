import { useFetchTeams } from 'api/hooks/teamsHooks';
import { SELECT_OPTIONS } from 'constants/selectOptions';
import { useAppForm } from 'hooks/form/useAppForm';

type TeamsFilterProps = {
  onTeamChange: (teamId: string) => void;
  teamId?: string;
};

export default function TeamsFilter({
  onTeamChange,
  teamId = '',
}: TeamsFilterProps) {
  const { data: teams } = useFetchTeams();

  const teamOptions =
    teams?.map((team) => ({
      key: team.id,
      value: team.name,
    })) ?? [];

  const form = useAppForm({
    defaultValues: {
      teamId: teamId,
    },
  });

  return (
    <form.AppField
      name="teamId"
      listeners={{
        onChange: ({ value }) => {
          onTeamChange(value);
        },
      }}
    >
      {(field) => (
        <field.Select
          label="Filter by Team"
          options={teamOptions}
          notSelectedOption={SELECT_OPTIONS.ALL_TEAMS}
        />
      )}
    </form.AppField>
  );
}
