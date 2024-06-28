import Link, { LinkProps } from 'next/link';
import { ReactNode } from 'react';

export const NextLink = (props: LinkProps & { children?: ReactNode }) => {
  const emptyHref = !props?.href;
  return (
    <Link {...props} passHref>
      <a
        className={emptyHref ? 'cursor-default' : undefined}
        href="passHref"
        onClick={(e) => {
          if (emptyHref) {
            e.preventDefault();
          }
        }}
      >
        {props?.children}
      </a>
    </Link>
  );
};
