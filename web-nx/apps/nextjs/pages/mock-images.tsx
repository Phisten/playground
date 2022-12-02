import Image from 'next/image';
import { Fragment } from 'react';
import useSWR from 'swr';

/* eslint-disable-next-line */
export interface MockImagesProps {}

export function MockImages(props: MockImagesProps) {


  const { data } = useSWR('https://picsum.photos/v2/list')

  return (
    <div>
      <h1 className="text-2xl">Welcome to MockImages!</h1>
      {[1, 2, 3].map((x) => {
        return (
          <Fragment key={`${x}`}>
            <div className="">
              <Image
                src={'https://picsum.photos/400/200'}
                unoptimized
                alt=""
                // fill=''
                width={400}
                height={200}
              />
            </div>
          </Fragment>
        );
      })}
    </div>
  );
}

export default MockImages;
