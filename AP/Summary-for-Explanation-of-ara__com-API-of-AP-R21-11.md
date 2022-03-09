# Reading of the document

## 6 API Element

Take Radar Service as the example. It is easier to understand.

### Instance Identifiers

Used on client/proxy side

a technical binding specific identifier (technical)

```
ara::com::InstanceIdentifier
```

name is constructed from AUTOSAR meta-model.

The C++ representation of such an "instance specifier" is the class ara::**core**::Instance**Specifier** (local name in the software developers realm)

```
ara::core::InstanceSpecifier
```

The API ara:com provides the translation between these two.

```cpp
namespace ara {
namespace com {
namespace runtime {
ara::com::InstanceIdentifierContainer ara::com::runtime::ResolveInstanceIDs(ara::core::InstanceSpecifier modelName);
}
} // namespace com
} // namespace ara
```

The above function convert from local name to technical one.

`ara::com::InstanceIdentifierContainer` represents a collection of `ara::com::InstanceIdentifier`, which means, meanings one local name one (for software component developer) can mapping to multiple technical bindings.

one local to multiple bindings for **skeleton/server** is common, the client can use their preferred binding.

one local to multiple bindings for **proxy/client** is rare, e.g., to support some fail-over approaches. (binding A not work then we can count on binding b)

`Service Instance Manifest`, `ara::core::InstanceSpecifier` must be unambiguous within the bundled `service instance manifest`.

### Proxy Class

Proxy class is generated from the **service interface description** of the AUTOSAR meta model.

Then we can see from the generated “abstract class” or a “**mock** class” against which the application developer could implemement his service consumer application.

Then see from the example code, we know or class `RadarServiceProxy`:

Pay attention to `static` , the generated code, they are the framework help functions (static function for whole class) for Service Proxy in client side.

*   HandleType
    *   hidden for app dev
    *   platform vendor is responsible
*   StartFindService
    *   handler, gets called any time the service availability of the matching services changes.
    *   instanceId, type is ara::com::InstanceIdentifier
    *   return FindServiceHandle, which shall be used to stop the availability monitoring and related firing of the given handler.
*   StartFindServe (with different signature)

```cpp
static ara::com::FindServiceHandle StartFindService(
ara::com::FindServiceHandler<RadarServiceProxy::HandleType> handler,
ara::com::InstanceIdentifier instanceId);

static ara::com::FindServiceHandle StartFindService(
ara::com::FindServiceHandler<RadarServiceProxy::HandleType> handler,
ara::core::InstanceSpecifier instanceSpec);

/**
* This is an overload of the StartFindService method using neither
* instance specifier nor instance identifier.
* Semantics is, that ALL instances of the service shall be found, by
* using all available/configured technical bindings.
*
*/
static ara::com::FindServiceHandle StartFindService(
ara::com::FindServiceHandler<RadarServiceProxy::HandleType> handler);
```

*   `static void StopFindService(ara::com::FindServiceHandle handle);`

*   `static ara::com::ServiceHandleContainer<RadarServiceProxy::HandleType>`  
    `FindService(ara::com::InstanceIdentifier instanceId);`

    *   synchronous

    *   does reflect the availability at the time of the method call.  
        No further (background) checks of availability are done.

*   `FindService` (overloads)

    *   ara::core::InstanceSpecifier

    *   void , ALL instance of the service shall be found

*   `explicit RadarServiceProxy(HandleType &handle);`

    *   Here we can see the HandleType (hidden from app dev? Implementation side, contains InstanceIdendifier, for Client to specify the service)

    *   HandleType is used to identify a service

*   Other class instance related operation config (c++ wise)

    *   not copy constructible

    *   not copy assignable

*   Public Members

    *   Events

    *   Fields

    *   methods


#### Constructor and Handle Concept

`ctor`

Signature:
`explicit RadarServiceProxy(HandleType &handle);`

After the call to the `ctor` you have a proxy instance for communicating with
the service. The `handle` contains the addressing information for the service,
which **Communication Management** binding implementation use to set the
contact.

What exactly this address information in `HandleType` contains is totally dependent on the
binding implementation/technical transport layer!

For an application developer, he only need to pay attention to his application
implementation, and be independent of Communication Management.


