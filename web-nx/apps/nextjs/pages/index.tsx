import Link from 'next/link';

export function Index() {
  return (
    <>
      <div className="grid">
        <div className="mx-auto">
          <Link href={'/NxDefault'}>
            <p className=" text-blue-500  text-5xl">NxDefault Page</p>
          </Link>
        </div>
      </div>
    </>
  );
}

export default Index;
