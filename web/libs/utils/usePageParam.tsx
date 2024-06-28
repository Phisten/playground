import { useRouter } from 'next/router';
import { getNextJsDynamicPathParamName } from '../nextDynamicPath';

/**
 * example :

export enum GenerateNextPath {
  CONTENT_NEWS_EDIT_NEWSBYID = '/content/news/edit/[NewsById]',
  CONTENT_NEWS_PREVIEW_NEWSBYID = '/content/news/preview/[NewsById]',
  ROOT = '/',
  STAFF = '/staff',
}

>> path: http://localhost/content/news/edit/12441 <<
const { param, paramName } = usePageParams(GenerateNextPath.CONTENT_NEWS_EDIT_NEWSBYID)
--> first render (hydrate): { param: undefined, paramName: 'NewsById' }
--> another render: { param: '12441', paramName: 'NewsById' }
--> onRoute to other page with same paramName:
    before route to http://localhost/content/news/preview/211
      { param: undefined, paramName: 'NewsById' }
    如果在 edit page 直接用 router.query.NewsById 來取得參數, 會在換頁到 preview 前 NewsById 被變更為 211
    因此在 router.query 頁面取得參數時需檢查 router.pathname 是否也正確, 否則在 paramName 相同時會導致換頁時取得意外的參數
 */

export const usePageParam = (nextJsPath: string) => {
  const router = useRouter();
  const lastPageParamName = getNextJsDynamicPathParamName(nextJsPath);
  const param =
    nextJsPath === router.pathname && lastPageParamName ? (router.query[lastPageParamName] as string) : undefined;

  return { param, router, paramName: lastPageParamName };
};
