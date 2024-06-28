import { MouseEventHandler, RefObject, useEffect, useRef } from 'react';

export const useDragScroll = (scrollContainer: RefObject<HTMLDivElement>) => {
  const hasMoveRef = useRef(false);

  const childPreventDefaultOnClick: MouseEventHandler = (e) => {
    if (hasMoveRef.current) {
      e.preventDefault();
    }
  };

  useEffect(() => {
    const container: HTMLDivElement | null = scrollContainer.current;
    let startY: number;
    let startX: number;
    let scrollLeft: number;
    let scrollTop: number;
    let isDown: boolean;

    const mouseIsDown = (e: MouseEvent) => {
      hasMoveRef.current = false;
      if (container) {
        isDown = true;
        startY = e.pageY - container.offsetTop;
        startX = e.pageX - container.offsetLeft;
        scrollLeft = container.scrollLeft;
        scrollTop = container.scrollTop;
      }
    };
    const mouseUp = (e: MouseEvent) => {
      isDown = false;
    };
    const mouseMove = (e: MouseEvent) => {
      if (isDown) {
        e.preventDefault();
        if (container) {
          const y = e.pageY - container.offsetTop;
          const walkY = y - startY;
          container.scrollTop = scrollTop - walkY;
          const x = e.pageX - container.offsetLeft;
          const walkX = x - startX;
          container.scrollLeft = scrollLeft - walkX;

          const threshold = 6;
          if (walkX > threshold || walkX < -threshold) {
            hasMoveRef.current = true;
          }
        }
      }
    };

    if (container) {
      container.addEventListener('mousedown', (e) => mouseIsDown(e));
      container.addEventListener('mouseup', (e) => mouseUp(e));
      container.addEventListener('mouseleave', (e) => mouseUp(e));
      container.addEventListener('mousemove', (e) => mouseMove(e));
    }

    return () => {
      if (container) {
        container.removeEventListener('mousedown', (e) => mouseIsDown(e));
        container.removeEventListener('mouseup', (e) => mouseUp(e));
        container.removeEventListener('mouseleave', (e) => mouseUp(e));
        container.removeEventListener('mousemove', (e) => mouseMove(e));
      }
    };
  });

  return { childPreventDefaultOnClick };
};
