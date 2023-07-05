import type { AspidaClient } from 'aspida'
import type { Methods as Methods0 } from '.'

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? '/api' : baseURL).replace(/\/$/, '')
  const PATH0 = '/auth/logout'
  const GET = 'GET'

  return {
    get: (option?: { config?: T | undefined } | undefined) =>
      fetch<void, Methods0['get']['resHeaders'], Methods0['get']['status']>(prefix, PATH0, GET, option).send(),
    $get: (option?: { config?: T | undefined } | undefined) =>
      fetch<void, Methods0['get']['resHeaders'], Methods0['get']['status']>(prefix, PATH0, GET, option).send().then(r => r.body),
    $path: () => `${prefix}${PATH0}`
  }
}

export type ApiInstance = ReturnType<typeof api>
export default api
