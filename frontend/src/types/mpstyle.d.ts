export {};

declare global {
  const MPStyle: {
    Parser: {
      toHTML: (input: string) => string;
    };
  };
}
