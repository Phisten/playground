import { Block } from '@web-nx/ui';
import Link from 'next/link';
import { NextPath } from '../utils/path';
import dayjs from 'dayjs';

export function Index() {
  const timeString = '2023-05-25 12:34:56 +03:00';
  const parsedTime = dayjs(timeString, 'YYYY-MM-DD HH:mm:ss ZZ');

  return (
    <>
      <div className="grid bg-slate-300 w-full">
        <div className="mx-auto">
          <Block title="">
            {parsedTime.format('YYYY-MM-DD HH:mm:ss ZZ')}
            <Link href={NextPath.NXDEFAULT}>
              <p className="text-blue-500 text-5xl text-center">
                NxDefault Page
              </p>
            </Link>
            <Link href={NextPath.MOCK_IMAGES}>
              <p>mock-images Page</p>
            </Link>
            <Link href={NextPath.NANOSTORES}>
              <p>NanoStores React</p>
            </Link>
            <Link href={NextPath.STYLE}>
              <p>Styles</p>
            </Link>
          </Block>
          <Block title="2313">123133</Block>
        </div>
      </div>
    </>
  );
}

export default Index;
