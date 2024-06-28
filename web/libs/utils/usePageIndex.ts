import { useQueryState } from 'next-usequerystate';
import { useCallback } from 'react';

// const { pageIndex, setPageIndex, limit, offset } = usePageIndex(10);
/** index start at 0, page start at 1 */
export const usePageIndex = (pageLength: number) => {
  const [pageStr, setPageStr] = useQueryState('page', { defaultValue: '1' });
  const page = Number(pageStr);
  const pageIndex = page - 1;

  const setPage = useCallback(
    (newPage: number) => {
      setPageStr(String(newPage));
      const newPageIndex = page - 1;
      return { newPage, newPageIndex, offset: newPageIndex * pageLength };
    },
    [page, pageLength, setPageStr],
  );

  const setPageIndex = useCallback(
    (newPageIndex: number) => {
      setPage(newPageIndex + 1);
      return { newPage: newPageIndex - 1, newPageIndex, offset: newPageIndex * pageLength };
    },
    [pageLength, setPage],
  );

  return {
    pageIndex,
    setPageIndex,
    page,
    setPage,
    limit: pageLength,
    offset: pageIndex * pageLength,
  };
};
