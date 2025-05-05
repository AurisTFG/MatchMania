import { SELECT_OPTIONS } from 'constants/selectOptions';

export function isSelectOptionValid(
  option: string | null | undefined,
): boolean {
  return !!option && option !== SELECT_OPTIONS.NOT_SELECTED.key;
}
