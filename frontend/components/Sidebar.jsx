'use client'

import useLogout from "@/hooks/useLogout"

export default function Sidebar({ children }) {
  const { logout } = useLogout()
  return (
    <>
      <button onClick={logout}>Logout</button>
      {children}
    </>
  )
}