For AUTOSAR Binding Implementation detail. We skip for now.

#### Finding Services
| function | life cycle | static |
| --- | --- | --- |
| `StartFindService` | notifies the caller via a given callback(handler) <br> any time the availability changes | yes |
| `FindService` | at the point in time of the call | yes |

Both of those methods come in three different overrides: 
* one taking an ara::com::InstanceIdentifier
* one taking an ara::core::InstanceSpecifier
* one taking NO argument

For `StartFindService`:
```cpp
static ara::com::FindServiceHandle StartFindService(
ara::com::FindServiceHandler<RadarServiceProxy::HandleType> handler,
ara::com::InstanceIdentifier instanceId);

static ara::com::FindServiceHandle StartFindService(
ara::com::FindServiceHandler<RadarServiceProxy::HandleType> handler,
ara::core::InstanceSpecifier instanceSpec);

/**
* This is an overload of the StartFindService method using neither
* instance specifier nor instance identifier.
* Semantics is, that ALL instances of the service shall be found, by
* using all available/configured technical bindings.
*
*/
static ara::com::FindServiceHandle StartFindService(
ara::com::FindServiceHandler<RadarServiceProxy::HandleType> handler);
```
For `StartFindService` -> `FindServiceHandler`,
```cpp
using FindServiceHandler = std::function<void(ServiceHandleContainer<T>, FindServiceHandle)>;
```
```cpp
FindServiceHandler<RadarServiceProxy::HandleType> = std::function<void(ServiceHandleContainer<RadarServiceProxy::HandleType>, FindServiceHandle)>
```

`StartFindService` returns a `FindServiceHandle`, which can be used to stop the
ongoing background activity of monitoring service instance availability via
call to `StopFindService`.


For `FindService`:
```cpp
static ara::com::ServiceHandleContainer<RadarServiceProxy::HandleType>
FindService(ara::com::InstanceIdentifier instanceId);
```
##### Auto Update Proxy Instance

```cpp
/**
* Reference to radar instance, we work with,
* initialized during startup
*/
RadarServiceProxy *myRadarProxy;

void radarServiceAvailabilityHandler(ServiceHandleContainer<
  RadarServiceProxy::HandleType> curHandles, FindServiceHandle handle) {
    for (RadarServiceProxy::HandleType handle : curHandles) {
        if (handle.GetInstanceId() == myRadarProxy->GetHandle().GetInstanceId()) {
            /**
            * This call on the proxy instance shall NOT lead to an exception,
            * regarding service instance not reachable, since proxy instance
            * should be already auto updated at this point in time.
            * Implementation details.
            */
            ara::core::Future<Calibrate::Output> out =
            myRadarProxy->Calibrate("test");

            // ... do something with out.
        }
    }
}
```
**Access to proxy instance within FindService handler**

### Events

Example code generated:

It is the type of one member in `RadarServiceProxy`. 

