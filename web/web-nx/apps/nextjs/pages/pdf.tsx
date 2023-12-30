import React from 'react';
import { useIsClient } from 'usehooks-ts';
import { Pdf } from '../content/pdf';

export const PdfPage = () => {
  const isHydrated = useIsClient();
  return isHydrated ? (
    <>
      <Pdf />
    </>
  ) : null;
};

export default PdfPage;
