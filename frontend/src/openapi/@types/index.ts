/* eslint-disable */
export type Response_CsrfToken = {
  csrf_token?: string | undefined
}

export type Response_EnabledSocialLogin = {
  socials?: string[] | undefined
}

export type Response_Error = {
  description?: string | undefined
  field?: string | undefined
  help?: string | undefined
  message?: string | undefined
}

export type Response_Errors = {
  errors?: Response_Error[] | undefined
}

export type Response_User = {
  email?: string | undefined
  email_verified?: boolean | undefined
  picture?: string | undefined
  profile?: string | undefined
  sub?: string | undefined
  username?: string | undefined
}

export type Response_WebUIManifest = {
  header_name?: string | undefined
  token_base?: boolean | undefined
}