Event specific wrapper class. It is used to access events/event data.
```cpp
class BrakeEvent {
    /**
     * \brief Shortcut for the events data type.
     */
    using SampleType = RadarObjects;

    /**
     * \brief The application expects the CM to subscribe the event.
     *
     * The Communication Management shall try to subscribe and resubscribe
     * until \see Unsubscribe() is called explicitly.
     * The error handling shall be kept within the Communication Management.
     *
     * The function returns immediately. If the user wants to get notified,
     * when subscription has succeeded, he needs to register a handler
     * via \see SetSubscriptionStateChangeHandler(). This handler gets
     * then called after subscription was successful.
     *
     * \param maxSampleCount maximum number of samples, which can be held.
     */
    void Subscribe(size_t maxSampleCount);

    /**
     * \brief Query current subscription state.
     *
     * \return Current state of the subscription.
     */
    ara::com::SubscriptionState GetSubscriptionState() const;

    /**
     * \brief Unsubscribe from the service.
     */
    void Unsubscribe();

    /**
     * \brief Get the number of currently free/available sample slots.
     *
     * \return number from 0 - N (N = count given in call to Subscribe())
     *         or an ErrorCode in case of number of currently held samples
     *         already exceeds the max number given in Subscribe().
     */
    ara::core::Result<size_t> GetFreeSampleCount() const noexcept;

    /**
     * Setting a receive handler signals the Communication Management
     * implementation to use event style mode.
     * I.e., the registered handler gets called asynchronously by the
     * Communication Management as soon as new event data arrives for
     * that event. If the user wants to have strict polling behavior,
     * where no handler is called, NO handler should be registered.
     *
     * Handler may be overwritten anytime during runtime.
     *
     * Provided Handler needs not to be re-entrant since the
     * Communication Management implementation has to serialize calls
     * to the handler: Handler gets called once by the MW, when new
     * events arrived since the last call to GetNewSamples().
     *
     * When application calls GetNewSamples() again in the context of the
     * receive handler, MW must - in case new events arrived in the
     * meantime - defer next call to receive handler until after
     * the previous call to receive handler has been completed.
     */
    void SetReceiveHandler(ara::com::EventReceiveHandler handler);

    /**
     * Remove handler set by SetReceiveHandler()
     */
    void UnsetReceiveHandler();

    /**
     * Setting a subscription state change handler, which shall get
     * called by the Communication Management implementation as soon
     * as the subscription state of this event has changed.
     *
     * Communication Management implementation will serialize calls
     * to the registered handler. If multiple changes of the
     * subscription state take place during the runtime of a
     * previous call to a handler, the Communication Management
     * aggregates all changes to one call with the last/effective
     * state.
     *
     * Handler may be overwritten during runtime.
     */
    void SetSubscriptionStateChangeHandler(
         ara::com::SubscriptionStateChangeHandler handler);

    /**
     * Remove handler set by SetSubscriptionStateChangeHandler()
     */
    void UnsetSubscriptionStateChangeHandler();

    /**
     * \brief Get new data from the Communication Management
     * buffers and provide it in callbacks to the given callable f.
     *
     * \pre BrakeEvent::Subscribe has been called before
     * (and not be withdrawn by BrakeEvent::Unsubscribe)
     *
     * \param f
     * \parblock
     * callback, which shall be called with new sample.
     *
     * This callable has to fulfill signature
     * void(ara::com::SamplePtr<SampleType const>)
     * \parblockend
     *
     * \param maxNumberOfSamples
     * \parblock
     * upper bound of samples to be fetched from middleware buffers.
     * Default value means "no restriction", i.e. all newly arrived samples
     * are fetched as long as there are free sample slots.
     * \parblockend
     *
     * \return Result, which contains the number of samples,
     * which have been fetched and presented to user via calls to f or an
     * ErrorCode in case of error (e.g. precondition not fullfilled)
     */
    template <typename F>
    ara::core::Result<size_t> GetNewSamples(
         F&& f,
         size_t maxNumberOfSamples = std::numeric_limits<size_t>::max());
};
```
**Listing 6.7: Proxy side BrakeEvent Class**


Service Proxy - Event related Data type
RadarServiceProxy - BrakeEvent - RadarObjects

#### Event Subscription and Local Cache
You have to "subscript" for the event, in order to tell hte **Communication
Mangement**, that you are now interested in receiving events.

For `Subscribe`, parameter maxSampleCount can be passed. CM then knows it, also
"local cache" should be set up, which is also implemented by CM.

#### Monitoring Event Subscription
LEAVE BLANK.

#### Accessing Event Data - aka Samples

```cpp
class BrakeEvent {
    ...

    /**
     * \brief Get new data from the Communication Management
     * buffers and provide it in callbacks to the given callable f.
     *
     * \pre BrakeEvent::Subscribe has been called before
     * (and not be withdrawn by BrakeEvent::Unsubscribe)
     *
     * \param f
     * \parblock
     * callback, which shall be called with new sample.
     *
     * This callable has to fulfill signature
     * void(ara::com::SamplePtr<SampleType const>)
     * \parblockend
     *
     * \param maxNumberOfSamples
     * \parblock
     * upper bound of samples to be fetched from middleware buffers.
     * Default value means "no restriction", i.e. all newly arrived samples
     * are fetched as long as there are free sample slots.
     * \parblockend
     *
     * \return Result, which contains the number of samples,
     * which have been fetched and presented to user via calls to f or an
     * ErrorCode in case of error (e.g. precondition not fullfilled)
     */
    template <typename F>
    ara::core::Result<size_t> GetNewSamples(
         F&& f,
         size_t maxNumberOfSamples = std::numeric_limits<size_t>::max());
};
```
#### Event Sample Management via `SamplePtrs`
#### Event-Driven v.s. Polling-Based access

