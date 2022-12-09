import { atom } from 'nanostores';
import { useStore } from '@nanostores/react';
import { Fragment, PropsWithChildren, ReactNode } from 'react';
import { Block } from '@web-nx/ui';
import BaseField from 'libs/ui/src/lib/BaseField';

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

export const Page = () => {
  const testData = useStore(users);

  return (
    <div className="grid my-20 mx-auto text-center gap-10 max-w-[400px]">
      <Block
        className="grid gap-4 border rounded-md p-2"
        title={'NanoStore Demo'}
      >
        <form
          className="grid"
          onSubmit={(e) => {
            e.preventDefault();
            console.log('submit');
          }}
        >
          <BaseField label={'new user name'}>
            <input className="px-4 py-2"></input>
          </BaseField>
          <button
            type="button"
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
