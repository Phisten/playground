/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState, useEffect, useCallback } from 'react';

export const usePagination = (offset: number, initDeps: any[] = []) => {
  const [currentPage, setCurrentPage] = useState<number>(0);

  const sliceData = useCallback(
    (data: any[]) => {
      const maxSize = data.length;

      const start = currentPage * offset;
      const end = Math.min(maxSize, start + offset);

      return data.slice(start, end);
    },
    [offset, currentPage],
  );

  useEffect(() => {
    setCurrentPage(0);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, initDeps);

  return {
    currentPage: currentPage + 1,
    setCurrentPage,
    sliceData,
  };
};