```cpp
#include "RadarServiceProxy.hpp"
#include <memory>
#include <deque>

using namespace com::mycompany::division::radarservice;
using namespace ara::com;

/**
  * our radar proxy - initially the unique ptr is invalid.
  */
std::unique_ptr<proxy::RadarServiceProxy> myRadarProxy;

/**
  * a storage for BrakeEvent samples in fifo style
  */
std::deque<SamplePtr<const proxy::events::BrakeEvent::SampleType>>
     lastNActiveSamples;

/**
  * \brief application function, which processes current set of BrakeEvent
  * samples.
  * \param samples
  */
void processLastBrakeEvents(
      std::deque<SamplePtr<const proxy::events::BrakeEvent::SampleType>>&
     samples) {
      // do whatever with those BrakeEvent samples ...
}

/**
  * \brief event reception handler for BrakeEvent events, which we register
     to get informed about new events.
  */
void handleBrakeEventReception() {
      /**
       * we get newly arrived BrakeEvent events into our process space.
       * For each sample we get passed in, we check for a certain property
       * "active" and if it fulfills the check, we move it into our Last10-storage.
       * So this few lines basically implement filtering and a LastN policy.
       */
      myRadarProxy->BrakeEvent.GetNewSamples(
      [](SamplePtr<proxy::events::BrakeEvent::SampleType> samplePtr) {
              if(samplePtr->active) {
                  lastNActiveSamples.push_back(std::move(samplePtr));
                  if (lastNActiveSamples.size() > 10)
                      lastNActiveSamples.pop_front();
              }
          });

      // ... now process those samples ...
      processLastBrakeEvents(lastNActiveSamples);
}

int main(int argc, char** argv) {

      auto handles = proxy::RadarServiceProxy::FindService();

      if (!handles.empty()) {
          /* we have at least one valid handle - we are not very particular
     here and take the first one to
           * create our proxy */
          myRadarProxy = std::make_unique<proxy::RadarServiceProxy>(handles[0]);

          /* we are interested in receiving the event "BrakeEvent" - so we
          subscribe for it. We want to access up to 10 events, since our sample
          algo averages over at most 10.*/
          myRadarProxy->BrakeEvent.Subscribe(10);

          /* whenever new BrakeEvent events come in, we want be called, so we register a callback for it!
           * Note: If the entity we would subscribe to, would be a field
           * instead of an event, it would be crucial, to register our
           * reception handler BEFORE subscribing, to avoid race conditions.
           * After a field subscription, you would get instantly so called
           * "initial events" and to be sure not to miss them, you should care
           * for that your reception handler is registered before.*/
           myRadarProxy->BrakeEvent.SetReceiveHandler(
               handleBrakeEventReception);
      }

      // ... wait for application shutdown trigger by application exec mgmt.
}
```
                            Listing 6.8: Sample Code how to access Events

