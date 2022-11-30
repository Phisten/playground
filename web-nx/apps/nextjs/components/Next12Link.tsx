import Link, { LinkProps } from 'next/link';
import { ReactNode } from 'react';

export const Next12Link = (props: LinkProps & { children?: ReactNode }) => {
  return (
    <Link {...props} passHref>
      <a href="passHref">{props?.children}</a>
    </Link>
  );
};
