package target

var v1Error2 = []byte(`{"errors":[{"context":{"custom":{},"request":{"body":null,"cookies":{},"env":{"REMOTE_ADDR":"127.0.0.1","SERVER_NAME":"1.0.0.127.in-addr.arpa","SERVER_PORT":"8000"},"headers":{"accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8","accept-encoding":"gzip, deflate, br","accept-language":"en-US,en;q=0.9","connection":"keep-alive","content-length":"","content-type":"text/plain","host":"localhost:8000","upgrade-insecure-requests":"1","user-agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36"},"method":"GET","socket":{"encrypted":false,"remote_address":"127.0.0.1"},"url":{"full":"http://localhost:8000/oopsie","hostname":"localhost","pathname":"/oopsie","port":"8000","protocol":"http:"}},"user":{"id":null,"is_authenticated":false,"username":""}},"culprit":"opbeans.views.oopsie","exception":{"message":"AssertionError: ","module":"builtins","stacktrace":[{"abs_path":"/opbeans/lib/python3.6/site-packages/django/core/handlers/exception.py","context_line":"            response = get_response(request)","filename":"django/core/handlers/exception.py","function":"inner","library_frame":true,"lineno":41,"module":"django.core.handlers.exception","post_context":["        except Exception as exc:","            response = response_for_exception(request, exc)"],"pre_context":["    @wraps(get_response, assigned=available_attrs(get_response))","    def inner(request):","        try:"],"vars":{"exc":"AssertionError()","get_response":"\u003cbound method BaseHandler._get_response of \u003cdjango.core.handlers.wsgi.WSGIHandler object at 0x109764a20\u003e\u003e","request":"\u003cWSGIRequest: GET '/oopsie'\u003e"}},{"abs_path":"/opbeans/lib/python3.6/site-packages/django/core/handlers/base.py","context_line":"                response = self.process_exception_by_middleware(e, request)","filename":"django/core/handlers/base.py","function":"_get_response","library_frame":true,"lineno":187,"module":"django.core.handlers.base","post_context":["","        # Complain if the view returned None (a common error)."],"pre_context":["            try:","                response = wrapped_callback(request, *callback_args, **callback_kwargs)","            except Exception as e:"],"vars":{"callback":"\u003cfunction oopsie at 0x10ab1e9d8\u003e","callback_args":[],"callback_kwargs":{},"middleware_method":"\u003cbound method CsrfViewMiddleware.process_view of \u003cdjango.middleware.csrf.CsrfViewMiddleware object at 0x10aaf5080\u003e\u003e","request":"\u003cWSGIRequest: GET '/oopsie'\u003e","resolver":"\u003cRegexURLResolver 'opbeans.urls' (None:None) ^/\u003e","resolver_match":"ResolverMatch(func=opbeans.views.oopsie, args=(), kwargs={}, url_name=None, app_names=[], namespaces=[])","response":null,"self":"\u003cdjango.core.handlers.wsgi.WSGIHandler object at 0x109764a20\u003e","wrapped_callback":"\u003cfunction oopsie at 0x10ab1e9d8\u003e"}},{"abs_path":"/opbeans/lib/python3.6/site-packages/django/core/handlers/base.py","context_line":"                response = wrapped_callback(request, *callback_args, **callback_kwargs)","filename":"django/core/handlers/base.py","function":"_get_response","library_frame":true,"lineno":185,"module":"django.core.handlers.base","post_context":["            except Exception as e:","                response = self.process_exception_by_middleware(e, request)"],"pre_context":["        if response is None:","            wrapped_callback = self.make_view_atomic(callback)","            try:"],"vars":{"callback":"\u003cfunction oopsie at 0x10ab1e9d8\u003e","callback_args":[],"callback_kwargs":{},"middleware_method":"\u003cbound method CsrfViewMiddleware.process_view of \u003cdjango.middleware.csrf.CsrfViewMiddleware object at 0x10aaf5080\u003e\u003e","request":"\u003cWSGIRequest: GET '/oopsie'\u003e","resolver":"\u003cRegexURLResolver 'opbeans.urls' (None:None) ^/\u003e","resolver_match":"ResolverMatch(func=opbeans.views.oopsie, args=(), kwargs={}, url_name=None, app_names=[], namespaces=[])","response":null,"self":"\u003cdjango.core.handlers.wsgi.WSGIHandler object at 0x109764a20\u003e","wrapped_callback":"\u003cfunction oopsie at 0x10ab1e9d8\u003e"}},{"abs_path":"/Users/gil/.local/go/src/github.com/elastic/opbeans-python/opbeans/views.py","context_line":"    assert False","filename":"opbeans/views.py","function":"oopsie","library_frame":false,"lineno":188,"module":"opbeans.views","post_context":[],"pre_context":["","def oopsie(request):","    client.capture_message('About to blow up!')"],"vars":{"request":"\u003cWSGIRequest: GET '/oopsie'\u003e"}}],"type":"AssertionError"},"handled":false,"id":"47a58c1a-6ab5-4bd2-965d-3d0d0114a588","timestamp":"2018-01-09T03:36:12.813666Z","transaction":{"id":"b2f39758-f9fc-40ef-8571-a3198f4cb40a"}}],"process":{"argv":["./manage.py","runserver"],"pid":52687,"title":null},"service":{"agent":{"name":"python","version":"1.0.0"},"environment":null,"framework":{"name":"django","version":"1.11.8"},"language":{"name":"python","version":"3.6.3"},"name":"opbeans-python","runtime":{"name":"CPython","version":"3.6.3"},"version":null},"system":{"architecture":"x86_64","hostname":"localhost.localdomain","platform":"darwin"}}`)