import { MouseEventHandler, RefObject, useEffect, useRef } from "react";

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

    const getXY = (e: MouseEvent | TouchEvent) => {
      if (e instanceof MouseEvent) {
        return { eventX: e.pageX, eventY: e.pageY, touches: 1 };
      } else if (e instanceof TouchEvent) {
        return {
          eventX: e.touches[0].pageX,
          eventY: e.touches[0].pageY,
          touches: e.touches.length,
        };
      }
      return { eventX: -1, eventY: -1, touches: 0 };
    };

    const mouseIsDown = (e: MouseEvent | TouchEvent) => {
      const { eventX, eventY, touches } = getXY(e);
      hasMoveRef.current = false;
      if (container && touches === 1) {
        isDown = true;
        startY = eventY - container.offsetTop;
        startX = eventX - container.offsetLeft;
        scrollLeft = container.scrollLeft;
        scrollTop = container.scrollTop;
      }
    };
    const mouseUp = (e: MouseEvent | TouchEvent) => {
      isDown = false;
    };
    const mouseMove = (e: MouseEvent | TouchEvent) => {
      if (isDown) {
        e.preventDefault();
        const { eventX, eventY } = getXY(e);

        if (container) {
          const y = eventY - container.offsetTop;
          const walkY = y - startY;
          container.scrollTop = scrollTop - walkY;
          const x = eventX - container.offsetLeft;
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
      container.addEventListener("mousedown", (e) => mouseIsDown(e));
      container.addEventListener("mouseup", (e) => mouseUp(e));
      container.addEventListener("mouseleave", (e) => mouseUp(e));
      container.addEventListener("mousemove", (e) => mouseMove(e));
      container.addEventListener("touchstart", (e) => mouseIsDown(e));
      container.addEventListener("touchend", (e) => mouseUp(e));
      container.addEventListener("touchmove", (e) => mouseMove(e));
    }

    return () => {
      if (container) {
        container.removeEventListener("mousedown", (e) => mouseIsDown(e));
        container.removeEventListener("mouseup", (e) => mouseUp(e));
        container.removeEventListener("mouseleave", (e) => mouseUp(e));
        container.removeEventListener("mousemove", (e) => mouseMove(e));
        container.removeEventListener("touchstart", (e) => mouseIsDown(e));
        container.removeEventListener("touchend", (e) => mouseUp(e));
        container.removeEventListener("touchmove", (e) => mouseMove(e));
      }
    };
  });

  return { childPreventDefaultOnClick };
};
