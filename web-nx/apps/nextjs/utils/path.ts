import { GenerateNextPath as NextPath } from './generateNextPath';
export { GenerateNextPath as NextPath } from './generateNextPath';

// use "yarn client.pages"
// generate-next-links: https://github.com/lindeneg/generate-next-links

/** next/link with DynamicPath
    <Link
      href={{
        pathname: NextPath.NEWS_DETAILBYID, // = '/news/[detailById]'
        query: {
          detailById: props.id
        },
      }}
    >
 */
export function nextDynamicPath(nextPath: NextPath, param: string) {
  return nextPath.replace(/\[\S+\]/, param);
}
