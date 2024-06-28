import { MutableRefObject, useEffect, useState } from 'react';
import { useIsomorphicLayoutEffect } from 'usehooks-ts';
function useScrollLock(initialLocked = false, lockedElement?: MutableRefObject<HTMLDivElement>) {
  const [locked, setLocked] = useState(initialLocked);
  useIsomorphicLayoutEffect(() => {
    if (!locked) {
      return;
    }
    const element = lockedElement?.current ?? document.body;
    const originalOverflow = element.style.overflow;
    const originalPaddingRight = element.style.paddingRight;
    element.style.overflow = 'hidden';
    const root = document.getElementById('___gatsby');
    const scrollBarWidth = root ? root.offsetWidth - root.scrollWidth : 0;
    if (scrollBarWidth) {
      element.style.paddingRight = `${scrollBarWidth}px`;
    }
    return () => {
      element.style.overflow = originalOverflow;
      if (scrollBarWidth) {
        element.style.paddingRight = originalPaddingRight;
      }
    };
  }, [locked]);
  useEffect(() => {
    if (locked !== initialLocked) {
      setLocked(initialLocked);
    }
  }, [initialLocked]);
  return [locked, setLocked];
}
export default useScrollLock;
