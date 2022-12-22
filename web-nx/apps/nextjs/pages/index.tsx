import { Block } from '@web-nx/ui';
import Link from 'next/link';
import { NextPath } from '../utils/path';

export function Index() {
  return (
    <>
      <div className="grid bg-slate-300 w-full">
        <div className="mx-auto">
          <Block title="">
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
