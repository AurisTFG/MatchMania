import dayjs from 'dayjs';

export function getStartOfDay(offsetDays = 0): Date {
  return dayjs()
    .startOf('day')
    .add(dayjs().utcOffset(), 'minute')
    .add(offsetDays, 'day')
    .toDate();
}

export function getToday(): Date {
  return dayjs().toDate();
}
