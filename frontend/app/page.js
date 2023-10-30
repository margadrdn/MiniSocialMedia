import PostView from "@/components/PostView";

export default async function Home() {
  const posts = await getPosts()

  return (
    <div className="bg-indigo-400">
      {posts.map((post) => <PostView post={post} />)}
    </div>
  )
}

async function getPosts() {
  const res = await fetch("http://127.0.0.1:8080/posts", { cache: 'no-store' })
  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error('Failed to fetch data')
  }
  return res.json()
}