import { NextResponse } from "next/server"

export default function middleware(request) {
  const username = request.cookies.get('username')?.value
  const token = request.cookies.get('usertoken')?.value

  let auth_token = null

  if (username && token) {
    auth_token = { username, token }
  }

  if (request.nextUrl.pathname.startsWith('/auth') && !auth_token) {
    return
  }

  if (!auth_token) {
    return NextResponse.redirect(new URL('/auth', request.url))
  }

  if (request.url.includes('/auth') && auth_token) {
    return NextResponse.redirect(new URL('/', request.url))
  }

  if (auth_token) {
    return
  }
}

export const config = {
  matcher: ['/', '/auth']
}