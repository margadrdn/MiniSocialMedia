export default function useCookie() {
  const setCookie = (name, value, options) => {
    document.cookie = `${name}=${value}; ${options}`;
  };

  return { setCookie }
}