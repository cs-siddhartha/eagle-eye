import { createFileRoute, Link } from '@tanstack/react-router'
import {
  Activity,
  ArrowRight,
  BellRing,
  ChartNoAxesCombined,
  Check,
  CircleDot,
  Gauge,
  Search,
  ShieldCheck,
  Sparkles,
} from 'lucide-react'
import { ModeToggle } from '~/components/mode-toggle'
import { Button } from '~/components/ui/button'

export const Route = createFileRoute('/')({
  component: Home,
})

// Gives first-time visitors a focused overview of the observability product
// and directs them toward the existing dashboard routes.
function Home() {
  return (
    <>

      <main className="min-h-screen bg-background text-foreground">
        <header className="border-y bg-background/95">
          <div className="mx-auto flex max-w-7xl items-center justify-between px-6 py-4 lg:px-8">
            <Link to="/" className="flex items-center gap-3">
              <span className="flex size-9 items-center justify-center bg-primary text-primary-foreground">
                <Activity className="size-5" />
              </span>
              <span className="text-sm font-semibold tracking-[0.2em] uppercase">
                Signal
              </span>
            </Link>

            <nav className="hidden items-center gap-8 text-xs font-semibold tracking-widest uppercase md:flex">
              <a href="#platform" className="text-muted-foreground hover:text-foreground">
                Platform
              </a>
              <a href="#workflow" className="text-muted-foreground hover:text-foreground">
                Workflow
              </a>
              <Link to="/posts" className="text-muted-foreground hover:text-foreground">
                Changelog
              </Link>
            </nav>

            <div className="flex items-center gap-2">
              <ModeToggle />
              <Button asChild className="hidden sm:inline-flex">
                <Link to="/users">
                  Open dashboard
                  <ArrowRight data-icon="inline-end" />
                </Link>
              </Button>
            </div>
          </div>
        </header>

        <section className="relative overflow-hidden border-b">
          <div className="absolute inset-0 -z-10 bg-[linear-gradient(to_right,var(--border)_1px,transparent_1px),linear-gradient(to_bottom,var(--border)_1px,transparent_1px)] bg-[size:64px_64px] opacity-40" />
          <div className="mx-auto grid max-w-7xl gap-16 px-6 py-24 lg:grid-cols-[1.1fr_0.9fr] lg:px-8 lg:py-32">
            <div className="flex flex-col justify-center">
              <div className="mb-8 flex w-fit items-center gap-2 border bg-background px-3 py-2 text-xs font-semibold tracking-widest uppercase">
                <Sparkles className="size-3.5" />
                Observability without the noise
              </div>

              <h1 className="max-w-4xl text-5xl font-semibold tracking-[-0.05em] sm:text-6xl lg:text-7xl">
                See the signal.
                <span className="block text-muted-foreground">
                  Fix what matters.
                </span>
              </h1>

              <p className="mt-8 max-w-2xl text-lg leading-8 text-muted-foreground">
                Unify logs, metrics, and traces in one focused workspace.
                Signal helps engineering teams detect issues faster, understand
                impact, and move from alert to resolution with confidence.
              </p>

              <div className="mt-10 flex flex-col gap-3 sm:flex-row">
                <Button asChild size="lg">
                  <Link to="/users">
                    Start monitoring
                    <ArrowRight data-icon="inline-end" />
                  </Link>
                </Button>
                <Button asChild variant="outline" size="lg">
                  <a href="#platform">Explore platform</a>
                </Button>
              </div>

              <div className="mt-10 flex flex-wrap gap-x-8 gap-y-3 text-sm text-muted-foreground">
                <span className="flex items-center gap-2">
                  <Check className="size-4 text-foreground" />
                  Five-minute setup
                </span>
                <span className="flex items-center gap-2">
                  <Check className="size-4 text-foreground" />
                  OpenTelemetry native
                </span>
                <span className="flex items-center gap-2">
                  <Check className="size-4 text-foreground" />
                  No credit card
                </span>
              </div>
            </div>

            <div className="relative border bg-card p-3 shadow-2xl shadow-foreground/5">
              <div className="border bg-background">
                <div className="flex items-center justify-between border-b px-4 py-3">
                  <div className="flex items-center gap-2">
                    <CircleDot className="size-4 text-emerald-500" />
                    <span className="text-xs font-semibold tracking-widest uppercase">
                      Production
                    </span>
                  </div>
                  <span className="text-xs text-muted-foreground">Last 30 min</span>
                </div>

                <div className="grid grid-cols-3 border-b">
                  <div className="border-r p-4">
                    <p className="text-xs text-muted-foreground">Requests</p>
                    <p className="mt-2 text-2xl font-semibold">2.4M</p>
                    <p className="mt-1 text-xs text-emerald-600">+12.8%</p>
                  </div>
                  <div className="border-r p-4">
                    <p className="text-xs text-muted-foreground">P95 latency</p>
                    <p className="mt-2 text-2xl font-semibold">184ms</p>
                    <p className="mt-1 text-xs text-emerald-600">-8.2%</p>
                  </div>
                  <div className="p-4">
                    <p className="text-xs text-muted-foreground">Error rate</p>
                    <p className="mt-2 text-2xl font-semibold">0.08%</p>
                    <p className="mt-1 text-xs text-muted-foreground">Stable</p>
                  </div>
                </div>

                <div className="p-5">
                  <div className="mb-6 flex items-center justify-between">
                    <div>
                      <p className="text-sm font-semibold">Request volume</p>
                      <p className="text-xs text-muted-foreground">
                        All services
                      </p>
                    </div>
                    <ChartNoAxesCombined className="size-5 text-muted-foreground" />
                  </div>

                  <div className="flex h-52 items-end gap-1.5">
                    {/* Builds a lightweight dashboard preview without adding a chart dependency. */}
                    {[38, 52, 44, 68, 58, 76, 63, 82, 74, 91, 78, 88, 72, 95, 86, 100, 82, 92, 76, 88].map(
                      (height, index) => (
                        <div
                          key={`${height}-${index}`}
                          className="flex-1 bg-primary/80"
                          style={{ height: `${height}%` }}
                        />
                      ),
                    )}
                  </div>
                </div>

                <div className="border-t p-4">
                  <div className="flex items-center gap-3 border bg-muted/40 p-3">
                    <BellRing className="size-4" />
                    <div className="min-w-0 flex-1">
                      <p className="truncate text-xs font-semibold">
                        Checkout latency recovered
                      </p>
                      <p className="text-xs text-muted-foreground">
                        Resolved automatically · 2m ago
                      </p>
                    </div>
                    <span className="text-xs font-semibold text-emerald-600">
                      Healthy
                    </span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>

        <section id="platform" className="border-b">
          <div className="mx-auto max-w-7xl px-6 py-24 lg:px-8">
            <div className="max-w-2xl">
              <p className="text-xs font-semibold tracking-[0.2em] text-muted-foreground uppercase">
                One operating picture
              </p>
              <h2 className="mt-4 text-3xl font-semibold tracking-tight sm:text-4xl">
                Everything you need to understand production.
              </h2>
            </div>

            <div className="mt-14 grid border md:grid-cols-3">
              <article className="border-b p-8 md:border-r md:border-b-0">
                <Search className="size-6" />
                <h3 className="mt-8 text-lg font-semibold">Explore in context</h3>
                <p className="mt-3 leading-7 text-muted-foreground">
                  Move between traces, logs, and metrics without losing the
                  thread of an investigation.
                </p>
              </article>
              <article className="border-b p-8 md:border-r md:border-b-0">
                <Gauge className="size-6" />
                <h3 className="mt-8 text-lg font-semibold">Measure what matters</h3>
                <p className="mt-3 leading-7 text-muted-foreground">
                  Track service health against the customer outcomes and SLOs
                  your team owns.
                </p>
              </article>
              <article className="p-8">
                <ShieldCheck className="size-6" />
                <h3 className="mt-8 text-lg font-semibold">Respond with clarity</h3>
                <p className="mt-3 leading-7 text-muted-foreground">
                  Route actionable alerts with enough context to resolve an
                  incident before impact spreads.
                </p>
              </article>
            </div>
          </div>
        </section>

        <section id="workflow">
          <div className="mx-auto max-w-7xl px-6 py-24 lg:px-8">
            <div className="grid items-end gap-10 border bg-primary px-8 py-12 text-primary-foreground md:grid-cols-[1fr_auto] md:px-12">
              <div>
                <p className="text-xs font-semibold tracking-[0.2em] uppercase opacity-70">
                  Ready when you are
                </p>
                <h2 className="mt-4 max-w-2xl text-3xl font-semibold tracking-tight sm:text-4xl">
                  Make production easier to understand.
                </h2>
                <p className="mt-4 max-w-xl leading-7 opacity-70">
                  Connect your first service and start seeing useful telemetry
                  in minutes.
                </p>
              </div>
              <Button asChild size="lg" variant="secondary">
                <Link to="/users">
                  Get started
                  <ArrowRight data-icon="inline-end" />
                </Link>
              </Button>
            </div>
          </div>
        </section>
      </main>
    </>
  )
}
