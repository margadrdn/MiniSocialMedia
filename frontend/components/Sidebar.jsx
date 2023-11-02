import Link from "next/link";

export default function Sidebar() {
  return (
    <div className="">
      <Link className="block bg-green-200 rounded-full p-2 m-2" href="/">
        Home
      </Link>
      <Link className="block bg-green-200 rounded-full p-2 m-2" href="/messages">
        Messages
      </Link>
    </div>
  )
}