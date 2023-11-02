import NewPostView from "@/components/NewPostView";
import PostView from "@/components/PostView";

export default async function Home() {
  const posts = await getPosts()

  return (
    <div className="grow overflow-scroll">
      <NewPostView />
      {posts.map((post) => <PostView post={post} />)}
    </div>
  )
}

async function getPosts() {
  const res = await fetch("http://127.0.0.1:8080/posts", { next: {revalidate: 0 }})
  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error('Failed to fetch data')
  }
  return res.json()
}