import { useFetchLeagues } from 'api/hooks/leaguesHooks';
import { SELECT_OPTIONS } from 'constants/selectOptions';
import { useAppForm } from 'hooks/form/useAppForm';

type LeaguesFilterProps = {
  onLeagueChange: (leagueId: string) => void;
  leagueId?: string;
};

export default function LeaguesFilter({
  onLeagueChange,
  leagueId = '',
}: LeaguesFilterProps) {
  const { data: leagues } = useFetchLeagues();

  const leagueOptions =
    leagues?.map((league) => ({
      key: league.id,
      value: league.name,
    })) ?? [];

  const form = useAppForm({
    defaultValues: {
      leagueId: leagueId,
    },
  });

  return (
    <form.AppField
      name="leagueId"
      listeners={{
        onChange: ({ value }) => {
          onLeagueChange(value);
        },
      }}
    >
      {(field) => (
        <field.Select
          label="Filter by League"
          options={leagueOptions}
          notSelectedOption={SELECT_OPTIONS.ALL_LEAGUES}
        />
      )}
    </form.AppField>
  );
}
