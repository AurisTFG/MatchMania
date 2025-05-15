import {
  Box,
  CircularProgress,
  Container,
  Paper,
  Typography,
} from '@mui/material';
import { useFetchMatches } from 'api/hooks/matchHooks';
import { useFetchQueues } from 'api/hooks/queueHooks';
import UserAvatar from 'components/UserAvatar';
import QueueForm from 'components/queues/QueueForm';
import withAuth from 'hocs/withAuth';
import { useAuth } from 'providers/AuthProvider';
import { Permission } from 'types/enums/permission';

function MatchmakingPage() {
  const { user } = useAuth();
  const { data: queues, isLoading: isQueuesLoading } = useFetchQueues();
  const { data: matches, isLoading: isMatchesLoading } = useFetchMatches();

  const filteredQueues =
    queues?.filter((queue) => queue.teams.length > 0) ?? [];
  const filteredMatches =
    matches?.filter((match) => match.teams.length > 0) ?? [];

  const isUserInQueue = queues?.some((queue) =>
    queue.teams.some((team) =>
      team.players.some((player) => player.id === user?.id),
    ),
  );

  const isUserInMatch = matches?.some((match) =>
    match.teams.some((team) =>
      team.players.some((player) => player.id === user?.id),
    ),
  );

  const selectedTeamId = isUserInMatch
    ? matches
        ?.find((match) =>
          match.teams.some((team) =>
            team.players.some((player) => player.id === user?.id),
          ),
        )
        ?.teams.find((team) =>
          team.players.some((player) => player.id === user?.id),
        )?.id
    : isUserInQueue
      ? queues
          ?.find((queue) =>
            queue.teams.some((team) =>
              team.players.some((player) => player.id === user?.id),
            ),
          )
          ?.teams.find((team) =>
            team.players.some((player) => player.id === user?.id),
          )?.id
      : undefined;

  const selectedLeagueId = isUserInMatch
    ? matches?.find((match) =>
        match.teams.some((team) =>
          team.players.some((player) => player.id === user?.id),
        ),
      )?.league.id
    : isUserInQueue
      ? queues?.find((queue) =>
          queue.teams.some((team) =>
            team.players.some((player) => player.id === user?.id),
          ),
        )?.league.id
      : undefined;

  return (
    <Container maxWidth="md">
      <Box sx={{ mt: 4 }}>
        <QueueForm
          queueState={
            isUserInMatch ? 'in-match' : isUserInQueue ? 'queued' : 'idle'
          }
          selectedTeamId={selectedTeamId}
          selectedLeagueId={selectedLeagueId}
        />

        <Paper
          elevation={2}
          sx={{
            mt: 4,
            p: 3,
            borderRadius: 2,
            backgroundColor: 'background.paper',
            boxShadow: 2,
          }}
        >
          <Typography
            variant="h6"
            fontWeight={600}
          >
            Queues
          </Typography>
          {isQueuesLoading ? (
            <CircularProgress sx={{ mt: 2 }} />
          ) : (
            <>
              {filteredQueues.length === 0 ? (
                <Typography
                  sx={{ mt: 2 }}
                  color="text.secondary"
                >
                  No active queues found.
                </Typography>
              ) : (
                filteredQueues.map((queue) => (
                  <Box
                    key={queue.id}
                    sx={{ mt: 2, borderBottom: '1px solid #ccc', pb: 2 }}
                  >
                    <Typography
                      variant="subtitle1"
                      fontWeight={500}
                    >
                      Game Mode: {queue.gameMode}
                    </Typography>
                    <Box
                      sx={{
                        display: 'flex',
                        alignItems: 'center',
                        gap: 1,
                        mt: 1,
                      }}
                    >
                      <UserAvatar
                        imageUrl={queue.league.logoUrl}
                        name={queue.league.name}
                        size={24}
                      />
                      <Typography>{queue.league.name}</Typography>
                    </Box>
                    <Box sx={{ mt: 1 }}>
                      <Typography fontWeight={500}>Teams:</Typography>
                      <Box sx={{ pl: 2 }}>
                        {queue.teams.map((team) => (
                          <Box
                            key={team.id}
                            sx={{
                              display: 'flex',
                              alignItems: 'center',
                              gap: 1,
                              mt: 0.5,
                            }}
                          >
                            <UserAvatar
                              imageUrl={team.logoUrl}
                              name={team.name}
                              size={24}
                            />
                            <Typography>{team.name}</Typography>
                          </Box>
                        ))}
                      </Box>
                    </Box>
                  </Box>
                ))
              )}
            </>
          )}
        </Paper>

        <Paper
          elevation={2}
          sx={{
            mt: 4,
            p: 3,
            borderRadius: 2,
            backgroundColor: 'background.paper',
            boxShadow: 2,
          }}
        >
          <Typography
            variant="h6"
            fontWeight={600}
          >
            Matches
          </Typography>
          {isMatchesLoading ? (
            <CircularProgress sx={{ mt: 2 }} />
          ) : filteredMatches.length === 0 ? (
            <Typography
              sx={{ mt: 2 }}
              color="text.secondary"
            >
              No ongoing matches found.
            </Typography>
          ) : (
            filteredMatches.map((match, i) => (
              <Box
                key={i}
                sx={{ mt: 2, borderBottom: '1px solid #ccc', pb: 2 }}
              >
                <Typography
                  variant="subtitle1"
                  fontWeight={500}
                >
                  Game Mode: {match.gameMode}
                </Typography>
                <Box
                  sx={{ display: 'flex', alignItems: 'center', gap: 1, mt: 1 }}
                >
                  <UserAvatar
                    imageUrl={match.league.logoUrl}
                    name={match.league.name}
                    size={24}
                  />
                  <Typography>{match.league.name}</Typography>
                </Box>
                <Box sx={{ mt: 1 }}>
                  <Typography fontWeight={500}>Teams:</Typography>
                  <Box sx={{ pl: 2 }}>
                    {match.teams.map((team, idx) => (
                      <Box
                        key={idx}
                        sx={{
                          display: 'flex',
                          alignItems: 'center',
                          gap: 1,
                          mt: 0.5,
                        }}
                      >
                        <UserAvatar
                          imageUrl={team.logoUrl}
                          name={team.name}
                          size={24}
                        />
                        <Typography>{team.name}</Typography>
                      </Box>
                    ))}
                  </Box>
                </Box>
              </Box>
            ))
          )}
        </Paper>
      </Box>
    </Container>
  );
}

export default withAuth(MatchmakingPage, {
  permission: Permission.ViewQueue,
});
