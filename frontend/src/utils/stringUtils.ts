export const isStringEmpty = (str: string | null | undefined): boolean => {
  return str === null || str === undefined || str.trim().length === 0;
};

export const isUuidEmpty = (uuid: string | null | undefined): boolean => {
  return (
    uuid === null ||
    uuid === undefined ||
    uuid.length === 0 ||
    uuid === '00000000-0000-0000-0000-000000000000'
  );
};
