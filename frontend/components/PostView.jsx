export default function PostView({ post }) {
  return (
    <div className="w-1/2 bg-indigo-100 p-4 mx-auto m-4">
      <div>
        {post.id}
      </div>
      <div>
        {post.author}
      </div>
      <div>
        {post.body}
      </div>
    </div>
  )
}