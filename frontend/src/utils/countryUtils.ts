export function getCountryIconUrl(code: string | null | undefined): string {
  if (!code) {
    return '';
  }
  const lowerCaseCode = code.toLowerCase();
  const countryIconUrl = `https://flagcdn.com/20x15/${lowerCaseCode}.png`;

  return countryIconUrl;
}

export function getCountryName(code: string | null | undefined): string {
  if (!code) {
    return 'N/A';
  }

  return countryCodeToName[code.toLowerCase()] || 'N/A';
}

export const countryCodeToName: Record<string, string> = {
  us: 'United States',
  de: 'Germany',
  fr: 'France',
  gb: 'United Kingdom',
  ca: 'Canada',
  lt: 'Lithuania',
};
