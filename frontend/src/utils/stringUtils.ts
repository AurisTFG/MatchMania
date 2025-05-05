export const isStringEmpty = (str?: string | null): boolean => !str?.trim();

export const isUuidEmpty = (uuid?: string | null): boolean =>
  isStringEmpty(uuid) || uuid === '00000000-0000-0000-0000-000000000000';