Explanation:
In `myRadarProxy->BrakeEvent.GetNewSamples`
- the first parameter, `F&& f`.

    Function:

    [&& Meaning](https://stackoverflow.com/questions/28066777/const-and-specifiers-for-member-functions-in-c):

    && means, that this overload will be used only for rvalue object. So for
    this F&&, it depends on F. And F is `void(ara::com::SamplePtr<SampleType const>)`

    ```cpp
      [](SamplePtr<proxy::events::BrakeEvent::SampleType> samplePtr) {
      ...
      }
      // https://stackoverflow.com/questions/12483753/how-do-define-anonymous-functions-in-c
      // https://en.wikipedia.org/wiki/Anonymous_function#C++_(since_C++11)
    ```
- the second parameter has default value


#### Buffering Strategies
Implementation specific. Skip.

### Methods
#### Event-Driven v.s. Polling access to method results

```cpp
 enum class future_status : uint8_t
 {
 ready,    ///< the shared state is ready
 timeout,     ///< the shared state did not become ready before the specified timeout has passed
 };

 template <typename T, typename E = ErrorCode>
 class Future {
    public:

       Future() noexcept = default;
       ~Future();

       Future(Future const&) = delete;
       Future& operator=(Future const&) = delete;

       Future(Future&& other) noexcept;
       Future& operator=(Future&& other) noexcept;

       /**
        * @brief Get the value.
        *
        * This function shall behave the same as the corresponding std::future function.
        *
        * @returns value of type T
        * @error Domain:error the error that has been put into the corresponding Promise via Promise::SetError
        *
        */
       T get();

       /**
        * @brief Get the result.
        *
        * Similar to get(), this call blocks until the value or an error is available.
        * However, this call will never throw an exception.
        *
        * @returns a Result with either a value or an error
        * @error Domain:error the error that has been put into the corresponding Promise via Promise::SetError
        *
        */
       Result<T, E> GetResult() noexcept;

       /**
        * @brief Checks if the Future is valid, i.e. if it has a shared state.
        *
        * This function shall behave the same as the corresponding std::future function.
        *
        * @returns true if the Future is usable, false otherwise
        */
       bool valid() const noexcept;

       /**
        * @brief Wait for a value or an error to be available.
        *
        * This function shall behave the same as the corresponding std::future function.
        */
       void wait() const;

       /**
        * @brief Wait for the given period, or until a value or an error is available.
        *
        * This function shall behave the same as the corresponding std::future function.
        *
        * @param timeoutDuration maximal duration to wait for
        * @returns status that indicates whether the timeout hit or if a value is available
        */
       template <typename Rep, typename Period>
       future_status wait_for(std::chrono::duration<Rep, Period> const& timeoutDuration) const;

     /**
      * @brief Wait until the given time, or until a value or an error is available.
      *
      * This function shall behave the same as the corresponding std::future function.
      *
      * @param deadline latest point in time to wait
      * @returns status that indicates whether the time was reached or if a
     */
      template <typename Clock, typename Duration>
      future_status wait_until(std::chrono::time_point<Clock, Duration> const
    & deadline) const;

      /**
       * @brief Register a callable that gets called when the Future becomes ready.
       *
       * When @a func is called, it is guaranteed that get() and GetResult() will not block.
       *
       * @a func may be called in the context of this call or in the context of Promise::set_value()
       * or Promise::SetError() or somewhere else.
       *
       * The return type of @a then depends on the return type of @a func (aka continuation).
       *
       * Let U be the return type of the continuation (i.e. std::result_of_t< std::decay_t<F>(ara::core::Future<T,E>)>).
       * If U is ara::core::Future<T2,E2> for some types T2, E2, then the return type of @a then is ara::core::Future<T2,E2>,
       * otherwise it is ara::core::Future<U>. This is known as implicit unwrapping.
       *
       * @param func a callable to register
       * @returns a new Future instance for the result of the continuation
       */
      template <typename F>
      auto then(F&& func) -> SEE_COMMENT_ABOVE;

      /**
       * @brief Return whether the asynchronous operation has finished.
       *
       * If this function returns true, get(), GetResult() and the wait calls are guaranteed not to block.
       *
       * @returns true if the Future contains a value or an error, false otherwise
       */
      bool is_ready() const;
};
```
                             Listing 6.11: ara::core::Future Class
### Fields
Just like Event + Method

## Skeleton Class
### Example Code
```cpp
class RadarServiceSkeleton {
  public:
      /**
       * Ctor taking instance identifier as parameter and having default
       * request processing mode kEvent.
       */
      RadarServiceSkeleton(ara::com::InstanceIdentifier instanceId,
      ara::com::MethodCallProcessingMode mode =
      ara::com::MethodCallProcessingMode::kEvent);

      /**
       * Ctor taking instance identifier container as parameter and having
       * default request processing mode kEvent.
       * This specifically supports multi-binding.
       */
      RadarServiceSkeleton(ara::com::InstanceIdentifierContainer instanceIds,
      ara::com::MethodCallProcessingMode mode =
      ara::com::MethodCallProcessingMode::kEvent);

      /**
       * Ctor taking instance specifier as parameter and having default
       * request processing mode kEvent.
       */
      RadarServiceSkeleton(ara::core::InstanceSpecifier instanceSpec,
      ara::com::MethodCallProcessingMode mode =
      ara::com::MethodCallProcessingMode::kEvent);

      /**
       * skeleton instances are nor copy constructible.
       */
      RadarServiceSkeleton(const RadarServiceSkeleton& other) = delete;

      /**
       * skeleton instances are nor copy assignable.
       */
      RadarServiceSkeleton& operator=(const RadarServiceSkeleton& other) = delete;

      /**
       * The Communication Management implementer should care in his dtor
       * implementation, that the functionality of StopOfferService()
       * is internally triggered in case this service instance has
       * been offered before. This is a convenient cleanup functionality.
       */
      ~RadarServiceSkeleton();

      /**
       * Offer the service instance.
       * method is idempotent - could be called repeatedly.
       */
      void OfferService();

      /**
       * Stop Offering the service instance.
       * method is idempotent - could be called repeatedly.
       *
       * If service instance gets destroyed - it is expected that the
       * Communication Management implementation calls StopOfferService()
       * internally.
       */
      void StopOfferService();

      /**
       * For all output and non-void return parameters
       * an enclosing struct is generated, which contains
       * non-void return value and/or out parameters.
       */
      struct CalibrateOutput {
              bool result;
          };

      /**
       * For all output and non-void return parameters
       * an enclosing struct is generated, which contains
       * non-void return value and/or out parameters.
       */
      struct AdjustOutput {
              bool success;
              Position effective_position;
          };

      /**
       * This fetches the next call from the Communication Management
       * and executes it. The return value is a ara::core::Future.
       * In case of an Application Error, an ara::core::ErrorCode is stored
       * in the ara::core::Promise from which the ara::core::Future
       * is returned to the caller.
       * Only available in polling mode.
       */
      ara::core::Future<bool> ProcessNextMethodCall();

      /**
       * \brief Public member for the BrakeEvent
       */
      events::BrakeEvent BrakeEvent;

      /**
       * \brief Public member for the UpdateRate
       */
      fields::UpdateRate UpdateRate;

      /**
       * The following methods are pure virtual and have to be implemented
       */
      virtual ara::core::Future<CalibrateOutput> Calibrate(
      std::string configuration) = 0;
      virtual ara::core::Future<AdjustOutput> Adjust(
      const Position& position) = 0;
      virtual void LogCurrentState() = 0;
};
```
### Instantiation

[**Skeleton**](https://www.merriam-webster.com/dictionary/skeleton): a usually rigid supportive or protective structure or framework
of an organism.

Service Implementation should subclass from the skeleton class.

Identifier (in the `ctor`) should be unique.

A static member function `Preconstruct` checks the provided identifier.

### Offering Service Instance

```cpp
using namespace ara::com;

/**
  * Our implementation of the RadarService -
  * subclass of RadarServiceSkeleton
  */
class RadarServiceImpl;

int main(int argc, char** argv) {
      // read instanceId from commandline
      ara::core::string_view instanceIdStr(argv[1]);
      RadarServiceImpl myRadarService(InstanceIdentifier(instanceIdStr));

      // do some service specific initialization here ....
      myRadarService.init();

      // now service instance is ready -> make it visible/available
      myRadarService.OfferService();

      // go into some wait state in main thread - waiting for AppExecMgr
      // signals or the like ....

      return 0;
}
```

### Polling and event-driven processing modes
**Procssing**
Default one is `kEvent`.

```cpp
/**
  * Request processing modes for the service implementation side
  * (skeleton).
  *
  * \note Should be provided by platform vendor exactly like this.
  */
enum class MethodCallProcessingMode { kPoll, kEvent, kEventSingleThread };
```

#### Polling Mode
In `kPoll`, the Communication Management implementation will not call any of the
provided service methods asynchronously.

```cpp
/**
  * This fetches the next call from the Communication Management
  * and executes it. The return value is a ara::core::Future.
  * In case of an Application Error, an ara::core::ErrorCode is stored
  * in the ara::core::Promise from which the ara::core::Future
  * is returned to the caller.
  * Only available in polling mode.
  */
ara::core::Future<bool> ProcessNextMethodCall();
```

#### Event-Driven Mode

### Methods

```cpp
/**
 * For all output and non-void return parameters
 * an enclosing struct is generated, which contains
 * non-void return value and/or out parameters.
 */
struct AdjustOutput {
      bool success;
      Position effective_position;
};
```

