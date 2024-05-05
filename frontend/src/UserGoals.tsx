import { LoaderFunction, useLoaderData } from 'react-router-dom';
import { Nav } from './components/Nav';
import { LoaderData } from './utils/loader-data';
import { Heading } from './components/Heading';
import { getGoals } from './utils/goal';

export const loader = (async () => {
  return await getGoals();
}) satisfies LoaderFunction;

export const UserGoals = () => {
  const data = useLoaderData() as LoaderData<typeof loader>;
  return (
    <div className="min-h-screen">
      <Nav />

      <main className="px-10 pt-10">
        {data.goals && data.goals.length ? (
          data.goals.map((g) => {
            // TODO: handle real goals
            return (
              <section className="flex justify-between align-middle mb-6">
                <Heading el="h2" type="h3" className="self-center">
                  {g.name}
                </Heading>
                <section className="text-right">
                  {/**
                   * TODO: Inject correct goal durations
                   */}
                  <p>Today's goal: 1hr</p>
                  {/**
                   * TODO: Inject correct goal durations
                   */}
                  <p>Time tracked: 20min</p>
                </section>
              </section>
            );
          })
        ) : (
          <section>No goals</section>
        )}
      </main>
    </div>
  );
};
