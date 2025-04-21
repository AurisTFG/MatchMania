import dayjs from 'dayjs';

export function getStartOfDay(days = 0): Date {
  return dayjs()
    .startOf('day')
    .add(dayjs().utcOffset(), 'minute')
    .add(days, 'day')
    .toDate();
}
