const name = 'csrf_token';
export const setCsrfToken = (token: string) => {
  if (token) {
    sessionStorage.setItem(name, JSON.stringify({ csrf_token: token }));
  }
};

export const loadCsrfToken = (): string => {
  const token = sessionStorage.getItem(name);
  if (token) {
    const tokenObject: { csrf_token: string } = JSON.parse(token);
    return tokenObject['csrf_token'];
  }
  return '';
};
