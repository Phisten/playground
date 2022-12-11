import { atom } from 'nanostores';
import { useStore } from '@nanostores/react';
import {
  FormEventHandler,
  Fragment,
  PropsWithChildren,
  ReactNode,
  useState,
} from 'react';
import { Block } from '@web-nx/ui';
import BaseField from 'libs/ui/src/lib/BaseField';
// import { randomUUID } from 'crypto';

type User = {
  id: string;
  name: string;
};
const users = atom<User[]>([
  { id: '1', name: 'userA' },
  { id: '2', name: 'userB' },
]);

const addUser = (newUser: User) => {
  return users.set([...users.get(), newUser]);
};
type FormData = {
  userId?: string;
};

export const Page = () => {
  const testData = useStore(users);

  const handleSubmit: FormEventHandler<HTMLFormElement> = (event) => {
    const formData = new FormData(event.currentTarget);
    event.preventDefault();

    console.log({ formData });
    formData.forEach((v: FormDataEntryValue, key, parent) => {
      console.log({ v, key, parent });
    });

    addUser({
      id: global.crypto.randomUUID().toString(),
      name: formData.get('userName').toString() ?? 'unknownUser',
    });
  };

  return (
    <div className="grid my-20 mx-auto text-center gap-10 max-w-[400px]">
      <Block
        className="grid gap-4 border rounded-md p-2"
        title={'NanoStore Demo'}
      >
        <form className="grid" onSubmit={handleSubmit}>
          <BaseField label={'new user name'}>
            <input
              className="px-4 py-2"
              name="userName"
              placeholder="please input user name"
            ></input>
          </BaseField>
          <button
            type="submit"
            onClick={() => {
              console.log('first');
            }}
          >
            btn
          </button>
        </form>
      </Block>

      <Block
        className="grid gap-4 border rounded-md p-2"
        title={'-- User List --'}
      >
        {testData.map((v) => (
          <Fragment key={v.id}>
            <Block className="text-sm" title={v.id}>
              <p className="text-base">{v.name}</p>
            </Block>
          </Fragment>
        ))}
      </Block>
    </div>
  );
};

export default Page;
