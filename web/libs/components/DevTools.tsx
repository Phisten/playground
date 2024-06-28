import { Button, Input, Message, Modal, ModalBody, ModalHeader } from '@mezzanine-ui/react';
import { Ref, useRef, useState } from 'react';

import { devConsoleLog } from './devConsoleLog';

declare const VERSION: string;
declare const DEV_MODE: boolean;
declare const DEV_CONSOLE_LOG: boolean;

const versionId = VERSION ?? '';

export const DevTools = () => {
  const [open, setOpen] = useState(false);
  const [hiddenButton, setHiddenButton] = useState(false);

  const testMessageInputRef: Ref<HTMLInputElement> = useRef(null);

  if (!DEV_CONSOLE_LOG) return null;

  return (
    <>
      <div className={'fixed z-[100] left-0 bottom-10 w-10 h-10 group/DebTools'}>
        <Button
          size="large"
          style={hiddenButton ? { backgroundColor: 'unset' } : { backgroundColor: 'var(--diag-black)' }}
          className={
            'transition-all ease-in-out duration-150 left-[calc(12px-160px)] group-hover/DebTools:left-2 ' +
            'absolute h-10 w-[160px] rounded-lg !text-dWhite opacity-75 text-center'
          }
          onClick={() => {
            setOpen(true);
          }}
        >
          {`${versionId ? `Build: ${versionId}` : 'DevTools'}`}
        </Button>
      </div>
      {DEV_MODE && (
        <Modal className="" onClose={() => setOpen(false)} open={open}>
          <ModalHeader>開發測試工具</ModalHeader>
          <ModalBody className="grid gap-4">
            {/* <div className='flex gap-4'>
            當前環境為：
            <a href={'http://localhost:4100'}><Typography color={DEV_MODE ? 'primary' : 'text-disabled'}>開發環境</Typography></a>
            <a href={DEV_URL}><Typography color={isDevSite ? 'primary' : 'text-disabled'}>測試站</Typography></a>
          </div> */}
            <div className="grid grid-flow-col gap-4">
              <Input inputRef={testMessageInputRef} defaultValue="測試用錯誤訊息" />
              <Button
                className="h-full"
                onClick={() => {
                  const msg = testMessageInputRef.current?.value ?? '';
                  devConsoleLog(msg);
                  Message?.['error'](msg);
                }}
              >{`Message`}</Button>
            </div>

            <Button onClick={() => setHiddenButton((x) => !x)}>{`開發工具按鈕改為 ${
              hiddenButton ? '可視' : '隱藏'
            }`}</Button>
          </ModalBody>
        </Modal>
      )}
    </>
  );
};

export default DevTools;
