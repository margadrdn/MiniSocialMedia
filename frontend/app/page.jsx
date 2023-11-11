import Sidebar from "@/components/Sidebar";
import Posts from "@/components/post/Posts";

export default function Root() {
  return (
    <>
      <Sidebar>
        <Posts />
      </Sidebar>
    </>
  )
}