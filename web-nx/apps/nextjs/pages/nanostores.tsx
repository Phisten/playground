import { atom } from 'nanostores';
import { useStore } from '@nanostores/react';
import { Fragment } from 'react';

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
      <h1>NanoStore Demo</h1>
      <div className="grid gap-4 border rounded-sm">
        <label>new user name</label>
        <input className="px-4 py-2"></input>
      </div>

      <div className="grid gap-4">
        {testData.map((v) => (
          <Fragment key={v.id}>
            <div>
              <p>{v.name}</p>
            </div>
          </Fragment>
        ))}
      </div>
    </div>
  );
};

export default Page;
