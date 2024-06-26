import { atom } from 'nanostores';
import { useStore } from '@nanostores/react';
import { FormEventHandler, Fragment } from 'react';
import { Block, BaseField, BlockConfigProvider } from '@web-nx/ui';
import { Button } from '@mui/material';
import { useIsClient } from 'usehooks-ts';

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

const DataList = (props: { testData: User[] }) => {
  const { testData } = props;

  return (
    <Block
      className="grid gap-4 border rounded-md p-2"
      title={'-- User List --'}
    >
      {testData.map((v) => (
        <Fragment key={v.id}>
          <Block title={v.id}>
            <div className="grid items-center grid-cols-[1fr_100px_100px]">
              <p className="text-base">{v.name}</p>
              <Button>edit</Button>
              <Button>remove</Button>
            </div>
          </Block>
        </Fragment>
      ))}
    </Block>
  );
};

const PageContent = (props: {
  testData: User[];
  handleSubmit: FormEventHandler<HTMLFormElement>;
}) => {
  const { testData, handleSubmit } = props;
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

      <DataList testData={testData} />
    </div>
  );
};

const Page = () => {
  const testData = useStore(users);
  const isHydrated = useIsClient();

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
    <BlockConfigProvider
      config={{
        default: { className: 'shadow-md text-sm' },
      }}
    >
      <PageContent testData={testData} handleSubmit={handleSubmit} />
    </BlockConfigProvider>
  );
};

export default Page;
